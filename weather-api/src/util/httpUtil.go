package util

import (
	"io/ioutil"
	"net/http"
)

// SendGetHTTPRequest GETのHTTPリクエストを投げる
func SendGetHTTPRequest(sendURL string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", sendURL, nil)
	if err != nil {
		return ""
	}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}
