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
// Instance to hold base-level response from the API
type ShortenResponse struct {
    Status  string `json:"status_txt"`
    // Nested field
    ShortenResponseData `json:"data"`
}

// Instance to hold data we actually care about from the API
type ShortenResponseData struct {
    ShortUrl  string `json:"url"`
}

func main() {
    accessToken := ""
    // Make sure we encode our URI
    longUrl := url.QueryEscape("http://samee.ninja")

    resp, err := http.Get(baseApiURI + shortenEndPoint + "?access_token=" + accessToken + "&longUrl=" + longUrl)
    if err != nil {
      fmt.Println("Error occurred")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    // Print JSON response
    fmt.Println(string(body))

    var respWrapper ShortenResponse
    err = json.Unmarshal([]byte(string(body)), &respWrapper)
    if err != nil {
      fmt.Println("Error occurred")
    }

    // Print shortened link
    fmt.Println(pojo.ShortenResponseData.ShortUrl)
}
