package linguee

import "regexp"

const baseURL = "https://www.linguee.com"

var (
	foundRe         = regexp.MustCompile(`<div class="result_list">.*?<div class="target">(.*?)</div>.*?</div>`)
	sentenceSplitRe = regexp.MustCompile(`\p{L}+|[^\p{L}\s]`)

	wordRegex  = regexp.MustCompile(`^[\p{L}\p{N}_]+$`)
	delimRegex = regexp.MustCompile(`^[^\p{L}\p{N}_]+$`)
)
