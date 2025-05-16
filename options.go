package translate

type Option func(*Client) error

func WithProvider(p Provider) Option {
	return func(client *Client) error {
		client.providers = append(client.providers, p)
		return nil
	}
}

func WithSkipErrors() Option {
	return func(client *Client) error {
		client.skipErrors = true
		return nil
	}
}
