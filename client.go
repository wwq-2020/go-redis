package redis

// Client Client
type Client struct {
}

// Conf Conf
type Conf struct {
	Addr    string
	Network string
}

// New New
func New(c *Conf) *Client {
	return &Client{}
}
