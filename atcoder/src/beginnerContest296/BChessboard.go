package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func BChessboardMain() {

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	result := ""
	for i := 0; i < 8; i++ {
		sStr := BChessboardReadLine(rdr)
		if result == "" {
			for j, c := range sStr {
				sMoji := string([]rune{c})
				if sMoji == "*" {
					yoko := string([]rune{rune('a' + j)})
					tate := strconv.FormatInt(int64(8-i), 10)
					result = yoko + tate
				}
			}
		}
	}
	fmt.Println(result)
}

func BChessboardReadLine(rdr *bufio.Reader) string {
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
