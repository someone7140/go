package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CSecurity2Main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sStrList := strings.Split(CSecurity2Rdr(rdr), "")
	bCount := int64(0)
	for i, j := 0, len(sStrList)-1; i < j; i, j = i+1, j-1 {
		sStrList[i], sStrList[j] = sStrList[j], sStrList[i]
	}

	for _, sStr := range sStrList {
		s, _ := strconv.Atoi(sStr)
		amariCountB := int(bCount % 10)
		sTarget := (s - amariCountB)
		if sTarget < 0 {
			sTarget = 10 + sTarget
		}

		bCount = bCount + int64(sTarget)
	}

	fmt.Println(bCount + int64(len(sStrList)))
}

func CSecurity2Rdr(rdr *bufio.Reader) string {
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
