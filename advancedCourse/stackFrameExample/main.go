package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	code := make(chan int) //Указание канала для горутины
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHTTPRequest(code) // передача канала в конкретную горутину
			wg.Done()
		}()

	}
	go func() {
		wg.Wait()
		close(code)
	}()
	for resp := range code {

		// resp:= <-code//Запись значения из канала горутины куда-то, если ничего не придёт, то рутина будет намертво висеть
		fmt.Println("Status Code:", resp)
	}
}

// Свой клиент, а не http.Get: у http.DefaultClient таймаут нулевой,
// то есть ожидание ответа ничем не ограничено.
var client = &http.Client{Timeout: 10 * time.Second}

func getHTTPRequest(codeCh chan int) /*(int, error)*/ { // Тут канал как аргумент, в него потом записывается значение статус кода
	const baseURL = "http://google.com/"
	base, err := url.Parse(baseURL)

	resp, err := client.Get(base.String())
	if err != nil {
		//return 0, fmt.Errorf("google: обращение по адресу %s: %w", base.String(), err)
		fmt.Printf("google: обращение по адресу %s: %w", base.String(), err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		//return 0, fmt.Errorf("google: статус %d: %s",resp.StatusCode)
		fmt.Printf("google: статус %d: %s", resp.StatusCode)
	}
	codeCh <- resp.StatusCode // Запись значения статус кода в канал горутины

	//return resp.StatusCode, nil
}
