package main

import (
	"bufio"
	"fmt"
	"os"
)

func BExtendedABCMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := BExtendedABCReadLine(rdr)

	result := "Yes"

	if len(s) > 0 {
		now := ""
		next := "A"
		for i, c := range s {
			sMoji := string([]rune{c})
			if i == 0 {
				if sMoji == "A" {
					now = "A"
					next = "B"
				} else if sMoji == "B" {
					now = "B"
					next = "C"
				} else if sMoji == "C" {
					now = "C"
					next = "C"
				} else {
					result = "No"
					break
				}
			} else {
				if now == "" && next == "A" && sMoji == "A" {
					now = "A"
					next = "B"
				} else if now == "A" && next == "B" && sMoji == "B" {
					now = "B"
					next = "C"
				} else if now == "A" && next == "B" && sMoji == "C" {
					now = "C"
					next = "C"
				} else if now == "B" && next == "C" && sMoji == "C" {
					now = "C"
					next = "C"
				} else if now == sMoji {
					// 何もしない
				} else {
					result = "No"
					break
				}
			}

		}
	}

	fmt.Println(result)
}

func BExtendedABCReadLine(rdr *bufio.Reader) string {
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
