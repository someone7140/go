package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	hwArrayArray := make([][]string, h)
	hwAlreadyArrayArray := make([][]int, h)
	for i := 0; i < h; i++ {
		wArray := make([]string, w)
		hRecord := DSnukeMazeReadLine(rdr)
		for j, c := range hRecord {
			wArray[j] = string([]rune{c})
		}
		hwArrayArray[i] = wArray
		hwAlreadyArrayArray[i] = make([]int, w)
	}
	result := "No"

	var judgeFunc func(targetH int, targetW int, moji string) bool
	judgeFunc = func(targetH int, targetW int, moji string) bool {
		if moji == "s" {
			return hwArrayArray[targetH][targetW] == "n"
		} else if moji == "n" {
			return hwArrayArray[targetH][targetW] == "u"
		} else if moji == "u" {
			return hwArrayArray[targetH][targetW] == "k"
		} else if moji == "k" {
			return hwArrayArray[targetH][targetW] == "e"
		} else {
			return hwArrayArray[targetH][targetW] == "s"
		}
	}

	var judgeAlready func(targetH int, targetW int) bool
	judgeAlready = func(targetH int, targetW int) bool {
		return hwAlreadyArrayArray[targetH][targetW] == 100
	}

	var loopFunc func(nowH int, nowW int)
	loopFunc = func(nowH int, nowW int) {
		if nowH == h-1 && nowW == w-1 {
			result = "Yes"
		} else {
			nowMoji := hwArrayArray[nowH][nowW]
			// 上の座標
			if nowH != 0 {
				newH := nowH - 1
				newW := nowW
				judge := judgeFunc(newH, newW, nowMoji) && !judgeAlready(newH, nowW)
				if judge {
					hwAlreadyArrayArray[newH][newW] = 100
					loopFunc(newH, newW)
				}
			}
			// 下の座標
			if nowH != h-1 {
				newH := nowH + 1
				newW := nowW
				judge := judgeFunc(newH, newW, nowMoji) && !judgeAlready(newH, newW)
				if judge {
					hwAlreadyArrayArray[newH][newW] = 100
					loopFunc(newH, newW)
				}
			}
			// 左の座標
			if nowW != 0 {
				newH := nowH
				newW := nowW - 1
				judge := judgeFunc(newH, newW, nowMoji) && !judgeAlready(newH, newW)
				if judge {
					hwAlreadyArrayArray[newH][newW] = 100
					loopFunc(newH, newW)
				}
			}
			// 右の座標
			if nowW != w-1 {
				newH := nowH
				newW := nowW + 1
				judge := judgeFunc(newH, newW, nowMoji) && !judgeAlready(newH, newW)
				if judge {
					hwAlreadyArrayArray[newH][newW] = 100
					loopFunc(newH, newW)
				}
			}
		}

	}
	if hwArrayArray[0][0] != "s" {
		result = "No"
	} else {
		loopFunc(0, 0)
	}

	fmt.Println(result)

}

func DSnukeMazeReadLine(rdr *bufio.Reader) string {
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
