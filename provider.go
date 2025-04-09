package translate

import "context"

type Provider interface {
	Translate(ctx context.Context, from, to Lang, text string) (string, error)
}
