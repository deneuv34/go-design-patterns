package main

import (
	"encoding/json"
	"fmt"
)

// Message Json struct
type Message struct {
	Body   []byte
	Format string
}

// MessageBuilder interface
type MessageBuilder interface {
	SetHeader(header string)
	SetBody(body string)
	SetFooter(footer string)
	BuildMessage() (*Message, error)
}

// MessageStruct type
type MessageStruct struct {
	Header string
	Body   string
	Footer string
}

// SetHeader Function
func (s *MessageStruct) SetHeader(header string) {
	s.Header = header
}

// SetBody function
func (s *MessageStruct) SetBody(body string) {
	s.Body = body
}

// SetFooter function
func (s *MessageStruct) SetFooter(footer string) {
	s.Footer = footer
}

// BuildMessage function to build the Message Struct into Message struct
func (s *MessageStruct) BuildMessage() (*Message, error) {
	m := make(map[string]string)
	m["header"] = s.Header
	m["body"] = s.Body
	m["footer"] = s.Footer

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return &Message{Body: data, Format: "JSON"}, nil
}

// Sender Struct
type Sender struct{}

// SendMessage Function
func (m *Sender) SendMessage(builder MessageBuilder) (*Message, error) {
	builder.SetHeader("Dear Rangga")
	builder.SetBody("Hello, How are you")
	builder.SetFooter("Regards, Adhitya")

	return builder.BuildMessage()
}

func main() {
	sender := &Sender{}

	jsonMessage, err := sender.SendMessage(&MessageStruct{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonMessage.Body)) // {"body":"Hello, How are you","footer":"Regards, Adhitya","header":"Dear Rangga"}
}
