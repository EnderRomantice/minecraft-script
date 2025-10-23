package runtime

type Environment struct {
	variables map[string]interface{}
}

func New() *Environment {
	return &Environment{
		variables: make(map[string]interface{}),
	}
}

func (e *Environment) Get(name string) (interface{}, bool) {
	val, ok := e.variables[name]
	return val, ok
}

func (e *Environment) Set(name string, value interface{}) {
	e.variables[name] = value
}