package main

import (
	"bufio"
	"fmt"
	"os"
)

func ASimilarStringMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sString := ASimilarStringReadLine(rdr)
	tString := ASimilarStringReadLine(rdr)
	result := "Yes"

	sChars := []rune(sString)
	tChars := []rune(tString)

	for i := 0; i < n; i++ {
		s := string([]rune{sChars[i]})
		t := string([]rune{tChars[i]})

		if s != t {
			if s == "1" || s == "l" {
				if t != "1" && t != "l" {
					result = "No"
					break
				}
			} else if s == "0" || s == "o" {
				if t != "0" && t != "o" {
					result = "No"
					break
				}
			} else {
				result = "No"
				break
			}
		}
	}
	fmt.Println(result)

}

func ASimilarStringReadLine(rdr *bufio.Reader) string {
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
