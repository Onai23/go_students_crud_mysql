package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//set fungsi ParseBody()
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
