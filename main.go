package main

import (
	"fmt"
	"net/http"
	"sync"
)

const N int = 1500

func main() {
	var wg sync.WaitGroup
	responses := make([]*http.Response, N)
	for i := 0; i < N; i++ {
		resp, err := http.Get(fmt.Sprintf("http://testuser:testpass@localhost:7999/testindex/testtype/test%d?stream=true", i))
		if err != nil {
			fmt.Println(i, err)
		} else {
			wg.Add(1)
		}
		responses[i] = resp
	}

	wg.Wait()
}
