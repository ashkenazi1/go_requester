package go_requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Requester struct {
	Client  *http.Client
	Headers map[string]string
}

func New() *Requester {
	return &Requester{
		Client:  &http.Client{Timeout: 10 * time.Second},
		Headers: make(map[string]string),
	}
}

func (r *Requester) Get(url string) ([]byte, error) {
	return r.do("GET", url, nil)
}

func (r *Requester) Post(url string, body io.Reader) ([]byte, error) {
	return r.do("POST", url, body)
}

func (r *Requester) PostJSON(url string, body interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonBytes)
	return r.Post(url, bodyReader)
}

func (r *Requester) Patch(url string, body io.Reader) ([]byte, error) {
	return r.do("PATCH", url, body)
}

func (r *Requester) Put(url string, body io.Reader) ([]byte, error) {
	return r.do("PUT", url, body)
}

func (r *Requester) Delete(url string) ([]byte, error) {
	return r.do("DELETE", url, nil)
}

func (r *Requester) do(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(resBody))
	}

	return resBody, nil
}

func (r *Requester) SetHeader(key, value string) {
	r.Headers[key] = value
}

func (r *Requester) SetTimeout(timeout time.Duration) {
	r.Client.Timeout = timeout
}
