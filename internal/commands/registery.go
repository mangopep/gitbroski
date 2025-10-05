package commands

var Registery = make(map[string]func())

func Register(name string, handler func()) {
	Registery[name] = handler
}
