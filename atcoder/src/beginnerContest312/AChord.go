package main

import (
	"bufio"
	"fmt"
	"os"
)

func AChordMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := AChordReadLine(rdr)
	result := "No"

	if s == "ACE" || s == "BDF" || s == "CEG" || s == "DFA" || s == "EGB" || s == "FAC" || s == "GBD" {
		result = "Yes"
	}
	fmt.Println(result)

}

func AChordReadLine(rdr *bufio.Reader) string {
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
