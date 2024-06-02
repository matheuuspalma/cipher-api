package adapters

import (
	"cipher/application"
	"cipher/application/presentation/controller"
	"encoding/json"
	"net/http"
)

type encodeHeader struct {
	Name string `json:"name"`
}

type encodeBody struct {
	Data   string `json:"data"`
	Schema string `json:"schema"`
}

func Encode(output http.ResponseWriter, input *http.Request) {

	ctx := input.Context()

	req := Adapter{
		Body:   &encodeBody{},
		Header: &encodeHeader{},
	}

	err := req.RequestAdapt(input)

	if err != nil {
		ResponseAdapt(output, Response{StatusCode: 400, Err: err}) // bad request
		return
	}

	var response controller.Response
	var controllerRequest controller.EncodeRequestBody

	bodyByte, err := json.Marshal(req.Body)

	if err != nil {
		ResponseAdapt(output, Response{StatusCode: 500, Err: err})
		return
	}

	json.Unmarshal(bodyByte, &controllerRequest)
	response = application.EncodeController.Handle(ctx, controllerRequest)

	res := Response{Data: response.Data, StatusCode: response.StatusCode, Err: response.Err}

	ResponseAdapt(output, res)
}
