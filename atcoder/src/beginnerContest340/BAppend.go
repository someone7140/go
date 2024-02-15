package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BAppendMain() {
	var q int
	fmt.Scan(&q)

	var aStrArray []string
	var resultStrArray []string
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	len := 0
	for i := 0; i < q; i++ {
		qStrArray := strings.Split(BAppendReadLine(rdr), " ")
		if qStrArray[0] == "1" {
			aStrArray = append(aStrArray, qStrArray[1])
			len = len + 1
		} else {
			targetIndex, _ := strconv.Atoi(qStrArray[1])
			resultStrArray = append(resultStrArray, aStrArray[len-targetIndex])
		}
	}

	fmt.Println(strings.Join(resultStrArray, "\n"))

}

func BAppendReadLine(rdr *bufio.Reader) string {
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
