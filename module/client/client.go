package client

type Client interface {
	Dial(url string) error

}


