package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postRequest(url string, data interface{}, contentType string) string {
	jsonStr, _ := json.Marshal(data)
	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	defer func() {
		resp.Body.Close()
		fmt.Println("http.post finish")
	}()

	result, _ := ioutil.ReadAll(resp.Body)

	return string(result)
}

func main() {
	url := "http://127.0.0.1:8080/base/captcha"
	res := postRequest(url, map[string]interface{}{}, "application/json")
	fmt.Println(res)
}
