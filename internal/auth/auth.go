package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API key from the http header
// Example :
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth headers")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of api key")
	}
	return vals[1], nil
}
