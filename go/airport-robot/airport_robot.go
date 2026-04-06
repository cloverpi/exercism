package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct{}

func (l Italian) LanguageName() string     { return "Italian" }
func (l Italian) Greet(name string) string { return "Ciao " + name + "!" }

type Portuguese struct{}

func (l Portuguese) LanguageName() string     { return "Portuguese" }
func (l Portuguese) Greet(name string) string { return "Olá " + name + "!" }

func SayHello(name string, greet Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greet.LanguageName(), greet.Greet(name))
}
