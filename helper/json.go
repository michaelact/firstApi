package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(req *http.Request, result interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(res http.ResponseWriter, result interface{}) {
	res.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(res)
	err := encoder.Encode(result)
	PanicIfError(err)
}
