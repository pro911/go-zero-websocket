package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fech(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (windows NT 10.0;Win64; x64) Applewebkit/537.36 (KHTML,like Gecko) Chrome/99.0.4833")
	req.Header.Add("Cookie", "__gads=ID=a-03")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP get err:", err)
		return ""
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("http status code:", resp.StatusCode)
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return ""
	}
	return string(body)
}

func main() {
	url := "https://www.baidu.com/"
	bodystring := fech(url)
	fmt.Println(bodystring)
}
