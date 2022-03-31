package decorator

//Decorator interface que deben implementar las estructuras que se van a decorar
type Decorator interface {
	Validate(ID string, password string)
}
