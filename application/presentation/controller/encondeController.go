package controller

import (
	"cipher/application/domain/entities"
	"cipher/application/domain/usecases"
	"context"
	"fmt"
)

type EncodeRequestBody struct {
	Data   string
	Schema string
}

type IEnconde32ControllerRequest interface {
	Handle(ctx context.Context, request EncodeRequestBody) (response Response)
}

type Encode32ControllerRequest struct {
	enconde32 usecases.IEnconde32UseCase
}

func (c Encode32ControllerRequest) Handle(ctx context.Context, request EncodeRequestBody) Response {

	var response entities.Encode
	var err error

	if request.Data == "" {
		return Response{Err: fmt.Errorf("invalid data")}
	}
	switch request.Schema {
	case "32":
		response, err = c.enconde32.Execute(ctx, usecases.Enconde32Request{Data: request.Data, Schema: request.Schema})
	default:
		return Response{Err: fmt.Errorf("invalid schema")}
	}

	return Response{Data: response.ToByte(), Err: err, StatusCode: 200}
}

func NewEncode32Controller(enconde32 usecases.IEnconde32UseCase) IEnconde32ControllerRequest {
	return Encode32ControllerRequest{enconde32: enconde32}
}
