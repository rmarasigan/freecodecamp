package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// This is to unmarshal the data
func ParseBody(r *http.Request, x interface{}) {
	// Reading the body
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
