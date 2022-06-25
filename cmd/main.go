package main

import (
	"InterfaceMastering/Logic"
	"InterfaceMastering/Logic/loogging"
)

func main() {

	hello := Logic.NewHelloLogic()
	hello = loogging.ForLoggingHello(hello)

	hello.Hello("Hello there")
}
