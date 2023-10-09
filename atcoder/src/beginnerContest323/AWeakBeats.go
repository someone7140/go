package main

import (
	"bufio"
	"fmt"
	"os"
)

func AWeakBeatsMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := AWeakBeatsReadLine(rdr)
	result := "Yes"
	for i, c := range s {
		sMoji := string([]rune{c})
		if i%2 == 1 && sMoji == "1" {
			result = "No"
			break
		}
	}

	fmt.Println(result)

}

func AWeakBeatsReadLine(rdr *bufio.Reader) string {
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
