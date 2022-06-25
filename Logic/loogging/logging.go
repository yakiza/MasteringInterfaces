package loogging

import (
	"InterfaceMastering"
	"fmt"
)

type LoggingHandler struct {
	Usecase InterfaceMastering.HeloUseCase
}

func (l LoggingHandler) Hello(x string) string {
	fmt.Println("Logging happend here")
	defer func() {
		fmt.Println("Goth in to the defer")
	}()

	return l.Usecase.Hello(x)
}

func ForLoggingHello(usecase InterfaceMastering.HeloUseCase) InterfaceMastering.HeloUseCase {
	return LoggingHandler{Usecase: usecase}
}
