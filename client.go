package translate

import "context"

type Client struct {
	provider Provider
}

func New(options ...Option) (*Client, error) {
	c := &Client{}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) Translate(ctx context.Context, from, to Lang, text string) (string, error) {
	if c.provider == nil {
		return "", ErrUnknownProvider
	}
	return c.provider.Translate(ctx, from, to, text)
}
