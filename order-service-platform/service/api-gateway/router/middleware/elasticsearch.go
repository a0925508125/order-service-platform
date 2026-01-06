package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/gin-gonic/gin"
)

// 全域 BulkIndexer
var bulkIndexer esutil.BulkIndexer

func InitBulkIndexer(esClient *elasticsearch.Client) {
	var err error
	bulkIndexer, err = esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         "http-logs",
		Client:        esClient,
		NumWorkers:    4,               // 同時寫入的執行緒數（配合 CPU 核心數）
		FlushBytes:    5e+6,            // 達到 5MB 時送出一次
		FlushInterval: 5 * time.Second, // 每 5 秒強制送出一次
	})
	if err != nil {
		fmt.Printf("Error creating the indexer: %s", err)
	}
}

func ESLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		var userID int64
		if v, exists := c.Get("user_id"); exists {
			if id, ok := v.(int64); ok {
				userID = id
			}
		}

		params := make(map[string]interface{})
		// GET 或 URL Query
		for k, v := range c.Request.URL.Query() {
			if len(v) == 1 {
				params[k] = v[0]
			} else {
				params[k] = v
			}
		}

		// 讀取 Body
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Next()

		// 準備 Log 物件
		logEntry := HTTPLog{
			Path:      c.FullPath(),
			Method:    c.Request.Method,
			Status:    c.Writer.Status(),
			Latency:   time.Since(start).Milliseconds(),
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			UserID:    userID,
			Ts:        time.Now().Unix(),
			Params:    params,
		}

		data, _ := json.Marshal(logEntry)

		// 使用 BulkIndexer 添加任務，而非直接發送請求
		err := bulkIndexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "index",
				Body:   bytes.NewReader(data),
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						fmt.Printf("ERROR: %s\n", err)
					} else {
						fmt.Printf("ERROR: %s: %s\n", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			fmt.Printf("Unexpected error: %s", err)
		}
	}
}
