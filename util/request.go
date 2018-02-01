package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func noRedirect(req *http.Request, via []*http.Request) error {
	return errors.New("redirects are disabled")
}

func validateResponse(response http.Response) error {
	statusCode := response.StatusCode
	if statusCode != 200 {
		return fmt.Errorf("unexpected response status, expected 200, got %d", statusCode)
	}

	contentType := response.Header.Get("content-type")
	if !strings.HasPrefix(contentType, "application/json") {
		return fmt.Errorf("unexpected response content-type, expected application/json, got %s", contentType)
	}

	return nil
}

// Request ...
func Request(uri string, bearer *string) ([]byte, error) {
	client := &http.Client{
		Timeout:       time.Second * 10,
		CheckRedirect: noRedirect,
	}

	req, _ := http.NewRequest("GET", uri, nil)

	if bearer != nil {
		parts := []string{
			"Bearer",
			*bearer,
		}
		value := strings.Join(parts, " ")
		req.Header.Set("Authorization", value)
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = validateResponse(*response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}
