package usecases

import (
	"cipher/application/domain/entities"
	"cipher/application/domain/externals"
	"cipher/application/infrastructure"
	"context"
)

type Enconde32Request externals.Encode32request

type IEnconde32UseCase interface {
	Execute(ctx context.Context, data Enconde32Request) (entities.Encode, error)
}

type Enconde32UseCase struct {
	encode infrastructure.Encode32
}

func (u Enconde32UseCase) Execute(ctx context.Context, data Enconde32Request) (entities.Encode, error) {

	response, err := u.encode.Encode(data.Data)

	return entities.Encode{EncodedData: response}, err
}

func NewEncode32UseCase() IEnconde32UseCase {
	return Enconde32UseCase{}
}
