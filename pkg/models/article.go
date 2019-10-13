package models

//Article a basic article struct
type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLoImage   string `json:"urlToImage"`
	PublishedAt string `json:"publishedat"`
	Content     string `json:"content"`
}
