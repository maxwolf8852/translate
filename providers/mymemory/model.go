package mymemory

type Model struct {
	ResponseData struct {
		TranslatedText string `json:"translatedText"`
		Match          int    `json:"match"`
	} `json:"responseData"`
	QuotaFinished   bool   `json:"quotaFinished"`
	MtLangSupported any    `json:"mtLangSupported"`
	ResponseDetails string `json:"responseDetails"`
	ResponseStatus  int    `json:"responseStatus"`
	ResponderID     any    `json:"responderId"`
	ExceptionCode   any    `json:"exception_code"`
	Matches         []struct {
		ID             string `json:"id"`
		Segment        string `json:"segment"`
		Translation    string `json:"translation"`
		Source         string `json:"source"`
		Target         string `json:"target"`
		Quality        int    `json:"quality"`
		Reference      any    `json:"reference"`
		UsageCount     int    `json:"usage-count"`
		Subject        string `json:"subject"`
		CreatedBy      string `json:"created-by"`
		LastUpdatedBy  string `json:"last-updated-by"`
		CreateDate     string `json:"create-date"`
		LastUpdateDate string `json:"last-update-date"`
		Match          int    `json:"match"`
		Penalty        int    `json:"penalty"`
	} `json:"matches"`
}
