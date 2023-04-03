package main

import (
	"bufio"
	"fmt"
	"os"
)

func BQualBMain() {
	var n, k int
	fmt.Scan(&n, &k)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := BQualBReadLine(rdr)
	result := ""
	count := 0
	for _, c := range s {
		sMoji := string([]rune{c})
		if count == k {
			result = result + "x"
		} else {
			if sMoji == "o" {
				count = count + 1
				result = result + "o"
			} else {
				result = result + "x"
			}
		}
	}
	fmt.Println(result)
}

func BQualBReadLine(rdr *bufio.Reader) string {
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
