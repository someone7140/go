package main

import (
	"bufio"
	"fmt"
	"os"
)

func AabMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := AabReadLine(rdr)

	aFlag := false
	bFlag := false
	result := "No"
	for _, c := range s {
		sMoji := string([]rune{c})
		if aFlag && sMoji == "b" {
			result = "Yes"
			break
		} else if bFlag && sMoji == "a" {
			result = "Yes"
			break
		} else if sMoji == "a" {
			aFlag = true
			bFlag = false
		} else if sMoji == "b" {
			bFlag = true
			aFlag = false
		} else {
			aFlag = false
			bFlag = false
		}
	}
	fmt.Println(result)

}

func AabReadLine(rdr *bufio.Reader) string {
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
