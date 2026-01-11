package main

import (
	"bufio"
	"fmt"
	"os"
)

func AopaddingMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := AopaddingRdr(rdr)
	sabun := n - len(s)
	for i := 0; i < sabun; i++ {
		s = "o" + s
	}

	fmt.Println(s)
}

func AopaddingRdr(rdr *bufio.Reader) string {
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
