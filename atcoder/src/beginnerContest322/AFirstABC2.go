package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AFirstABC2Main() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := AFirstABC2ReadLine(rdr)

	result := strings.Index(s, "ABC")

	if result >= 0 {
		fmt.Println(result + 1)
	} else {
		fmt.Println(result)
	}

}

func AFirstABC2ReadLine(rdr *bufio.Reader) string {
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
