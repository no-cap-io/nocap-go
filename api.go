package NoCap

import (
	"io/ioutil"
	"net/http"
)

const (
	CapEndpoint = "https://no-cap.io"
)

// request sends an HTTP request and returns the response
// body. An error is returned if the http request wasn't
// successful.
func request(req *http.Request) (string, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
