package utils

// This file is used to be able to read user requests as they are in json form

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PasreBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
}