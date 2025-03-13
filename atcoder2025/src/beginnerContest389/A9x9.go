package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A9x9Main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sLine := A9x9Rdr(rdr)
	sStrArray := strings.Split(sLine, "x")
	s1, _ := strconv.Atoi(sStrArray[0])
	s2, _ := strconv.Atoi(sStrArray[1])

	fmt.Println(s1 * s2)

}

func A9x9Rdr(rdr *bufio.Reader) string {
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
