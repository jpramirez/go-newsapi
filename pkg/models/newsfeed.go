package models

//NewsFeed struct type for V2 of apinews.org
type NewsFeed struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}
