package client

type Pinger interface {
	ForwardPing(payload map[string]interface{}) error
}
