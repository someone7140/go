package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BMinimizeAbs1Main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(BMinimizeAbs1ReadLine(rdr), " ")

	results := make([]int, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrs[i])
		temp := 0
		if a <= l {
			temp = l
		} else if a >= r {
			temp = r
		} else {
			temp = a
		}

		results[i] = temp
	}

	resultStrs := make([]string, n)
	for i := 0; i < n; i++ {
		resultStrs[i] = strconv.FormatInt(int64(results[i]), 10)
	}
	fmt.Println(strings.Join(resultStrs, " "))

}

func BMinimizeAbs1ReadLine(rdr *bufio.Reader) string {
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
