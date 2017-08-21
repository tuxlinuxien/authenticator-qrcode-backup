package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

// Entry .
type Entry struct {
	Account string `json:"account"`
	Hash    string `json:"hash"`
	Issuer  string `json:"issuer"`
	Secret  string `json:"secret"`
	Type    string `json:"type"`
}

// otpauth://$TYPE/$ACCOUNT?secret=$SECRETissuer=$ISSUER
func (e Entry) toURI() string {
	account, _ := url.QueryUnescape(e.Account)
	uri := url.URL{
		Scheme: "otpauth",
		Host:   e.Type,
		Path:   account,
	}
	value := url.Values{}
	value.Add("issuer", e.Issuer)
	value.Add("secret", e.Secret)
	return uri.String() + "?" + value.Encode()
}

func showBitmap(q *qrcode.QRCode) {
	bitmap := q.Bitmap()
	for y := range bitmap {
		for x := range bitmap[y] {
			if bitmap[y][x] {
				fmt.Print("\033[40;37m  \033[0m")
			} else {
				fmt.Print("\033[47;30m  \033[0m")
			}
		}
		fmt.Println()
	}
}

// if qr.shrinkPoints[x][y] == 1 {
// 				fmt.Print("\033[40;37m  \033[0m")
// 			} else {
// 				fmt.Print("\033[47;30m  \033[0m")
// 			}

func main() {
	dec := json.NewDecoder(os.Stdin)
	c := map[string]Entry{}
	dec.Decode(&c)

	for _, e := range c {
		totpURI := e.toURI()
		q, _ := qrcode.New(totpURI, qrcode.Low)
		showBitmap(q)
		fmt.Println(q.Content)
	}
}
