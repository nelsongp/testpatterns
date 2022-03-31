package notifier

import (
	"fmt"
	"time"
)

type DB struct{}

func (d *DB) Notify(ID, password string) {
	writeInDataBase(ID, password)
}

func writeInDataBase(ID, password string) {
	fmt.Println("Voy a escribir en base de datos")
	fmt.Printf("ID: %v, Password: %v, Date: %v", ID, password, time.Now())
}
