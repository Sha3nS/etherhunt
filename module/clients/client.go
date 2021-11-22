package clients

type Client interface {
	Dial(url string) error
}


