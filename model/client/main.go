package client

import (
	"encoding/json"

	"../../util"
)

// Client ...
type Client struct {
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method"`
	ClientID                string `json:"client_id"`
	ClientSecret            string `json:"client_secret"`
}

// Read ...
func Read(registrationClientURI string, registrationAccessToken string) (*Client, error) {
	body, err := util.Request(registrationClientURI, &registrationAccessToken)
	if err != nil {
		return nil, err
	}

	var client Client
	err = json.Unmarshal(body, &client)

	if err != nil {
		return nil, err
	}

	return &client, nil
}
