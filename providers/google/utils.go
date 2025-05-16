package google

import (
	"crypto/rand"
	"math/big"
	"net/http"
)

type customTransport struct {
	http.RoundTripper
	userAgent string
}

func newUserAgentTransport(trans http.RoundTripper, userAgent string) http.RoundTripper {
	return &customTransport{trans, userAgent}
}

func (ct *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", ct.userAgent)
	return ct.RoundTripper.RoundTrip(req)
}

func chooseRandom(elements []string) (string, error) {
	// Generate a random index within the range of the slice length
	max := big.NewInt(int64(len(elements)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	// Return the randomly selected element
	return elements[n.Int64()], nil
}
