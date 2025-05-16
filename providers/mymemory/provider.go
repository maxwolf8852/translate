package mymemory

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"translate"
)

type Provider struct {
	client *http.Client
}

func New() *Provider {
	return &Provider{client: &http.Client{}}
}

func (p *Provider) Translate(ctx context.Context, from, to translate.Lang, text string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, makeRequest(from, to, text), nil)
	if err != nil {
		return "", err
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// pass
	case http.StatusTooManyRequests:
		return "", translate.ErrTooManyRequests
	default:
		return "", fmt.Errorf("%w: %d", translate.ErrStatusCode, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var out Model
	if err := json.Unmarshal(body, &out); err != nil {
		return "", err
	}

	if out.ResponseStatus != 200 {
		return "", fmt.Errorf("%w: %d", translate.ErrStatusCode, out.ResponseStatus)
	}

	if out.ResponseData.TranslatedText != "" {
		return out.ResponseData.TranslatedText, nil
	}

	for _, match := range out.Matches {
		return match.Translation, nil
	}

	return "", translate.ErrNotFound
}

func makeRequest(from, to translate.Lang, text string) string {
	return fmt.Sprintf("%s?q=%s&langpair=%s|%s", baseUrl, url.QueryEscape(text), from, to)
}
