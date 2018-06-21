package api

type Greeter struct{}

func (Greeter) Greet(name string) string {
	return "Hello " + name
}
