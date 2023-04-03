package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BPostalCardMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sArray := make([]string, n)
	for i := 0; i < n; i++ {
		sArray[i] = BPostalCardReadLine(rdr)
	}

	result := 0
	tArray := make([]string, m)
	for i := 0; i < m; i++ {
		tArray[i] = BPostalCardReadLine(rdr)
	}
	for i := 0; i < n; i++ {

		for j := 0; j < m; j++ {
			if strings.HasSuffix(sArray[i], tArray[j]) {
				result = result + 1
				break
			}
		}
	}

	fmt.Println(result)
}

func BPostalCardReadLine(rdr *bufio.Reader) string {
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
