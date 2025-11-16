package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type CTruckDriverRunLength struct {
	Moji  string
	Count int
}

func CTruckDriverMain() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sStrArray := strings.Split(CTruckDriverRdr(rdr), "")
	var sRunList []CTruckDriverRunLength
	before := ""
	count := 0
	for i := 0; i < n; i++ {
		s := sStrArray[i]
		if before != s && before != "" {
			sRunList = append(sRunList, CTruckDriverRunLength{Moji: before, Count: count})
			count = 1
		} else {
			count = count + 1
		}
		before = s

		if i == n-1 {
			sRunList = append(sRunList, CTruckDriverRunLength{Moji: before, Count: count})
		}
	}
	aCount := 0
	bCount := 0
	result := 0
	lIndex := 0
	rIndex := 0
	sRunLen := len(sRunList)

	deque := list.New()
	for {
		sRun := sRunList[rIndex]
		deque.PushBack(sRunList[rIndex])
		if sRun.Moji == "a" {
			aCount = aCount + sRun.Count
		}
		if sRun.Moji == "b" {
			bCount = bCount + sRun.Count
		}

		if aCount >= a && bCount < b {
			front := deque.Front().Value.(CTruckDriverRunLength)
			back := deque.Back().Value.(CTruckDriverRunLength)
			if front.Moji == "b" && back.Moji == "b" {
				result = result + 1
			} else if back.Moji == "a" && front.Moji == "a" {
				dequeLen := deque.Len()
				if dequeLen == 1 {
					for i := a; i <= aCount; i++ {
						result = result + (i - a + 1)
					}
				} else {
					backCount := back.Count
					frontCount := front.Count
					for i := a; i <= backCount+frontCount; i++ {
						result = result + (i - a + 1)
					}
					if frontCount >= a {
						result = result - (frontCount - a + 1)
					}
					if backCount >= a {
						result = result - (backCount - a + 1)
					}
				}
			} else if back.Moji == "a" {
				backAcount := back.Count
				if backAcount <= a {
					result = result + 1
				} else {
					result = result + (backAcount - a + 1)
				}
			} else if front.Moji == "a" {
				frontAcount := front.Count
				if frontAcount <= a {
					result = result + 1
				} else {
					result = result + (frontAcount - a + 1)
				}
			}
			rIndex = rIndex + 1
		} else if bCount >= b {
			for {
				front := deque.Front().Value.(CTruckDriverRunLength)
				if front.Moji == "a" {
					aCount = aCount - front.Count
				}
				if front.Moji == "b" {
					bCount = bCount - front.Count
				}
				deque.Remove(deque.Front())
				lIndex = lIndex + 1
				if bCount < b {
					rIndex = rIndex + 1
					break
				}
			}
		} else {
			rIndex = rIndex + 1
		}

		if lIndex >= sRunLen || rIndex >= sRunLen {
			break
		}
	}

	fmt.Println(result)
}

func CTruckDriverRdr(rdr *bufio.Reader) string {
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
