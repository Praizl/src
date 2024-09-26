package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func urlChecker(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		//fmt.Println(err)
		s := fmt.Sprintf("%s is down!\n", url)
		s += fmt.Sprintf("Error:%v\n", err)
		fmt.Println(s)
		c <- url

	} else {

		s := fmt.Sprintf("%s>>>>>Status Code:%d\n", url, resp.StatusCode)

		s += fmt.Sprintf("%s is up\n", url)
		fmt.Println(s)
		c <- url

	}
}

func main() {

	channel := make(chan string)

	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	for _, url := range urls {
		go urlChecker(url, channel)

	}

	fmt.Println(runtime.NumGoroutine())
	//for i := 0; i < len(urls); i++ {
	//fmt.Println(<-channel)
	//}

	/*for {
		go urlChecker(<-channel, channel)
		fmt.Println(strings.Repeat("<>", 30))
		time.Sleep(time.Minute / 10)
	}*/

	//iterating over the channel
	for url := range channel {
		go func(u string) {
			time.Sleep(time.Minute / 10)
			urlChecker(u, channel)
		}(url)
	}

}
