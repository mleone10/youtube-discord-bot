package main

import (
	"encoding/json"
	"net/http"
)

func parseResponse(r *http.Response, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
