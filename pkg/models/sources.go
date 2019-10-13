package models

//Sources struct type for V2 of apinews.org
type Sources struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}
