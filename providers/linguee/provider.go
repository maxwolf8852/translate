package linguee

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/maxwolf8852/translate"

	"golang.org/x/net/html"
)

type Provider struct {
	client *http.Client
}

func New() *Provider {
	return &Provider{client: &http.Client{}}
}

func (p *Provider) Translate(ctx context.Context, from, to translate.Lang, text string) (string, error) {
	words := splitSentence(text)
	translated := make([]string, 0, len(words))
	for i := range words {
		w, err := p.translateWord(ctx, from, to, words[i])
		if err != nil {
			return "", err
		}

		translated = append(translated, w)
	}

	return joinWords(translated), nil
}

func joinWords(words []string) string {
	var sb strings.Builder
	for i := range words {
		if i != 0 && wordRegex.MatchString(words[i]) {
			sb.WriteString(" ")
		}
		sb.WriteString(words[i])
	}
	return sb.String()
}

func (p *Provider) translateWord(ctx context.Context, from, to translate.Lang, word string) (string, error) {
	if len(word) > 50 {
		return "", fmt.Errorf("%w: word exceeds max length", translate.ErrInvalidInput)
	}

	if word == "" {
		return "", nil
	}

	source, ok := code2lang[from]
	if !ok {
		return "", translate.ErrUnsupportedLang
	}

	target, ok := code2lang[to]
	if !ok {
		return "", translate.ErrUnsupportedLang
	}

	url := fmt.Sprintf("%[1]s/%[2]s-%[3]s/search/?source=%[2]s&query=%[4]s", baseURL, source, target, url.QueryEscape(word))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
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

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	// Extract the translation results from the HTML
	elements := extractElements(doc)
	if len(elements) == 0 {
		return word, nil
	}

	var translations []string
	for _, el := range elements {
		translations = append(translations, el)
	}

	if len(translations) == 0 {
		return word, nil
	}

	return strings.TrimSpace(translations[0]), nil
}

func splitSentence(sentence string) []string {
	return sentenceSplitRe.FindAllString(sentence, -1)
}

func extractElements(n *html.Node) []string {
	var elements []string
	var parseNode func(*html.Node)
	parseNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "dictLink featured" {
					// Extract text
					var text string
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.TextNode {
							text = text + c.Data
						}
					}
					elements = append(elements, strings.TrimSpace(text))
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseNode(c)
		}
	}
	parseNode(n)
	return elements
}
