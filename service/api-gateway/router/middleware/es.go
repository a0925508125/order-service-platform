package middleware

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"time"

// 	es "github.com/elastic/go-elasticsearch/v8"
// 	"github.com/gin-gonic/gin"
// )

// var esLogChan = make(chan []byte, 10000)

// func StartESWriter(esClient *es.Client) {
// 	go func() {
// 		for b := range esLogChan {
// 			res, err := esClient.Index(
// 				"http-logs",
// 				bytes.NewReader(b),
// 			)
// 			if err != nil {
// 				fmt.Println("ES write error:", err)
// 				continue
// 			}
// 			res.Body.Close()
// 		}
// 	}()
// }

// func ESLogger(esClient *es.Client) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		start := time.Now()

// 		// 先讀取 body，避免 ShouldBindJSON 消耗掉
// 		var bodyBytes []byte
// 		if c.Request.Body != nil {
// 			bodyBytes, _ = io.ReadAll(c.Request.Body)
// 			// 重新放回去，讓後續 c.Next() 可以正常讀
// 			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
// 		}

// 		c.Next()
// 		latency := time.Since(start).Milliseconds()
// 		status := c.Writer.Status()

// 		var userID int64
// 		if v, exists := c.Get("user_id"); exists {
// 			if id, ok := v.(int64); ok {
// 				userID = id
// 			}
// 		}

// 		params := make(map[string]interface{})
// 		// GET 或 URL Query
// 		for k, v := range c.Request.URL.Query() {
// 			if len(v) == 1 {
// 				params[k] = v[0]
// 			} else {
// 				params[k] = v
// 			}
// 		}

// 		// POST / PUT JSON
// 		if len(bodyBytes) > 0 && (c.Request.Method == "POST" || c.Request.Method == "PUT") {
// 			var bodyParams map[string]interface{}
// 			if err := json.Unmarshal(bodyBytes, &bodyParams); err == nil {
// 				for k, v := range bodyParams {
// 					params[k] = v
// 				}
// 			} else {
// 				// 如果不是 JSON，就當作原始字串
// 				params["raw_body"] = string(bodyBytes)
// 			}
// 		}

// 		logEntry := HTTPLog{
// 			Path:      c.FullPath(),
// 			Method:    c.Request.Method,
// 			Status:    status,
// 			Latency:   latency,
// 			UserID:    userID,
// 			Ts:        time.Now().Unix(),
// 			Params:    params,
// 			Timestamp: time.Now().UTC().Format(time.RFC3339),
// 		}

// 		b, _ := json.Marshal(logEntry)

// 		go func() {
// 			res, err := esClient.Index(
// 				"http-logs",
// 				bytes.NewReader(b),
// 				esClient.Index.WithRefresh("true"),
// 			)
// 			if err != nil {
// 				fmt.Println("ES write error:", err)
// 				return
// 			}
// 			defer res.Body.Close()
// 			if res.IsError() {
// 				fmt.Println("ES response error:", res.String())
// 			} else {
// 				fmt.Println("ES write success:", string(b))
// 			}
// 		}()
// 	}
// }
