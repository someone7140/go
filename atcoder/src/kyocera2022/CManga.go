package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CMangaMain() {
	var n int
	fmt.Scan(&n)

	result := 0
	queueLen := 0
	aArray := make([]int, n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CMangaRdr(rdr), " ")

	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray[i] = a
	}

	sort.Ints(aArray)

	for i := 0; i < n; i++ {
		a := aArray[i]
		if a == (result + 1) {
			result = result + 1
		} else {
			queueLen = queueLen + 1
		}
	}
	result = result + queueLen/2
	fmt.Println(result)
}

func CMangaRdr(rdr *bufio.Reader) string {
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
