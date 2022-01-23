package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errReqeustFailed = errors.New("Request Failed")

func main() {
	// var results = make(map[string]string)
	// urls := []string{
	// 	"https://www.airbnb.com/",
	// 	"https://www.google.com/",
	// 	"https://www.amazon.com/",
	// 	"https://www.reddit.com/",
	// 	"https://www.google.com/",
	// 	"https://soundcloud.com/",
	// 	"https://www.facebook.com/",
	// 	"https://www.instagram.com/",
	// 	"https://academy.nomadcoders.co/",
	// }
	// for _, url := range urls {
	// 	result := "OK"
	// 	err := hitURL(url)
	// 	if err != nil {
	// 		result = "FAIL"
	// 	}
	// 	results[url] = result
	// }
	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errReqeustFailed
	}
	return nil
}

func sexyCount(person string) {
	for i:= 0;i<10;i++{
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}