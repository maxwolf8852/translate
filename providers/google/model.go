package google

type output struct {
	Sentences []sentence `json:"sentences"`
}

type sentence struct {
	Trans   string `json:"trans"`
	Orig    string `json:"orig"`
	Backend int    `json:"backend"`
}
