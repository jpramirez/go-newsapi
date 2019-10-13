# Go-newsapi


Small client for newsapi.org API.

Example of implementation in cmd/NewServer 

Importing


```golang
var parameters models.NewsQueryParameter
parameters.Category = "technology"
parameters.PageSize = 10
parameters.Language = "en"

ret, err := _feed.GetFeed(constants.TopHeadLinesEndpoint, parameters)
```

Listing all Sources 

```golang
//Testing Sources
var sourcePrameters models.NewsQueryParameter
sourcePrameters.Category = "technology"
retSources, err := _feed.GetSources(sourcePrameters)

fmt.Println(retSources.Status)
fmt.Println(retSources.Sources)
```
