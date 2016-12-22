# go-bitly
go-bitly is a wrapper written in go for the bitly REST API. 

### authentication
Right now, you can insert your key in the accessToken field. I plan on having a `config.json` to store these keys, and have them parsed in the future.
```
accessToken := "[YOUR_KEY_GOES_HERE]"
```
#### shorten

To shorten URLs, simply run the following command:
```
go run bitylywrapper.go -s [YOUR URL]
```

#### expand

To expand bit.ly URLs, simply run the following command:
```
go run bitylywrapper.go -e [YOUR URL]
```
