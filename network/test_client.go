package main

import (
	"fmt"
	"net/http"
)

func httpClient() {

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("resp.Body: %v\n", resp.Body)
	}
	defer resp.Body.Close()

}

func printLog() {
	fmt.Printf("\"233\": %v\n", "233")
}
