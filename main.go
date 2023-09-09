package main

type Message struct {
	Body string
}

func NewMessage(body string) Message {
	return Message{body}
}

func PrintMessage(m Message) {
	println(m.Body)
}
