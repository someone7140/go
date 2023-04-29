package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ANchoicequestionMain() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	sum := a + b

	result := 0
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	cStrArray := strings.Split(ANchoicequestionReadLine(rdr), " ")
	for i := 0; i < n; i++ {
		c, _ := strconv.Atoi(cStrArray[i])
		if c == sum {
			result = i + 1
			break
		}
	}

	fmt.Println(result)

}

func ANchoicequestionReadLine(rdr *bufio.Reader) string {
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
