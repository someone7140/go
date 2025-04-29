package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DTakahashitheWallBreakerMain() {
	var h, w int
	fmt.Scan(&h, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	hwArrayArray := make([][]string, h)
	hwCountArrayArray := make([][]int, h)
	hwVisitedArrayArray := make([][]bool, h)
	for i := 0; i < h; i++ {
		wLine := strings.Split(DTakahashitheWallBreaker(rdr), "")
		hwArrayArray[i] = wLine

		initialRow1 := make([]int, w)
		initialRow2 := make([]bool, w)
		for j := 0; j < w; j++ {
			initialRow1[j] = -1
			initialRow2[j] = false
		}
		hwCountArrayArray[i] = initialRow1
		hwVisitedArrayArray[i] = initialRow2
	}
	abcdLine := strings.Split(DTakahashitheWallBreaker(rdr), " ")
	a, _ := strconv.Atoi(abcdLine[0])
	a = a - 1
	b, _ := strconv.Atoi(abcdLine[1])
	b = b - 1
	c, _ := strconv.Atoi(abcdLine[2])
	c = c - 1
	d, _ := strconv.Atoi(abcdLine[3])
	d = d - 1

	var loopFunc func(targetH int, targetW int, nowCount int)
	loopFunc = func(targetH int, targetW int, nowCount int) {
		hwVisitedArrayArray[targetH][targetW] = true
		hwCountArrayArray[targetH][targetW] = nowCount
		// 上に行く
		if targetH > 0 {
			nextH := targetH - 1
			nextCount := nowCount
			if !hwVisitedArrayArray[nextH][targetW] || hwCountArrayArray[nextH][targetW] > nowCount {
				if hwArrayArray[nextH][targetW] == "#" {
					nextCount = nowCount + 1
					hwCountArrayArray[nextH][targetW] = nextCount
					// もう一つ上
					if nextH > 0 {
						nextH2 := nextH - 1
						if !hwVisitedArrayArray[nextH2][targetW] || hwCountArrayArray[nextH2][targetW] > nextCount {
							hwCountArrayArray[nextH2][targetW] = nextCount
							loopFunc(nextH2, targetW, nextCount)
						}
					}

				}
				loopFunc(nextH, targetW, nextCount)
			}
		}
		// 下に行く
		if targetH < h-1 {
			nextH := targetH + 1
			nextCount := nowCount
			if !hwVisitedArrayArray[nextH][targetW] || hwCountArrayArray[nextH][targetW] > nowCount {
				if hwArrayArray[nextH][targetW] == "#" {
					nextCount = nowCount + 1
					hwCountArrayArray[nextH][targetW] = nextCount
					// もう一つ下
					if nextH < h-1 {
						nextH2 := nextH + 1
						if !hwVisitedArrayArray[nextH2][targetW] || hwCountArrayArray[nextH2][targetW] > nextCount {
							hwCountArrayArray[nextH2][targetW] = nextCount
							loopFunc(nextH2, targetW, nextCount)
						}
					}

				}
				loopFunc(nextH, targetW, nextCount)
			}
		}
		// 左に行く
		if targetW > 0 {
			nextW := targetW - 1
			nextCount := nowCount
			if !hwVisitedArrayArray[targetH][nextW] || hwCountArrayArray[targetH][nextW] > nowCount {
				if hwArrayArray[targetH][nextW] == "#" {
					nextCount = nowCount + 1
					hwCountArrayArray[targetH][nextW] = nextCount
					// もう一つ上
					if nextW > 0 {
						nextW2 := nextW - 1
						if !hwVisitedArrayArray[targetH][nextW2] || hwCountArrayArray[targetH][nextW2] > nextCount {
							hwCountArrayArray[targetH][nextW2] = nextCount
							loopFunc(targetH, nextW2, nextCount)
						}
					}

				}
				loopFunc(targetH, nextW, nextCount)
			}
		}
		// 右に行く
		if targetW < w-1 {
			nextW := targetW + 1
			nextCount := nowCount
			if !hwVisitedArrayArray[targetH][nextW] || hwCountArrayArray[targetH][nextW] > nowCount {
				if hwArrayArray[targetH][nextW] == "#" {
					nextCount = nowCount + 1
					hwCountArrayArray[targetH][nextW] = nextCount
					// もう一つ右
					if nextW < w-1 {
						nextW2 := nextW + 1
						if !hwVisitedArrayArray[targetH][nextW2] || hwCountArrayArray[targetH][nextW2] > nextCount {
							hwCountArrayArray[targetH][nextW2] = nextCount
							loopFunc(targetH, nextW2, nextCount)
						}
					}

				}
				loopFunc(targetH, nextW, nextCount)
			}
		}
	}

	loopFunc(a, b, 0)
	fmt.Println(hwCountArrayArray[c][d])
}

func DTakahashitheWallBreaker(rdr *bufio.Reader) string {
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
