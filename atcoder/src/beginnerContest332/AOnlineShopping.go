package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AOnlineShoppingMain() {
	var n, s, k int
	fmt.Scan(&n, &s, &k)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sum := 0

	for i := 0; i < n; i++ {
		pqfListStr := strings.Split(AOnlineShoppingReadLine(rdr), " ")
		price, _ := strconv.Atoi(pqfListStr[0])
		kosuu, _ := strconv.Atoi(pqfListStr[1])
		sum = sum + price*kosuu
	}

	if sum < s {
		sum = sum + k
	}
	fmt.Println(sum)

}

func AOnlineShoppingReadLine(rdr *bufio.Reader) string {
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
