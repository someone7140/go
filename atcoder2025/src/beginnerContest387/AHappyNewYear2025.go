package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AHappyNewYear2025Main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	abLine := AHappyNewYear2025Rdr(rdr)
	sStrArray := strings.Split(abLine, " ")
	a, _ := strconv.Atoi(sStrArray[0])
	b, _ := strconv.Atoi(sStrArray[1])

	fmt.Println((a + b) * (a + b))

}

func AHappyNewYear2025Rdr(rdr *bufio.Reader) string {
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
