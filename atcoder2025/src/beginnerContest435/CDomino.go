package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CDominoMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CDominoRdr(rdr), " ")
	aArray := make([]int, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray[i] = a
	}

	result := 1
	for i := 0; i < n; i++ {
		nextIndex := i + aArray[i] - 1
		temp := nextIndex + 1
		if result > temp {
			continue
		} else if temp >= n {
			result = n
			break
		} else if result == temp && i+1 == result {
			break
		} else {
			result = temp
		}
	}

	fmt.Println(result)
}

func CDominoRdr(rdr *bufio.Reader) string {
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
