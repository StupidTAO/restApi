package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "http://127.0.0.1:8080/api/accounts_rank/100"
	number := 5000
	var res *http.Response
	var err error

	successCount := 0
	failCount := 0

	start := time.Now()
	for i := 0; i < number; i++ {
		res, err = http.Get(url)
		if err != nil {
			fmt.Println("occur error: ", err)
			return
		}
		if res.StatusCode != 200 {
			failCount++
			fmt.Println("[Warning]access page exception")
		} else {
			successCount++
		}

	 	fmt.Println(i, res.Close)
		res.Body.Close()

	}

	end := time.Now()

	spend := end.UnixNano() - start.UnixNano()
	fmt.Printf("time spend: %d\n", spend)
	fmt.Printf("success count: %d, failed count: %d\n", successCount, failCount)

}