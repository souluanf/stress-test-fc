package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Client struct {
	concurrency int
	client      *http.Client
}

func NewClient(concurrency int) *Client {
	return &Client{
		concurrency: concurrency,
		client:      &http.Client{},
	}
}

func (c *Client) LoadTest(url string, totalRequests int) (map[int]int, error) {
	var wg sync.WaitGroup
	results := make(map[int]int)
	ch := make(chan int, c.concurrency)
	mu := sync.Mutex{}
	for i := 0; i < 600; i++ {
		results[i] = 0
	}
	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		ch <- 1
		go func() {
			defer func() {
				wg.Done()
				<-ch
			}()
			resp, err := c.client.Get(url)
			if err != nil {
				fmt.Println("Erro ao fazer request:", err)
				return
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					fmt.Println("Erro ao fechar body:", err)
				}
			}(resp.Body)
			mu.Lock()
			results[resp.StatusCode]++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return results, nil
}
