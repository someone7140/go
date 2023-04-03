package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AManyABProblemsMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var results []string
	for i := 0; i < n; i++ {
		abArr := strings.Split(AManyABProblemsReadLine(rdr), " ")
		a, _ := strconv.Atoi(abArr[0])
		b, _ := strconv.Atoi(abArr[1])
		result := a + b
		results = append(results, strconv.FormatInt(int64(result), 10))
	}
	fmt.Println(strings.Join(results, "\n"))
}

func AManyABProblemsReadLine(rdr *bufio.Reader) string {
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
