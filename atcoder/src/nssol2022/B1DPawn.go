package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func B1DPawnMain() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aStrArray := strings.Split(B1DPawnRdr(rdr), " ")
	aArray := make([]int, k)
	for i := 0; i < k; i++ {
		aArray[i], _ = strconv.Atoi(aStrArray[i])
	}

	lStrArray := strings.Split(B1DPawnRdr(rdr), " ")
	for i := 0; i < q; i++ {
		l, _ := strconv.Atoi(lStrArray[i])
		a := aArray[l-1]
		if a != n {
			// 右端
			if l == k {
				aArray[l-1] = a + 1
			} else {
				if (aArray[l] - aArray[l-1]) != 1 {
					aArray[l-1] = a + 1
				}
			}
		}
	}

	results := make([]string, k)
	for i := 0; i < k; i++ {
		results[i] = strconv.FormatInt(int64(aArray[i]), 10)
	}
	fmt.Println(strings.Join(results, " "))
}

func B1DPawnRdr(rdr *bufio.Reader) string {
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
