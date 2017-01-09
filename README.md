# go-bitly
go-bitly is a command-line utility written in go for the bitly REST API. Will eventually turn into a wrapper, but I was toying around with go, and command-line flags, so this is the state it's in as of the moment.

### authentication
Right now, you can insert your key in the accessToken field. I plan on having a `config.json` to store these keys, and have them parsed in the future.
```
accessToken := "[YOUR_KEY_GOES_HERE]"
```
#### shorten

To shorten URLs, simply run the following command:
```
go run bitylyWrapper.go -s [YOUR URL]
```

#### expand

To expand bit.ly URLs, simply run the following command:
```
go run bitylyWrapper.go -e [YOUR URL]
```
