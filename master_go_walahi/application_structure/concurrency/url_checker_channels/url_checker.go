package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func urlChecker(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		//fmt.Println(err)
		s := fmt.Sprintf("%s is down!\n", url)
		s += fmt.Sprintf("Error:%v\n", err)
		c <- s
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s>>>>>Status Code:%d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			_ = err
			file := strings.Split(url, "//")[1]
			file += ".txt"

			s += fmt.Sprintf("writing response body to %s\n", file)

			fmt.Println(strings.Repeat("<>", 20))

			err = os.WriteFile(file, bodyBytes, 0644)
			if err != nil {
				//log.Fatal()
				s += "Error encountered while writing file\n"
			}
			s += fmt.Sprintf("%s is up\n", url)
			c <- s
		}
	}

}

func main() {

	channel := make(chan string)

	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	for _, url := range urls {
		go urlChecker(url, channel)

	}

	fmt.Println(runtime.NumGoroutine())
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-channel)
		fmt.Println(strings.Repeat("<>", 30))
	}

}
