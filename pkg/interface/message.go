package message

type Message interface {
	GetMessage() string
	GetProps() map[string]interface{}
}
