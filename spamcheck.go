// Package spamcheck uses Postmark's spam API. The API is free to use. No API
// key is required. For more information visit https://spamcheck.postmarkapp.com/.
/*
	scr := spamcheck.NewRequest("I am a nigerian prince and will give you $1 million")
	scr = spamcheck.Short

	resp, err := scr.CheckScore()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	// {true 7.9 [] }
*/
package spamcheck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://spamcheck.postmarkapp.com/filter"

// Options dictates the amount of information returned from the API.
type Options string

const (
	// Long for a full report of processing rules. Default option.
	Long Options = "long"
	// Short for a score request.
	Short Options = "short"
)

// Request is used to send information to the API.
type Request struct {
	Email   string  `json:"email"`
	Options Options `json:"options"`
}

// Response is what is returned from the API.
type Response struct {
	Success bool       `json:"success"`
	Score   string     `json:"score"`
	Rules   []spamRule `json:"rules"`
	Report  string     `json:"report"`
}

type spamRule struct {
	Score       string `json:"score"`
	Description string `json:"description"`
}

// GetReport returns a long response
func (r *Request) GetReport() {

}

// GetScore gets only the score
func (r *Request) GetScore() {

}

func process(email, options string) {

}

// NewRequest is a wrapper for creating a new Request.
func NewRequest(input string) *Request {
	spamReq := &Request{
		Email:   input,
		Options: Long,
	}

	return spamReq
}

// CheckScore checks the spam score
func (s *Request) CheckScore() (Response, error) {
	if s.Options == "" {
		s.Options = Long
	}

	var spamResponse Response

	jsn, err := json.Marshal(s)
	if err != nil {
		return spamResponse, err
	}

	postData, err := json.Marshal(map[string]string{
		"email":   s.Email,
		"options": string(s.Options),
	})

	resp, err := http.Post(
		"https://spamcheck.postmarkapp.com/filter",
		"application/json",
		bytes.NewBuffer(postData),
	)
	defer resp.Body.Close()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsn))
	if err != nil {
		return spamResponse, err
	}
	r := &Response{}
	err = json.NewDecoder(resp.Body).Decode(r)
	fmt.Println(r.Score)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return spamResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return spamResponse, err
	}

	json.Unmarshal(body, &spamResponse)

	return spamResponse, nil
}
