package main

import (
	"bufio"
	"fmt"
	"os"
)

func ASequenceofStringsMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aArray := make([]string, n)

	for i := 0; i < n; i++ {
		aArray[i] = ASequenceofStringsReadLine(rdr)
	}

	resultArray := make([]string, n)
	for i := 0; i < n; i++ {
		resultArray[i] = aArray[n-1-i]
	}

	for i := 0; i < n; i++ {
		fmt.Println(resultArray[i])
	}

}

func ASequenceofStringsReadLine(rdr *bufio.Reader) string {
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
