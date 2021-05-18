package message

type Message interface {
	GetMessage() string
	GetProperties() map[string]interface{}
}
