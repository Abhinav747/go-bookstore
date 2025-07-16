// used for marshalling and unmarshalling
// we we recevie a request then it is in the form of json so we have to convert into go struct
package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// help in parsing in create book function..... when user send the data in form of request to create a book.....the data is in json format
func ParseBody(r *http.Request, x any){//interface{} can be replaced by any... Under the hood, type any = interface{},,Empty interface - can hold ANY type
	if body, err := io.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal(body,x); err != nil{
			return
		}

	}
}

//Read the JSON data from an HTTP request body
//Parse (unmarshal) that JSON into a Go data structure
//Handle any errors that might occur during this process