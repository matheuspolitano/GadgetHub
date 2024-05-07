package chat

type Messager interface {
	Send(msg string, to string) error
}
