package adapters

import "net/http"

type Response struct {
	Data       []byte
	Err        error
	StatusCode int
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
	response.WriteHeader(500) //build map of error
	response.Header().Add("Content-Type", "application/json")
	// response.Write(res.Data) process errpr
}
