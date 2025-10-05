package commands

var Registery = make(map[string]func(args ...string))

func Register(name string, handler func(args ...string)) {
	Registery[name] = handler
}
