package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func ACBCMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := ACBCRdr(rdr)
	result := ""

	for _, c := range s {
		if unicode.IsUpper(c) {
			result = result + string([]rune{c})
		}
	}

	fmt.Println(result)
}

func ACBCRdr(rdr *bufio.Reader) string {
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
