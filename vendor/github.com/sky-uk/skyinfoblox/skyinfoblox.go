package skyinfoblox

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// NewInfobloxClient  Creates a new infobloxClient object.
func NewInfobloxClient(url string, user string, password string, ignoreSSL bool, debug bool) *InfobloxClient {
	infobloxClient := new(InfobloxClient)
	infobloxClient.URL = url
	infobloxClient.User = user
	infobloxClient.Password = password
	infobloxClient.IgnoreSSL = ignoreSSL
	infobloxClient.Debug = debug
	return infobloxClient
}

// InfobloxClient struct.
type InfobloxClient struct {
	URL       string
	User      string
	Password  string
	IgnoreSSL bool
	Debug     bool
}

// Do - makes the API call.
func (infobloxClient *InfobloxClient) Do(api api.InfobloxAPI) error {
	requestURL := fmt.Sprintf("%s%s", infobloxClient.URL, api.Endpoint())
	var requestPayload io.Reader

	// TODO: change this to JSON
	if api.RequestObject() != nil {
		requestJSONBytes, marshallingErr := json.Marshal(api.RequestObject())
		if marshallingErr != nil {
			log.Fatal(marshallingErr)
			return (marshallingErr)
		}
		if infobloxClient.Debug {
			log.Println("Request payload as JSON:")
			log.Println(string(requestJSONBytes))
			log.Println("--------------------------------------------------------------")
		}
		requestPayload = bytes.NewReader(requestJSONBytes)
	}
	if infobloxClient.Debug {
		log.Println("requestURL:", requestURL)
	}
	req, err := http.NewRequest(api.Method(), requestURL, requestPayload)
	if err != nil {
		log.Println("ERROR building the request: ", err)
		return err
	}

	req.SetBasicAuth(infobloxClient.User, infobloxClient.Password)

	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: infobloxClient.IgnoreSSL},
	}
	httpClient := &http.Client{Transport: tr}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println("ERROR executing request: ", err)
		return err
	}
	defer res.Body.Close()
	return infobloxClient.handleResponse(api, res)
}

func (infobloxClient *InfobloxClient) handleResponse(api api.InfobloxAPI, res *http.Response) error {
	api.SetStatusCode(res.StatusCode)
	bodyText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ERROR reading response: ", err)
		return err
	}

	api.SetRawResponse(bodyText)

	if infobloxClient.Debug {
		log.Println(string(bodyText))
	}

	if isJSON(res.Header.Get("Content-Type")) && api.StatusCode() == 200 {
		JSONerr := json.Unmarshal(bodyText, api.ResponseObject())
		if JSONerr != nil {
			log.Println("ERROR unmarshalling response: ", JSONerr)
			return err
		}
	} else {
		api.SetResponseObject(string(bodyText))
	}
	return nil
}

func isJSON(contentType string) bool {
	return strings.Contains(strings.ToLower(contentType), "/json")
}
