package main

import (
"fmt"
"net/http"
"io/ioutil"
"net/url"
"encoding/json"
)

const (
  baseApiURI   = "https://api-ssl.bitly.com"
  shortenEndPoint = "/v3/shorten"
)

type ShortenResponse struct {
    Status  string `json:"status_txt"`
    ShortenResponseData `json:"data"`
}

type ShortenResponseData struct {
    ShortUrl  string `json:"url"`
}

func main() {
    accessToken := ""
    longUrl := url.QueryEscape("http://samee.ninja")

    var pojo ShortenResponse
    resp, err := http.Get(baseApiURI + shortenEndPoint + "?access_token=" + accessToken + "&longUrl=" + longUrl)
    if err != nil {
      fmt.Println("Error occurred")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))

    json.Unmarshal([]byte(string(body)), &pojo)
    fmt.Println(pojo.ShortenResponseData.ShortUrl)
}
