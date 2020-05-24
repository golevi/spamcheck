package spamcheck

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Options are either long or short
type Options string

const (
	// Short gets the score
	Short Options = "short"
	// Long gets a full report
	Long Options = "long"
)

type rules struct {
	Score       string `json:"score"`
	Description string `json:"description"`
}

// Request is the email and options
type Request struct {
	Email   string  `json:"email"`
	Options Options `json:"options"`
}

// Response is what is received from the api
type Response struct {
	Score  string  `json:"score"`
	Report string  `json:"report"`
	Rules  []rules `json:"rules"`
}

// Process the request
func (r *Request) Process() (*Response, error) {
	apiResponse := Response{}

	jsn, err := json.Marshal(&r)
	if err != nil {
		return &apiResponse, err
	}
	req, err := http.NewRequest("POST", "https://spamcheck.postmarkapp.com/filter", bytes.NewBuffer(jsn))
	if err != nil {
		return &apiResponse, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &apiResponse, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &apiResponse, err
	}

	// log.Println(string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &apiResponse)
	if err != nil {
		return &apiResponse, err
	}

	return &apiResponse, nil
}
