package object

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment environment
type Environment struct {
	store map[string]Object
}

// Get gets an object from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set sets an object in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
