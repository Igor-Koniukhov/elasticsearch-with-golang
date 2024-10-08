package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"es-with-go-sample/domain"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func BestKeanuActionMovies(ctx context.Context) {

	client := ctx.Value(domain.ClientKey).(*elasticsearch.Client)

	var searchBuffer bytes.Buffer
	search := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]string{
						"actors.en": "keanu reeves",
					},
				},
				"filter": []map[string]interface{}{
					{
						"term": map[string]string{
							"genres.keyword": "Action",
						},
					},
					{
						"range": map[string]interface{}{
							"rating": map[string]float64{
								"gte": 7.0,
							},
						},
					},
					{
						"range": map[string]interface{}{
							"year": map[string]int{
								"gte": 1995,
								"lte": 2005,
							},
						},
					},
				},
			},
		},
	}
	err := json.NewEncoder(&searchBuffer).Encode(search)
	if err != nil {
		panic(err)
	}

	response, err := client.Search(
		client.Search.WithContext(ctx),
		client.Search.WithIndex("movies"),
		client.Search.WithBody(&searchBuffer),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var SearchResponse = domain.SearchResponse{}
	err = json.NewDecoder(response.Body).Decode(&SearchResponse)
	if err != nil {
		panic(err)
	}

	if SearchResponse.Hits.Total.Value > 0 {
		var movieTitles []string
		for _, movieTitle := range SearchResponse.Hits.Hits {
			movieTitles = append(movieTitles, movieTitle.Source.Title)
		}
		fmt.Printf("✅ Best action movies from Keanu: [%s] \n", strings.Join(movieTitles, ", "))
	}
}
