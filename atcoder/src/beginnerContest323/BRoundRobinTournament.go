package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BRoundRobinTournamentPlayer struct {
	index    int
	winCount int
}

func BRoundRobinTournamentMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sArray := make([]BRoundRobinTournamentPlayer, n)

	for i := 0; i < n; i++ {
		s := BRoundRobinTournamentReadLine(rdr)
		winCount := 0
		for _, c := range s {
			sMoji := string([]rune{c})
			if sMoji == "o" {
				winCount = winCount + 1
			}
		}
		sArray[i] = BRoundRobinTournamentPlayer{
			index:    i + 1,
			winCount: winCount,
		}
	}

	sort.Slice(sArray, func(i, j int) bool {
		if sArray[i].winCount == sArray[j].winCount {
			return sArray[i].index < sArray[j].index
		} else {
			return sArray[i].winCount > sArray[j].winCount
		}
	})

	var resultSlice []string
	for i := 0; i < n; i++ {
		resultSlice = append(resultSlice, strconv.FormatInt(int64(sArray[i].index), 10))
	}

	fmt.Println(strings.Join(resultSlice, " "))

}

func BRoundRobinTournamentReadLine(rdr *bufio.Reader) string {
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
