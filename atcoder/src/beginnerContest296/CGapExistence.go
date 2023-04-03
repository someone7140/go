package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CGapExistenceMain() {
	var n, x int64
	fmt.Scan(&n, &x)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CGapExistenceReadLine(rdr), " ")
	aMap := map[int64]bool{}
	var i int64

	for i = 0; i < n; i++ {
		aInt, _ := strconv.Atoi(aStrArray[i])
		aMap[int64(aInt)] = true
	}

	result := "No"
	if x == 0 {
		result = "Yes"
	} else {
		for key := range aMap {
			target := key + x
			v, ok := aMap[target]
			if v && ok {
				result = "Yes"
				break
			}
		}
	}

	fmt.Println(result)
}

func CGapExistenceReadLine(rdr *bufio.Reader) string {
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
