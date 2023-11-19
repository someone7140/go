package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AWeakBeatsMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	resultStrs := strings.Split(ASpreadReadLine(rdr), "")

	fmt.Println(strings.Join(resultStrs, " "))

}

func ASpreadReadLine(rdr *bufio.Reader) string {
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
