package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BFourHiddenMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	tStrs := strings.Split(BFourHiddenRdr(rdr), "")
	uStrs := strings.Split(BFourHiddenRdr(rdr), "")
	result := "No"
	uLen := len(uStrs)
	tLen := len(tStrs)

	for i, t := range tStrs {
		if t == uStrs[0] || t == "?" {
			if uLen == 1 {
				result = "Yes"
				break
			}
			for i2 := 1; i2 < uLen; i2++ {
				if i+i2 >= tLen {
					break
				}
				if tStrs[i+i2] != "?" && tStrs[i+i2] != uStrs[i2] {
					break
				}
				if i2 == uLen-1 {
					result = "Yes"
				}
			}
		}
		if result == "Yes" {
			break
		}
	}
	fmt.Println(result)
}

func BFourHiddenRdr(rdr *bufio.Reader) string {
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
