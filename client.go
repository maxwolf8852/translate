package translate

import "context"

type Client struct {
	providers  []Provider
	skipErrors bool
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
	for i := range c.providers {
		out, err := c.providers[i].Translate(ctx, from, to, text)
		if err != nil && c.skipErrors && i != len(c.providers)-1 {
			continue
		}
		return out, err
	}

	return "", ErrNoProviders
}
