package bungie

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"reflect"
	"strings"
	"time"
)

const (
	baseURL      = "https://www.bungie.net/Platform/Destiny/"
	apiKeyHeader = "X-API-Key"
	apiKey       = "TODO"
	maxCacheSize = 20000
)

type cache map[string]interface{}

func (c cache) insert(key string, value interface{}) {
	if len(c) >= maxCacheSize {
		for k := range c {
			delete(c, k)
			if len(c) < maxCacheSize {
				break
			}
		}
	}
	c[key] = value
}

type API struct {
	client       http.Client
	cookie       string
	xcsrf        string
	cachePath    string
	cache        cache
	getThrottle  *time.Ticker
	postThrottle *time.Ticker
}

func New() (*API, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	return &API{
		client: http.Client{Timeout: 10 * time.Second},
		cachePath: usr.HomeDir + string(os.PathSeparator) + "bungie" +
			string(os.PathSeparator) + "cache" +
			string(os.PathSeparator),
		postThrottle: time.NewTicker(1 * time.Second),
		getThrottle:  time.NewTicker(50 * time.Millisecond),
		cache:        cache{},
	}, nil
}

func (b *API) SetCookie(cookie string) {
	b.cookie = cookie
}

func (b *API) SetXCSRF(xcsrf string) {
	b.xcsrf = xcsrf
}

type jsonStatusFields struct {
	ErrorCode       int64    `json:"ErrorCode"`
	ThrottleSeconds int64    `json:"ThrottleSeconds"`
	ErrorStatus     string   `json:"ErrorStatus"`
	Message         string   `json:"Message"`
	MessageData     struct{} `json:"MessageData"`
}

func (status jsonStatusFields) checkStatus() error {
	if status.ErrorCode != 1 {
		return fmt.Errorf("bungie.net returned error: %s (%d)\n%s",
			status.ErrorStatus, status.ErrorCode, status.Message)
	}
	return nil
}

type jsonResponse interface {
	checkStatus() error
}

func hash(url string) string {
	url = strings.Replace(url, "/", "_", -1)
	url = strings.Replace(url, "?", "_", -1)
	url = strings.Replace(url, "&", "_", -1)
	url = strings.Replace(url, "=", "_", -1)
	return url
}

func (b *API) lookup(url string, x jsonResponse) error {
	h := hash(url)

	if v, ok := b.cache[h]; ok {
		xValue := reflect.ValueOf(x)
		reflect.Indirect(xValue).Set(
			reflect.Indirect(reflect.ValueOf(v)))
		return nil
	}

	f, err := os.Open(b.cachePath + h)
	if err != nil {
		return err
	}
	defer f.Close()
	gzipReader, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzipReader.Close()
	err = fillFromReader(x, gzipReader)
	if err == nil {
		b.cache.insert(h, x)
	}
	return err
}

func (b *API) insert(url string, reader io.Reader, x jsonResponse) error {
	h := hash(url)
	f, err := os.Create(b.cachePath + h)
	if err != nil {
		return err
	}
	defer f.Close()
	writer := gzip.NewWriter(f)
	defer writer.Close()
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	writer.Write(bytes)
	err = fillFromBytes(x, bytes)
	if err == nil {
		b.cache.insert(h, x)
	}
	return err
}

func fillFromBytes(x jsonResponse, b []byte) error {
	buffer := bytes.NewBuffer(b)
	return fillFromReader(x, buffer)
}

func fillFromReader(x jsonResponse, reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(x)
	if err != nil {
		return err
	}
	return x.checkStatus()
}

type postResponse struct {
	jsonStatusFields
	Response int64 `json:"Response"`
}

func (b *API) post(url string, body interface{}) error {
	if b.cookie == "" || b.xcsrf == "" {
		return errors.New("Cannot post without cookie and xcsrf")
	}
	jsonBytes, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return err
	}
	readCloser, err := b.actuallyIssueRequest("POST", url,
		bytes.NewReader(jsonBytes))
	if err != nil {
		return err
	}
	defer readCloser.Close()
	return fillFromReader(&postResponse{}, readCloser)
}

func (b *API) get(url string, x jsonResponse, cache bool) error {
	if !cache {
		readCloser, err := b.actuallyIssueRequest("GET", url, nil)
		if err != nil {
			return err
		}
		defer readCloser.Close()
		return fillFromReader(x, readCloser)
	}
	err := b.lookup(url, x)
	if err == nil {
		return nil
	}
	readCloser, err := b.actuallyIssueRequest("GET", url, nil)
	if err != nil {
		return err
	}
	defer readCloser.Close()
	return b.insert(url, readCloser, x)
}

func (b *API) actuallyIssueRequest(httpVerb string, url string,
	body io.Reader) (io.ReadCloser, error) {
	switch httpVerb {
	case "GET":
		<-b.getThrottle.C
	case "POST":
		<-b.postThrottle.C
	default:
		return nil, errors.New("Unrecognized HTTP verb")
	}
	req, err := http.NewRequest(httpVerb, baseURL+url, body)
	log.Printf("Bungie API: %s %s%s\n", httpVerb, baseURL, url)
	if err != nil {
		return nil, err
	}
	req.Header.Set(apiKeyHeader, apiKey)
	if b.cookie != "" && b.xcsrf != "" {
		req.Header.Set("Cookie", b.cookie)
		req.Header.Set("x-csrf", b.xcsrf)
	}
	var resp *http.Response
	for i := 0; i < 3; i++ {
		resp, err = b.client.Do(req)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
