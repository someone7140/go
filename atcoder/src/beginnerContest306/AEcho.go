package main

import (
	"bufio"
	"fmt"
	"os"
)

func AEchoMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := AEchoReadLine(rdr)
	result := ""
	for _, c := range s {
		sMoji := string([]rune{c})
		result = result + sMoji + sMoji
	}
	fmt.Println(result)

}

func AEchoReadLine(rdr *bufio.Reader) string {
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
