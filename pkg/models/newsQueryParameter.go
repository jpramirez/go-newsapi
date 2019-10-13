package models

//NewsQueryParameter Format
type NewsQueryParameter struct {
	Country        string   `json:"country"`
	Category       string   `json:"category"`
	Sources        []Source `json:"sources"`
	Query          string   `json:"query"`
	QueryInTitle   string   `json:"queryintitle"`
	PageSize       int      `json:"pagesize"`
	Page           int      `json:"page"`
	APIKey         string   `json:"apikey"`
	Domains        []string `json:"domains"`
	ExcludeDomains []string `json:"excludedomains"`
	FromDate       string   `json:"from"`
	ToDate         string   `json:"to"`
	Language       string   `json:"language"`
	SortBy         string   `json:"sortby"`
}
