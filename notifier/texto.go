package notifier

import (
	"fmt"
	"time"
)

type Text struct{}

func (t *Text) Notify(ID, password string) {
	printLogInText(ID, password)
}

func printLogInText(ID, password string) {
	fmt.Println("Voy a escribir en texto")
	fmt.Printf("ID: %v, Password: %v, Date: %v", ID, password, time.Now())
}
