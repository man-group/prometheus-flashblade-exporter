// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

var xAuthToken string

type FlashbladeClient struct {
	Host       string
	client     *http.Client
	ApiVersion string
}

type Version struct {
	Versions []string `json:"versions"`
}

func NewFlashbladeClient(host string, insecure bool, version string) *FlashbladeClient {
	client := &http.Client{}
	var fb FlashbladeClient

	if insecure {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: transport}
	}

	fb = FlashbladeClient{client: client, Host: host}

	// Init x-auth-token
	fb.refreshXAuthToken()
	fb.ApiVersion = fb.getAPIVersion(version)
	return &fb
}

func (fbClient *FlashbladeClient) getAPIVersion(selectedVersion string) string {
	var (
		params map[string]string
		v      Version
	)
	err := fbClient.GetJSON("api_version", params, &v)
	if err != nil {
		log.Fatalln("Failed to get supported API-versions. Error =", err)
	}

	latestVersion := v.Versions[len(v.Versions)-1]

	latestMajorVersion, latestMinorVersion := getVersionsToCompare(latestVersion)
	selectedMajorVersion, selectedMinorVersion := getVersionsToCompare(selectedVersion)

	apiVersion := selectedVersion

	if latestMajorVersion < selectedMajorVersion {
		apiVersion = latestVersion
	} else if (latestMajorVersion == selectedMajorVersion) && (latestMinorVersion < selectedMinorVersion) {
		apiVersion = latestVersion
	}

	return apiVersion
}

func getVersionsToCompare(versionStr string) (int64, int64) {
	decimal := strings.Split(versionStr, ".")
	intPart, err := strconv.ParseInt(decimal[0], 10, 32)

	if err != nil {
		log.Fatalf("Couldn't convert string to int. Err =", err)
	}

	fracPart, err := strconv.ParseInt(decimal[1], 10, 16)

	if err != nil {
		log.Fatalf("Couldn't convert string to int. Err =", err)
	}

	return intPart, fracPart

}

func getAPITokenFromEnv() string {
	authToken, ok := os.LookupEnv("PUREFB_API")
	if !ok {
		log.Fatalln("No environment variable PUREFB_API found")
	}
	return authToken
}

// Retrieve an x-auth-token to put in the header of subsequent requests
func (fbClient *FlashbladeClient) refreshXAuthToken() {
	apiToken := getAPITokenFromEnv()

	url := fbClient.urlForEndpoint("login")
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("api-token", apiToken)

	resp, err := fbClient.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	xAuthToken = resp.Header.Get("x-auth-token")
	log.Println("Fetched x-auth-token for authenticating with FlashBlade")
}

func (fbClient *FlashbladeClient) urlForEndpoint(endpoint string) string {
	u, _ := url.Parse(fmt.Sprintf("https://%s/api", fbClient.Host))
	u.Path = path.Join(u.Path, fbClient.ApiVersion, endpoint)
	return u.String()
}

func (fbClient *FlashbladeClient) GetJSON(endpoint string, params map[string]string, target interface{}) error {
	resp := fbClient.Get(endpoint, params)

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalf("Cannot close response body: %v", err)
		}
	}()

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		log.Fatalf("Cannot parse response: %s", err)
	}

	return nil
}

func (fbClient *FlashbladeClient) Get(endpoint string, params map[string]string) *http.Response {
	url := fbClient.urlForEndpoint(endpoint)
	resp := fbClient.makeAuthedGetRequest(url, params)

	if resp.StatusCode == http.StatusForbidden {
		log.Printf("HTTP Status 403: %v - refreshing the token and retrying", resp.Request.URL.String())
		fbClient.refreshXAuthToken()
		resp = fbClient.makeAuthedGetRequest(url, params)
		if resp.StatusCode == http.StatusForbidden {
			log.Fatalln("Couldn't authenticate with the given token")
		}
	} else if resp.StatusCode != 200 {
		log.Printf("Failed GET with status %d for %v", resp.StatusCode, resp.Request.URL.String())
	}

	return resp
}

func (fbClient *FlashbladeClient) makeAuthedGetRequest(url string, params map[string]string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	query := req.URL.Query()
	for k, v := range params {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	req.Header.Add("x-auth-token", xAuthToken)
	log.Printf("GET %v", req.URL.String())
	resp, err := fbClient.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

// Useful for debugging the JSON responses
func printResponseBody(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(body)
}
