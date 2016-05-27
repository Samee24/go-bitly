package main

import (
"fmt"
"net/http"
"io/ioutil"
//"encoding/json"
)

func main() {
    fmt.Println("Hello World")
    resp, err := http.Get("https://api-ssl.bitly.com/v3/shorten?access_token=ACCESSTOKEN&longUrl=http%3A%2F%2Fgoogle.com%2F")
    if err != nil {
      fmt.Println("Error occurred")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
