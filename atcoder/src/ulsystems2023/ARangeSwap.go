package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ARangeSwapMain() {

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	npqrs := strings.Split(ARangeSwapReadLine(rdr), " ")
	n, _ := strconv.Atoi(npqrs[0])
	p, _ := strconv.Atoi(npqrs[1])
	q, _ := strconv.Atoi(npqrs[2])
	r, _ := strconv.Atoi(npqrs[3])
	s, _ := strconv.Atoi(npqrs[4])

	aArray := strings.Split(ARangeSwapReadLine(rdr), " ")
	var results []string
	sub1 := aArray[p-1 : q]
	sub2 := aArray[r-1 : s]

	for i := 0; i < n; i++ {
		if i == p-1 {
			results = append(results, sub2...)
		}
		if i == r-1 {
			results = append(results, sub1...)
		}

		if (i < p-1 || i > q-1) && (i < r-1 || i > s-1) {
			results = append(results, aArray[i])
		}
	}
	fmt.Println(strings.Join(results, " "))
}

func ARangeSwapReadLine(rdr *bufio.Reader) string {
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
