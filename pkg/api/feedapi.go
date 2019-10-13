package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	models "github.com/jpramirez/go-newsapi/pkg/models"
)

//NewsFeedAPI basic struct
type NewsFeedAPI struct {
	Config models.Config
}

//NewFeedAPI will create a new instance of NewsFeedAPI
func NewFeedAPI(config models.Config) (NewsFeedAPI, error) {
	var feedapi NewsFeedAPI
	log.Println("Starting News Reader Go")

	feedapi.Config = config

	// Stop the grpc verbose logging
	//grpclog.SetLogger(noplog)
	return feedapi, nil
}

//GetSources will get and list all the sources from newsapi
func (N *NewsFeedAPI) GetSources(Query models.NewsQueryParameter) (models.Sources, error) {
	var sourcesFeed models.Sources
	querystring, err := N.BuildQueryString("sources", Query)
	fmt.Println(querystring)
	resp, err := http.Get(N.Config.APIURL + querystring)
	if err != nil {
		log.Println("Error query the newsapi service sources")
		return sourcesFeed, err
	}
	json.NewDecoder(resp.Body).Decode(&sourcesFeed)
	return sourcesFeed, err
}

//GetFeed will get all the feeds from newsapi
func (N *NewsFeedAPI) GetFeed(endpoint string, Query models.NewsQueryParameter) (models.NewsFeed, error) {

	var newsfeed models.NewsFeed
	if endpoint == "sources" {
		log.Println("Error, please call getSources function instead")
	}
	querystring, err := N.BuildQueryString(endpoint, Query)
	fmt.Println(querystring)
	resp, err := http.Get(N.Config.APIURL + querystring)
	if err != nil {
		log.Println("Error query the newsapi service")
		return newsfeed, err
	}

	json.NewDecoder(resp.Body).Decode(&newsfeed)
	return newsfeed, err
}

//BuildQueryString is for building the parameters needed to execute a query to the api.
func (N *NewsFeedAPI) BuildQueryString(endpoint string, query models.NewsQueryParameter) (string, error) {
	var querystring string
	var err error

	switch endpoint {
	case "sources":
		if query.Category != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&category=" + query.Category
			} else {
				querystring = querystring + "?category=" + query.Category
			}
		}
		if query.Language != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&language=" + query.Language
			} else {
				querystring = querystring + "?language=" + query.Language
			}
		}

		if query.Country != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&country=" + query.Country
			} else {
				querystring = querystring + "?country=" + query.Country
			}
		}
	case "top-headlines":
		if query.Country != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&country=" + query.Country
			} else {
				querystring = querystring + "?country=" + query.Country
			}
		}
		if query.Category != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&category=" + query.Category
			} else {
				querystring = querystring + "?category=" + query.Category
			}
		}
		if query.Query != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&q=" + query.Query
			} else {
				querystring = querystring + "?q=" + query.Query
			}
		}
		if query.PageSize != 0 {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&pageSize=" + strconv.Itoa(query.PageSize)
			} else {
				querystring = querystring + "?q=" + strconv.Itoa(query.PageSize)
			}
		}
		if query.Page != 0 {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&page=" + strconv.Itoa(query.PageSize)
			} else {
				querystring = querystring + "?page=" + strconv.Itoa(query.PageSize)
			}
		}
	case "everything":

		if query.Query != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&q=" + query.Query
			} else {
				querystring = querystring + "?q=" + query.Query
			}
		}
		if query.QueryInTitle != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&qInTitle=" + query.QueryInTitle
			} else {
				querystring = querystring + "?qInTitle=" + query.QueryInTitle
			}
		}
		if len(query.Sources) != 0 {
			var qsources []string
			for _, source := range query.Sources {
				qsources = append(qsources, source.ID)
			}
			sSources := strings.Join(qsources, ",")
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&sources=" + sSources
			} else {
				querystring = querystring + "?sources=" + sSources
			}
		}

		if len(query.Domains) != 0 {
			sDomains := strings.Join(query.Domains, ",")
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&domains=" + sDomains
			} else {
				querystring = querystring + "?domains=" + sDomains
			}
		}
		if len(query.ExcludeDomains) != 0 {
			sDomains := strings.Join(query.ExcludeDomains, ",")
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&excludeDomains=" + sDomains
			} else {
				querystring = querystring + "?excludeDomains=" + sDomains
			}
		}

		if query.FromDate != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&from=" + query.FromDate
			} else {
				querystring = querystring + "?from=" + query.FromDate
			}
		}
		if query.ToDate != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&to=" + query.ToDate
			} else {
				querystring = querystring + "?to=" + query.ToDate
			}
		}
		if query.Language != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&language=" + query.Language
			} else {
				querystring = querystring + "?language=" + query.Language
			}
		}

		if query.Country != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&country=" + query.Country
			} else {
				querystring = querystring + "?country=" + query.Country
			}
		}
		if query.SortBy != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&sortBy=" + query.SortBy
			} else {
				querystring = querystring + "?sortBy=" + query.SortBy
			}
		}
		if query.Category != "" {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&category=" + query.Category
			} else {
				querystring = querystring + "?category=" + query.Category
			}
		}
		if query.PageSize != 0 {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&pageSize=" + strconv.Itoa(query.PageSize)
			} else {
				querystring = querystring + "?q=" + strconv.Itoa(query.PageSize)
			}
		}

		if query.Page != 0 {
			if strings.Contains(querystring, "?") {
				querystring = querystring + "&page=" + strconv.Itoa(query.PageSize)
			} else {
				querystring = querystring + "?page=" + strconv.Itoa(query.PageSize)
			}
		}
	}

	querystring = "/" + endpoint + querystring + "&apiKey=" + N.Config.APIKEY
	return querystring, err
}
