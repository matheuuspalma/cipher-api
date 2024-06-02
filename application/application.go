package application

import (
	"cipher/application/domain/usecases"
	"cipher/application/presentation/controller"
)

var EncodeController controller.IEnconde32ControllerRequest

func Init() {

	EncodeController = controller.NewEncode32Controller(usecases.NewEncode32UseCase())

}
