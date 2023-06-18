package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type CCentersResult struct {
	number      string
	centerIndex int
}

func CCentersMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	astrList := strings.Split(CCentersReadLine(rdr), " ")

	sMap := map[string]int{}
	var tempResultSlice []CCentersResult

	for i, aStr := range astrList {
		v, ok := sMap[aStr]
		if !ok {
			sMap[aStr] = 1
		} else if v == 1 {
			tempResultSlice = append(tempResultSlice, CCentersResult{
				number:      aStr,
				centerIndex: i,
			})
			sMap[aStr] = 2
		}
	}

	sort.Slice(tempResultSlice, func(i, j int) bool { return tempResultSlice[j].centerIndex > tempResultSlice[i].centerIndex })

	var resultSlice []string
	for _, temp := range tempResultSlice {
		resultSlice = append(resultSlice, temp.number)
	}
	fmt.Println(strings.Join(resultSlice, " "))
}

func CCentersReadLine(rdr *bufio.Reader) string {
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
