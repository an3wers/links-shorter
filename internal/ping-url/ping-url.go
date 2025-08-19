package pingurl

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func PingUrl() {
	path := flag.String("file", "./url.txt", "file with urls")

	flag.Parse()

	file, err := os.ReadFile(*path)

	if err != nil {
		panic(err.Error())
	}

	urlSlice := strings.Split(string(file), "\n")

	resCh := make(chan int)
	errCh := make(chan error)

	for _, url := range urlSlice {
		go ping(url, resCh, errCh)
	}

	for i := 0; i < len(urlSlice); i++ {
		select {
		case res := <-resCh:
			fmt.Println(res)
		case errRes := <-errCh:
			fmt.Println(errRes)
		}

	}
}

func ping(urlVal string, ch chan int, errCh chan error) {

	cleanUrl := strings.TrimSpace(urlVal)
	parsedUrl, err := url.Parse(cleanUrl)

	if err != nil {
		errCh <- err
		return
	}

	resp, err := http.Get(parsedUrl.String())

	if err != nil {
		errCh <- err
		return
	}

	if resp.StatusCode != 200 {
		errCh <- fmt.Errorf("status code: %d", resp.StatusCode)
		return
	}
	ch <- resp.StatusCode

}
