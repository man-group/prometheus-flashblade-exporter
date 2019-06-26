//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
//
// Copyright (c) 2019 Hudson River Trading LLC
// All rights reserved.
//


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
)

var xAuthToken string

type FlashbladeClient struct {
	Host   string
	client *http.Client
}

func NewFlashbladeClient(host string, insecure bool) *FlashbladeClient {
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

	return &fb
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
	u.Path = path.Join(u.Path, endpoint)
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
