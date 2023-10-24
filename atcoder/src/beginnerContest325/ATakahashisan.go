package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ATakahashisanMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	stStrs := strings.Split(ATakahashisanReadLine(rdr), " ")

	fmt.Println(stStrs[0] + " san")

}

func ATakahashisanReadLine(rdr *bufio.Reader) string {
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
