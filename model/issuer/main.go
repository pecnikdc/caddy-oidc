package issuer

import (
	"encoding/json"

	"../../util"
)

// Issuer ...
type Issuer struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	Issuer                string `json:"issuer"`
	TokenEndpoint         string `json:"token_endpoint"`
}

// Discover ...
func Discover(wellKnown string) (*Issuer, error) {
	body, err := util.Request(wellKnown, nil)
	if err != nil {
		return nil, err
	}

	var issuer Issuer
	err = json.Unmarshal(body, &issuer)

	if err != nil {
		return nil, err
	}

	return &issuer, nil
}
