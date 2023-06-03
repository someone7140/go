package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CVirusXy struct {
	x      int
	y      int
	kansen bool
}

func CVirusXyMain() {
	var n, d int
	fmt.Scan(&n, &d)

	d2 := int64(d * d)
	xyMap := map[int]CVirusXy{}
	notKansenMap := map[int]bool{}

	startX := 0
	startY := 0
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < n; i++ {
		xyString := strings.Split(CVirusReadLine(rdr), " ")
		x, _ := strconv.Atoi(xyString[0])
		y, _ := strconv.Atoi(xyString[1])
		if i == 0 {
			startX = x
			startY = y
			xyMap[i] = CVirusXy{
				x:      x,
				y:      y,
				kansen: true,
			}
		} else {
			kyori2 := int64((startX-x)*(startX-x)) + int64((startY-y)*(startY-y))
			if kyori2 <= d2 {
				xyMap[i] = CVirusXy{
					x:      x,
					y:      y,
					kansen: true,
				}
			} else {
				xyMap[i] = CVirusXy{
					x:      x,
					y:      y,
					kansen: false,
				}
				notKansenMap[i] = true
			}
		}
	}
	var loopFunc func(targetIndex int, targetXy CVirusXy)
	loopFunc = func(targetIndex int, targetXy CVirusXy) {
		for key := range notKansenMap {
			if targetIndex != key {
				v2, _ := xyMap[key]
				if !v2.kansen {
					kyori2 := int64((v2.x-targetXy.x)*(v2.x-targetXy.x)) + int64((v2.y-targetXy.y)*(v2.y-targetXy.y))
					if kyori2 <= d2 {
						xyMap[key] = CVirusXy{
							x:      v2.x,
							y:      v2.y,
							kansen: true,
						}
						delete(notKansenMap, key)
						loopFunc(key, v2)
					}
				}
			}
		}
	}
	// 最初以外で感染したか
	for i := 1; i < n; i++ {
		v, _ := xyMap[i]
		if v.kansen {
			loopFunc(i, v)
		}
	}

	var resultSlice []string
	for i := 0; i < n; i++ {
		v, _ := xyMap[i]
		if v.kansen {
			resultSlice = append(resultSlice, "Yes")
		} else {
			resultSlice = append(resultSlice, "No")
		}
	}
	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CVirusReadLine(rdr *bufio.Reader) string {
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
