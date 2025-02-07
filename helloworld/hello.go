package helloworld

import "fmt"

type Greeter struct {
	Message string
}

func NewGreeter() *Greeter {
	return &Greeter{
		Message: "Hello, World 2!",
	}
}

func (g *Greeter) Greet() {
	fmt.Println(g.Message)
}
