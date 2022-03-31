package main

import (
	"fmt"
	"github.com/nelsongp/testpatterns/credentialValidator"
	"github.com/nelsongp/testpatterns/logs"
)

func main() {
	//llammamos el decorator
	u := &credentialValidator.User{Id: "Nel", Password: "1234"}

	val := &logs.PrintLog{
		Val: u,
	}
	fmt.Println(val.Validate("asdf", "1234"))
}
