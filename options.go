package translate

type Option func(*Client) error

func WithProvider(p Provider) Option {
	return func(client *Client) error {
		client.provider = p
		return nil
	}
}
