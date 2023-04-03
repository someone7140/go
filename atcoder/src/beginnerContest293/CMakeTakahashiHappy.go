package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CMakeTakahashiHappyMain() {
	var h, w int
	fmt.Scan(&h, &w)

	aArrayArray := make([][]int, h)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	result := int64(0)

	for i := 0; i < h; i++ {
		aWArrayStr := strings.Split(CMakeTakahashiHappyReadLine(rdr), " ")
		aWArray := make([]int, w)
		for j := 0; j < w; j++ {
			aWArray[j], _ = strconv.Atoi(aWArrayStr[j])
		}
		aArrayArray[i] = aWArray
	}

	copyMap := func(mapParam map[int]bool) map[int]bool {
		copiedMap := map[int]bool{}
		for key, value := range mapParam {
			if value {
				copiedMap[key] = value
			}
		}
		return copiedMap
	}

	var loopFunc func(hParam, wParam int, aVisitedMap map[int]bool)
	loopFunc = func(hParam, wParam int, aVisitedMap map[int]bool) {
		if wParam == (w-1) && hParam == (h-1) {
			result = result + 1
		} else {
			if wParam < w-1 {
				nextW := wParam + 1
				res, ok := aVisitedMap[aArrayArray[hParam][nextW]]
				if !res && !ok {
					copyMap := copyMap(aVisitedMap)
					copyMap[aArrayArray[hParam][nextW]] = true
					loopFunc(hParam, nextW, copyMap)
				}
			}
			if hParam < h-1 {
				nextH := hParam + 1
				res, ok := aVisitedMap[aArrayArray[nextH][wParam]]
				if !res && !ok {
					copyMap := copyMap(aVisitedMap)
					copyMap[aArrayArray[nextH][wParam]] = true
					loopFunc(nextH, wParam, copyMap)
				}

			}
		}
	}

	if aArrayArray[0][0] != aArrayArray[1][0] {
		aVisitedMap := map[int]bool{}
		aVisitedMap[aArrayArray[0][0]] = true
		aVisitedMap[aArrayArray[1][0]] = true
		loopFunc(1, 0, aVisitedMap)
	}
	if aArrayArray[0][0] != aArrayArray[0][1] {
		aVisitedMap := map[int]bool{}
		aVisitedMap[aArrayArray[0][0]] = true
		aVisitedMap[aArrayArray[0][1]] = true
		loopFunc(0, 1, aVisitedMap)
	}

	fmt.Println(result)
}

func CMakeTakahashiHappyReadLine(rdr *bufio.Reader) string {
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
