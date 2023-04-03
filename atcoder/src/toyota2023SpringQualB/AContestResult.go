package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AContestResultMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aArr := strings.Split(AContestResultReadLine(rdr), " ")
	bArr := strings.Split(AContestResultReadLine(rdr), " ")
	result := 0
	for i := 0; i < m; i++ {
		b, _ := strconv.Atoi(bArr[i])
		a, _ := strconv.Atoi(aArr[b-1])
		result = result + a
	}
	fmt.Println(result)
}

func AContestResultReadLine(rdr *bufio.Reader) string {
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
