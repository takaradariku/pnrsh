package pnr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"crypto/tls"
	"net/url"
	"os"
	"time"
)

const reqEndpoint = `https://apis.alaskaair.com/1/guestservices/customermobile/mobileservices/reservation/details/{conf}?lastName={lname}`

var (
	requestHeaders = map[string]string{
		"Accept":                    "application/json",
		"User-Agent":                "Alaska Airlines for Android-8.16.0 (ALKApp/Android)",
		"Ocp-Apim-Subscription-Key": "de1d0ff837444468a5ea868945aab738",
	}

	client = http.Client{
		Timeout: time.Second * 15,

		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy: func(r *http.Request) (*url.URL, error) {
				if p := os.Getenv("HTTP_PROXY"); p != "" {
					return url.Parse(p)
				}

				return nil, nil
			},
		},
	}
)

func sendRequest(lastName, confirmationCode string) ([]byte, error) {
	endpoint := reqEndpoint
	endpoint = strings.Replace(endpoint, "{conf}", confirmationCode, -1)
	endpoint = strings.Replace(endpoint, "{lname}", lastName, -1)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []byte{}, err
	}

	for k, v := range requestHeaders {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)

	if res.StatusCode != 200 {
		return []byte{}, errors.New("status code was not 200")
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func performRequest(lastName, confirmationCode string) (res PNR, err error) {
	data, err := sendRequest(lastName, confirmationCode)

	if err != nil {
		return res, err
	}

	var response PNR
	if err := json.Unmarshal(data, &response); err != nil {
		return res, err
	}

	return response, nil
}
