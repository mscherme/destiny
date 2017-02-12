package bungie

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"time"
)

const (
	baseURL      = "https://www.bungie.net/Platform/Destiny/"
	apiKeyHeader = "X-API-Key"
	apiKey       = "7a1778148fa4464cab117143cb83f59c"
)

type API struct {
	client    http.Client
	cookie    string
	xcsrf     string
	cachePath string
}

func New() (*API, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	return &API{
		client: http.Client{Timeout: 10 * time.Second},
		cachePath: usr.HomeDir + string(os.PathSeparator) + "bungie" +
			string(os.PathSeparator) + "cache" + string(os.PathSeparator),
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
	f, err := os.Open(b.cachePath + h)
	if err != nil {
		return err
	}
	defer f.Close()
	return fill(x, bufio.NewReader(f))
}

func (b *API) insert(url string, reader io.Reader) error {
	h := hash(url)
	f, err := os.Create(b.cachePath + h)
	if err != nil {
		return err
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	_, err = writer.ReadFrom(reader)
	writer.Flush()
	return err
}

func fill(x jsonResponse, reader io.Reader) error {
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
	readCloser, err := b.actuallyIssueRequest("POST", url, bytes.NewReader(jsonBytes))
	if err != nil {
		return err
	}
	defer readCloser.Close()
	return fill(&postResponse{}, readCloser)
}

func (b *API) get(url string, x jsonResponse, cache bool) error {
	if !cache {
		readCloser, err := b.actuallyIssueRequest("GET", url, nil)
		if err != nil {
			return err
		}
		defer readCloser.Close()
		return fill(x, readCloser)
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
	err = b.insert(url, readCloser)
	if err != nil {
		return err
	}
	return b.lookup(url, x)
}

var throttle = time.Tick(50 * time.Millisecond)

func (b *API) actuallyIssueRequest(httpVerb string, url string, body io.Reader) (io.ReadCloser, error) {
	<-throttle
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
	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
