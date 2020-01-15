package main

import "fmt"

type bot interface {
	getGreeting() string
}

type chatbot interface {
	bot
}

type engBot struct{}
type spBot struct{}

func print(b chatbot) {
	fmt.Println(b.getGreeting())
}

func (engBot) getGreeting() string {
	return "hello"
}

func (spBot) getGreeting() string {
	return "hola"
}

func main() {

	sb := spBot{}
	eb := engBot{}

	print(sb)
	print(eb)

}
