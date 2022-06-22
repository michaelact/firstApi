package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(req *http.Request, result interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&result)
	PanicIfError(err)
}

func WriteToResponseBody(res http.ResponseWriter, result interface{}) {
	encoder := json.NewEncoder(res)
	err = encoder.Encode(&result)
	PanicIfError(err)
}
