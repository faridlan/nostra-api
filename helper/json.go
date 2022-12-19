package helper

import (
	"encoding/json"
	"net/http"
)

func RequestFromBody(request *http.Request, result any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response any) {
	writer.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
