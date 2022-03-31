package credentialValidator

import (
	"fmt"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

func (u *User) Validate(ID string, password string) error {
	fmt.Println("Estoy validando el usuario", ID, password)
	return nil
}
