package common

type DocumentTypes struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Status       bool   `json:"status"`
}
