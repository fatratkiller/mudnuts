package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const geckoEndpoint = "https://api.xxxx"

var geckoRequestHeader = map[string]string{
	"Accept": "application/json;version=20230302",
}

// GetTokens fetches the recently updated tokens from API.
func GetTokens() ([]Token, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", geckoEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for key, value := range geckoRequestHeader {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching tokens: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return tokenResp.Data, nil
}
