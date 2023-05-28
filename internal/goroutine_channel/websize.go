package goroutine_channel

import (
	"io"
	"log"
	"net/http"
)

/*
测量页面大小
*/

type Page struct {
	URL  string
	Size int
}

func ResponseSize(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//Get all Body
	//fmt.Println(string(body))
	//Get len
	channel <- Page{URL: url, Size: len(body)}
}
