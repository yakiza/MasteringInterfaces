package Logic

import (
	"InterfaceMastering"
	"fmt"
)

type HelloHandler struct{}

func (h HelloHandler) Hello(x string) string {
	fmt.Println("THi is inside the logic")
	return x
}

func NewHelloLogic() InterfaceMastering.HeloUseCase {
	return HelloHandler{}
}
