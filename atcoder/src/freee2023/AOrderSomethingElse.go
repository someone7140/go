package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AOrderSomethingElseMain() {
	var n, p, q int
	fmt.Scan(&n, &p, &q)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	minD := -1
	dListStr := AOrderSomethingElseReadLine(rdr)
	for i, dMoji := range strings.Split(dListStr, " ") {
		d, _ := strconv.Atoi(dMoji)
		if i == 0 {
			minD = d
		} else {
			if minD > d {
				minD = d
			}
		}
	}

	result := p
	if p > (minD + q) {
		result = minD + q
	}

	fmt.Println(result)

}

func AOrderSomethingElseReadLine(rdr *bufio.Reader) string {
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
