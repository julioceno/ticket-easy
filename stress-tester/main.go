package main

import (
	"bytes"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			createRequest(i)
		}()
	}

	wg.Wait()
}

func createRequest(index int) {
	var jsonStr = []byte(`{
			"eventId": "66d37f915deb0915f95b3552",
			"userId": "66d37ec1e08c07d3f15dee4c"
		}`)

	client := &http.Client{}
	for i := 0; i <= 10000; i++ {
		req, _ := http.NewRequest("POST", "http://localhost:8082/tickets", bytes.NewBuffer(jsonStr))
		req.Header.Set("x-api-key", "secret")
		req.Header.Set("Content-Type", "application/json")
		client.Do(req)
	}
}
