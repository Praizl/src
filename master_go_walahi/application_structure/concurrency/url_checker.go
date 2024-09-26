package main

/*import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

/*func urlChecker(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s>>>>>Status Code:%d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			_ = err
			file := strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("writing response body to %s\n", file)
			fmt.Println(strings.Repeat("<>", 20))
			err = os.WriteFile(file, bodyBytes, 0644)
			if err != nil {
				log.Fatal()
			}
		}
	}
	wg.Done()
}

func main() {

	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go urlChecker(url, &wg)

	}
	wg.Wait()
}*/
