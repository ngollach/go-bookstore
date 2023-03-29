package utils

//So when we start our book controller we need data that is unmarshall
/*
As we receive data from user it will be in Json format
So we need to UnMarshall it to able to use it in our controller
SO here we will unmarshall the json from the request
*/

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {

	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}

}
