package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AFilterMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(AFilterReadLine(rdr), " ")
	var resultSlice []string
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		if a%2 == 0 {
			resultSlice = append(resultSlice, aStrArray[i])
		}
	}
	fmt.Println(strings.Join(resultSlice, " "))

}

func AFilterReadLine(rdr *bufio.Reader) string {
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
