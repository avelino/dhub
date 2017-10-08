package repositories

type buildHistoryItem struct {
	Id        int    `json:"id"`
	Status    int    `json:"status"`
	TagName   string `json:"dockertag_name"`
	BuildCode string `json:"build_code"`
}

type buildHistory struct {
	Count int                `json:"count"`
	Itens []buildHistoryItem `json:"results"`
}
