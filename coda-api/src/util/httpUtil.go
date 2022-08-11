package util

import (
	"bytes"
	"encoding/base64"

	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"

	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SendGetHTTPRequest GETのHTTPリクエストを投げる
func SendGetHTTPRequest(sendURL string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", sendURL, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// SendGetHTTPRequestWithBearerToken BearerToken付きでGETのHTTPリクエストを投げる
func SendGetHTTPRequestWithBearerToken(sendURL string, bearerToken string) (string, error) {
	bearer := "Bearer " + bearerToken
	client := &http.Client{}
	req, err := http.NewRequest("GET", sendURL, nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// SendGetHTTPRequestWithUserAgent UserAgent付きのGETのHTTPリクエストを投げる
func SendGetHTTPRequestWithUserAgent(sendURL string, userAgent string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", sendURL, nil)
	if err != nil {
		return "", err
	}
	if userAgent != "" {
		req.Header.Set("User-Agent", userAgent)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(body)
	bReader := bytes.NewReader(body)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	return buf.String(), nil
}

// GetOgpUserAgentBotFormURL OGP取得のためのUserAgentを取得
func GetOgpUserAgentBotFormURL(targetUrl string) string {
	domain, _ := GetDomainFromUrl(targetUrl)
	if strings.HasSuffix(domain, "twitter.com") || strings.HasSuffix(domain, "facebook.com") || strings.HasSuffix(domain, "instagram.com") {
		// tiwtter・facebook・instagramの場合はbot
		return "bot"

	} else if strings.HasSuffix(domain, "amazon.co.jp") || strings.HasSuffix(domain, "youtube.com") {
		// amazon・youtubeはGooglebot
		return "Googlebot"
	}
	return "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"
}

// SendGetHTTPRequestForBase64Image Base64Image取得のHTTPリクエストを投げる
func SendGetHTTPRequestForBase64Image(sendURL string, userAgent string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", sendURL, nil)
	if err != nil {
		return "", err
	}
	if userAgent != "" {
		req.Header.Set("User-Agent", userAgent)
	}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	base64Data := base64.StdEncoding.EncodeToString(body)
	return base64Data, nil
}

// GetDomainFromUrl URLからドメインを取得
func GetDomainFromUrl(targetUrl string) (string, error) {
	u, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}
	domain := u.Hostname()
	return domain, nil
}
