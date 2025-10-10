package commands

var Registry = make(map[string]func(args ...string))

func Register(name string, handler func(args ...string)) {
	Registry[name] = handler
}
