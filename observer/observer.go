package observer

type Observer interface {
	Notify(ID string, password string)
}
