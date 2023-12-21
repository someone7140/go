package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
)

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		// certainがfalseの場合もそのまま使う
		_, name, _ := charset.DetermineEncoding(data, "")
		if name != "" {
			return name
		}
	}
	return "utf-8"
}

func htmlEncodeSampleMain() {
	// htmlページの取得
	// webPage := ("https://mayukasports.blogspot.com/")     // UTF-8のサイト
	webPage := ("http://abehiroshi.la.coocan.jp/top.htm") // Shift_JISのサイト
	resp, err := http.Get(webPage)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 文字コードチェック用のreadのためにbodyをコピーしておく
	body := new(bytes.Buffer)
	bodyForCheck := io.TeeReader(resp.Body, body)
	// 文字コードの判定
	charset := detectContentCharset(bodyForCheck)
	encode, err := htmlindex.Get(charset)
	if err != nil {
		fmt.Println(err)
		return
	}

	// htmlのRead
	var contentBytes []byte
	if name, _ := htmlindex.Name(encode); name != "utf-8" {
		encodeBody := encode.NewDecoder().Reader(body)
		contentBytes, err = io.ReadAll(encodeBody)
	} else {
		contentBytes, err = io.ReadAll(body)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(contentBytes))
}
