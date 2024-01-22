package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CLiningUp2Main() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CLiningUp2ReadLine(rdr), " ")

	aMap := map[int]int{}
	aReverseMap := map[int]int{}
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aMap[a] = i + 1
		aReverseMap[i+1] = a
	}

	num := aMap[-1]
	resultStrArray := make([]string, n)
	resultStrArray[0] = strconv.FormatInt(int64(num), 10)

	for i := 1; i < n; i++ {
		index := aMap[num]
		resultStrArray[i] = strconv.FormatInt(int64(index), 10)
		num = index
	}

	fmt.Println(strings.Join(resultStrArray, " "))
}

func CLiningUp2ReadLine(rdr *bufio.Reader) string {
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
