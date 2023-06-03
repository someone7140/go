package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AFirstPlayer struct {
	s string
	a int
}

func AFirstPlayerMain() {
	var n int
	fmt.Scan(&n)
	saSlice := make([]AFirstPlayer, n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	minIndex := -1
	minAge := -1
	for i := 0; i < n; i++ {
		sAString := strings.Split(AFirstPlayerReadLine(rdr), " ")
		a, _ := strconv.Atoi(sAString[1])
		sa := AFirstPlayer{
			s: sAString[0],
			a: a,
		}
		saSlice[i] = sa
		if minAge == -1 || minAge > a {
			minIndex = i
			minAge = a
		}
	}

	var resultSlice []string
	index := minIndex
	for i := 0; i < n; i++ {
		if index == n {
			index = 0
		}
		resultSlice = append(resultSlice, saSlice[index].s)
		index = index + 1
	}
	fmt.Println(strings.Join(resultSlice, "\n"))

}

func AFirstPlayerReadLine(rdr *bufio.Reader) string {
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
