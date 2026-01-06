package main

import (
	grpcclient "order-service-platform/service/api-gateway/grpc"
	"order-service-platform/service/api-gateway/router"
	"order-service-platform/service/api-gateway/router/middleware"

	"github.com/elastic/go-elasticsearch/v8"
)

func init() {
	// kafka.InitKafka()
	grpcclient.Init()
}

//docker-compose -f docker-compose-tool.yml -p tool up -d
//docker-compose -f docker-compose.yml up -d

func main() {
	// 初始化 ES client
	esClient, _ := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})

	r := router.CreateRouter()
	middleware.InitBulkIndexer(esClient) // 初始化 Indexer
	r.Use(middleware.ESLogger())
	router.SetupRouter(r)

	r.Run("0.0.0.0:8080")
}
