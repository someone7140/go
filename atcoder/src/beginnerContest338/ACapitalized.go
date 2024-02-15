package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func ACapitalizedMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := ACapitalizedReadLine(rdr)

	result := "Yes"
	for i, c := range s {
		if i == 0 {
			if !unicode.IsUpper(c) {
				result = "No"
				break
			}
		} else {
			if unicode.IsUpper(c) {
				result = "No"
				break
			}
		}
	}

	fmt.Println(result)

}

func ACapitalizedReadLine(rdr *bufio.Reader) string {
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
