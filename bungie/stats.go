package bungie

type Basic struct {
	Value        float64 `json:"value"`
	DisplayValue string  `json:"displayValue"`
}

type Stat struct {
	Basic Basic `json:"basic"`
}

type StatWithID struct {
	StatID string `json:"statId"`
	Basic  Basic  `json:"basic"`
}
