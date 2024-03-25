package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

type IpGeoApi struct {
	ApiKey string
}

func NewGeoApi(apiKeyFileName string) (IpGeoApi, error) {
	file, err := os.ReadFile(apiKeyFileName)
	if err != nil {
		return IpGeoApi{}, err
	}

	apiKey := string(file)
	return IpGeoApi{apiKey}, err
}

func (api IpGeoApi) LookupIp(ip string) (IpGeolocation, error) {
	/*
	* Ip validation
	 */
	isValidIpv4Ip, err := regexp.MatchString("^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$", ip)
	if err != nil {
		return IpGeolocation{}, err
	}

	isValidIpv6Ip, err := regexp.MatchString("(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))", ip)

	if err != nil {
		return IpGeolocation{}, err
	}

	if isValidIpv4Ip == false && isValidIpv6Ip == false {
		return IpGeolocation{}, fmt.Errorf("invalid ip")
	}

	/*
	*	Http/tls request
	 */

	// Setup the http client to use tls
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	client := &http.Client{Transport: transport}

	url := fmt.Sprintf("https://api.ipgeolocation.io/ipgeo?apiKey=%s&ip=%s", api.ApiKey, ip)

	// Make the request
	response, err := client.Get(url)
	if err != nil {
		return IpGeolocation{}, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return IpGeolocation{}, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return IpGeolocation{}, fmt.Errorf("%d %s", response.StatusCode, string(body))
	}

	// unmarshal
	var ipGeoLocation IpGeolocation
	err = json.Unmarshal(body, &ipGeoLocation)
	if err != nil {
		return IpGeolocation{}, err
	}

	return ipGeoLocation, nil
}
