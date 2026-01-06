package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	totalRequests = 10
	concurrency   = 200
	ip            = "localhost"
	postURL       = fmt.Sprintf("http://%s:8080/v1/order", ip)
	getURL        = postURL // GET 也用同一個 API
)

var client *http.Client

func init() {
	transport := &http.Transport{
		MaxIdleConns:        2000,
		MaxIdleConnsPerHost: 1000,
		IdleConnTimeout:     90 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client = &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
}

// POST 請求方法
func doPost(url string, orderID string) error {
	payload := []byte(fmt.Sprintf(`{"order_id":"%s","user_id":123,"quantity":1}`, orderID))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("[POST %s] Status: %d, Body: %s\n", orderID, resp.StatusCode, string(body))
	return nil
}

// GET 請求方法
func doGet(url string, orderID string) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("[GET %s] Status: %d, Body: %s\n", orderID, resp.StatusCode, string(body))
	return nil
}

func main() {

	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	wg.Add(totalRequests)

	for i := 0; i < totalRequests; i++ {
		sem <- struct{}{}
		go func(idx int) {
			defer wg.Done()
			defer func() { <-sem }()

			orderID := fmt.Sprintf("order-%d", idx)

			// 隨機 POST 或 GET（示例：偶數 POST，奇數 GET）
			if err := doPost(postURL, orderID); err != nil {
				fmt.Println("POST error:", err)
			}
			// if err := doGet(getURL, orderID); err != nil {
			// 	fmt.Println("GET error:", err)
			// }
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests finished!")
}
