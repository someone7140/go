package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CCycleGraphMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	abMap := map[int][]int{}
	for i := 0; i < m; i++ {
		ab := strings.Split(CCycleGraphRdr(rdr), " ")
		a, _ := strconv.Atoi(ab[0])
		b, _ := strconv.Atoi(ab[1])

		aList, ok1 := abMap[a]
		if ok1 {
			aList = append(aList, b)
			abMap[a] = aList
		} else {
			abMap[a] = []int{b}
		}

		bList, ok1 := abMap[b]
		if ok1 {
			bList = append(bList, a)
			abMap[b] = bList
		} else {
			abMap[b] = []int{a}
		}
	}

	visitedSet := make(map[int]struct{})
	nowPos := 1
	visitedSet[1] = struct{}{}
	result := "Yes"
	for {
		nextList := abMap[nowPos]
		lenNext := len(nextList)
		findNext := false
		kakuteiNext := -1
		for i := 0; i < lenNext; i++ {
			next := nextList[i]
			if next != nowPos {
				_, ok := visitedSet[next]
				if findNext && !ok {
					findNext = false
					result = "No"
					break
				}
				if !ok {
					findNext = true
					if nowPos == 1 {
						kakuteiNext = next
						break
					}
					kakuteiNext = next
				}
			}
		}

		if !findNext {
			break
		}
		visitedSet[kakuteiNext] = struct{}{}
		nowPos = kakuteiNext
	}

	if result == "Yes" && len(visitedSet) != n {
		result = "No"
	}
	fmt.Println(result)
}

func CCycleGraphRdr(rdr *bufio.Reader) string {
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
