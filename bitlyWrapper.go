package main

import (
"fmt"
"net/http"
"io/ioutil"
"net/url"
"encoding/json"
"os"
"flag"
)

const (
  accessToken="ee3b33ba448f179ac37bdf4bfb888fa032f78d18"
  baseApiURI="https://api-ssl.bitly.com"
  shortenEndPoint="/v3/shorten"
  expandEndPoint="/v3/expand"
  urlShortenError="Please provide a url to shorten"
  accessTokenError="Please provide an access token"
  noFlagError="Please provide a flag to either shorten (-s) or expand (-e) your link."
  tooManyFlagsError="Please provide only ONE flag to either shorten (-s) or expand (-e) your link."
)
// Instance to hold base-level response from the API
type Response struct {
    Status string `json:"status_txt"`
    // Nested field
    ResponseData `json:"data"`
}

// Instance to hold data we actually care about from the API
type ResponseData struct {
    ShortUrl string `json:"url"`
    ExpandedData []ExpandResponseData `json:"expand"`
}

// Instance to hold data we actually care about from the expand API call
type ExpandResponseData struct {
    LongUrl string `json:"long_url"`
}

func checkAccessToken(token string) {
  if len(token) < 1 {
      fmt.Println(accessTokenError)
      os.Exit(1)
    }
}

func checkArgs(args []string) {
  if len(args) < 1 {
      fmt.Println(urlShortenError)
      os.Exit(1)
    }
}

func checkFlagsAndAssignEndpoint(s bool, e bool, url string) string {
  if !s && !e {
      fmt.Println(noFlagError)
      os.Exit(1)
    }
  if s && e {
      fmt.Println(tooManyFlagsError)
      os.Exit(1)
  }
  if s {
    return baseApiURI + shortenEndPoint + "?access_token=" + accessToken + "&longUrl=" + url
  } else {
    return baseApiURI + expandEndPoint + "?access_token=" + accessToken + "&shortUrl=" + url
  }

}

func main() {
    shortenFlag := flag.Bool("s", false, "a bool for shortening")
    expandFlag := flag.Bool("e", false, "a bool for expanding")

    checkArgs(os.Args)
    checkAccessToken(accessToken)

    flag.Parse()

    // Make sure we encode our URI
    url := url.QueryEscape(os.Args[2])
    currentEndpoint := checkFlagsAndAssignEndpoint(*shortenFlag, *expandFlag, url)

    resp, err := http.Get(currentEndpoint)

    if err != nil {
      fmt.Printf("Error occurred\n %s", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    // Parse JSON
    var respWrapper = new(Response)
    err = json.Unmarshal([]byte(body), &respWrapper)
    if err != nil {
      fmt.Printf("Error occurred: %s", respWrapper.Status)
    }

    // Print shortened/expanded link
    if (*shortenFlag) {
       fmt.Println(respWrapper.ResponseData.ShortUrl)
    } else {
       fmt.Println(respWrapper.ResponseData.ExpandedData[0].LongUrl)
    }
}
