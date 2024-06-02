package adapters

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data       []byte
	Err        error
	StatusCode int
}

type HttpError struct {
	Err     error
	Message string
}

func ResponseAdapt(response http.ResponseWriter, res Response) {

	if res.Err != nil {
		errorProcess(response, res)
		return
	}
	response.WriteHeader(res.StatusCode)
	response.Header().Add("Content-Type", "application/json")
	response.Write(res.Data)
}

func errorProcess(response http.ResponseWriter, res Response) {
	err, _ := json.Marshal(HttpError{Err: res.Err, Message: res.Err.Error()})
	response.WriteHeader(res.StatusCode) //build map of error
	response.Header().Add("Content-Type", "application/json")
	response.Write(err)
	// response.Write(res.Data) process errpr
}
