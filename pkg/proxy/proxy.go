package proxy

type MLModel interface {
	GetMLName() string
	GetMLGroup() string
	GetEndPoint() string
	SendRequest(string) (string, error)
}
