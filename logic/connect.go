package logic

import (
	"context"
	"es-with-go-sample/domain"

	"github.com/elastic/go-elasticsearch/v8"
)

func ConnectWithElasticsearch(ctx context.Context) context.Context {
	NewClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})

	if err != nil {
		panic(err)
	}

	return context.WithValue(ctx, domain.ClientKey, NewClient)
}
