package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BRankingwithTiesPerson struct {
	index int
	point int
	rank  int
}

func BRankingwithTiesMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	pLine := BRankingwithTiesRdr(rdr)
	pStrArray := strings.Split(pLine, " ")
	pArray := make([]BRankingwithTiesPerson, n)

	for i := 0; i < n; i++ {
		p, _ := strconv.Atoi(pStrArray[i])
		pArray[i] = BRankingwithTiesPerson{
			index: i,
			point: p,
			rank:  0,
		}
	}

	sortedByPointArray := make([]BRankingwithTiesPerson, n)
	copy(sortedByPointArray, pArray)
	sort.Slice(sortedByPointArray, func(i, j int) bool { return sortedByPointArray[j].point < sortedByPointArray[i].point })
	tempRank := 1
	beforePoint := -1
	plusCount := 0
	sortedByPointWithRankArray := make([]BRankingwithTiesPerson, n)
	for i := 0; i < n; i++ {
		p := sortedByPointArray[i]
		if beforePoint != p.point {
			tempRank = tempRank + plusCount
			sortedByPointWithRankArray[i] = BRankingwithTiesPerson{
				index: p.index,
				point: p.point,
				rank:  tempRank,
			}
			plusCount = 1
		} else {
			sortedByPointWithRankArray[i] = BRankingwithTiesPerson{
				index: p.index,
				point: p.point,
				rank:  sortedByPointWithRankArray[i-1].rank,
			}
			plusCount = plusCount + 1
		}
		beforePoint = p.point
	}

	var resultSlice []string
	sort.Slice(sortedByPointWithRankArray, func(i, j int) bool {
		return sortedByPointWithRankArray[i].index < sortedByPointWithRankArray[j].index
	})

	for i := 0; i < n; i++ {
		resultSlice = append(resultSlice, strconv.FormatInt(int64(sortedByPointWithRankArray[i].rank), 10))
	}
	fmt.Println(strings.Join(resultSlice, "\n"))

}

func BRankingwithTiesRdr(rdr *bufio.Reader) string {
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
