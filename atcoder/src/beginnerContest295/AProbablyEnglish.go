package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AProbablyEnglishMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	wArray := strings.Split(AProbablyEnglishReadLine(rdr), " ")
	result := "No"
	for i := 0; i < n; i++ {
		w := wArray[i]
		if w == "and" || w == "not" || w == "that" || w == "the" || w == "you" {
			result = "Yes"
			break
		}
	}

	fmt.Println(result)

}

func AProbablyEnglishReadLine(rdr *bufio.Reader) string {
	buf := make([]byte, 0, 10000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
