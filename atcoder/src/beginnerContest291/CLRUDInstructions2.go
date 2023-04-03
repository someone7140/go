package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CLRUDInstructions2Main() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := CLRUDInstructions2ReadLine(rdr)
	sMap := map[string]bool{}
	result := "No"

	x := 0
	y := 0
	sMap["0-0"] = true
	for _, c := range s {
		sMoji := string([]rune{c})
		if sMoji == "R" {
			x = x + 1
		} else if sMoji == "L" {
			x = x - 1
		} else if sMoji == "U" {
			y = y + 1
		} else {
			y = y - 1
		}
		nowPosStr := strconv.FormatInt(int64(x), 10) + "-" + strconv.FormatInt(int64(y), 10)
		visited, ok := sMap[nowPosStr]
		if visited && ok {
			result = "Yes"
			break
		} else {
			sMap[nowPosStr] = true
		}
	}

	fmt.Println(result)

}

func CLRUDInstructions2ReadLine(rdr *bufio.Reader) string {
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
