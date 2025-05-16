package google

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"translate"
)

type Provider struct {
	client *http.Client

	host string
	acq  *tokenAcquirer
}

func New() *Provider {
	transport := &http.Transport{}
	// Skip verifies the server's certificate chain and host name.
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // skip verify
	host := defaultHost

	newUserAgentTransport(transport, defaultUserAgent)

	client := &http.Client{Transport: transport}

	acq := newTokenAcquirer(host, client)

	return &Provider{client: client, host: host, acq: acq}
}

func (p *Provider) Translate(ctx context.Context, from, to translate.Lang, text string) (string, error) {
	token, err := p.acq.do(ctx, text)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s/translate_a/single", p.host)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("client", "gtx")
	q.Add("sl", string(from))
	q.Add("tl", string(to))
	q.Add("hl", string(to))
	q.Add("tk", token)
	q.Add("q", text)
	q.Add("dt", "t")
	q.Add("dt", "bd")
	q.Add("dj", "1")
	q.Add("source", "popup")
	req.URL.RawQuery = q.Encode()

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
	case http.StatusBadRequest:
		return "", translate.ErrUnsupportedLang
	default:
		return "", fmt.Errorf("%w: %d", translate.ErrStatusCode, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var out output
	if err := json.Unmarshal(body, &out); err != nil {
		return "", err
	}

	translated := ""
	for _, s := range out.Sentences {
		translated += s.Trans
	}
	return translated, nil
}
