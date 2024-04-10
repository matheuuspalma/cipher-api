package adapters

import (
	"encoding/json"
	"net/http"
)

type encodeHeader struct {
	Name string `json:"name"`
}

type encodeBody struct {
	Data   string `json:"data"`
	Scheme string `json:"scheme"`
}

func Encode(output http.ResponseWriter, input *http.Request) {

	req := Adapter{
		Body:   &encodeBody{},
		Header: &encodeHeader{},
	}

	err := req.RequestAdapt(input)

	res := Response{Data: []byte{}, StatusCode: 201, Err: err}

	data, _ := json.Marshal(res)
	ResponseAdapt(output, Response{Data: data, StatusCode: 201, Err: err})
}
