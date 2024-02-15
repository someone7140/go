package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DSuperTakahashiBros struct {
	a int
	b int
	x int
}

func main() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var abxArray []DSuperTakahashiBros
	for i := 1; i < n; i++ {
		abxStrArray := strings.Split(DSuperTakahashiBrosReadLine(rdr), " ")
		a, _ := strconv.Atoi(abxStrArray[0])
		b, _ := strconv.Atoi(abxStrArray[1])
		x, _ := strconv.Atoi(abxStrArray[2])
		abxArray = append(abxArray, DSuperTakahashiBros{
			a: a,
			b: b,
			x: x,
		})
	}

	tennsuuMap := map[int]int64{}
	tennsuuMap[1] = 0
	var stageFunc func(target int, nowTensuu int64)
	stageFunc = func(target int, nowTensuu int64) {
		if target == n {
			return
		}

		nV, okN := tennsuuMap[n]
		// 普通に次に行く
		nextV, ok := tennsuuMap[target+1]
		calcNext := nowTensuu + int64(abxArray[target-1].a)
		if !okN || (okN && nV > calcNext) {
			if !ok {
				tennsuuMap[target+1] = calcNext
				stageFunc(target+1, calcNext)
			} else {
				if nextV > (nowTensuu + int64(abxArray[target-1].a)) {
					tennsuuMap[target+1] = calcNext
					stageFunc(target+1, calcNext)
				}
			}
		}

		// bとxを使う
		skipV, ok2 := tennsuuMap[abxArray[target-1].x]
		calcSkip := nowTensuu + int64(abxArray[target-1].b)
		if !okN || (okN && nV > calcSkip) {
			if !ok2 {
				tennsuuMap[abxArray[target-1].x] = calcSkip
				stageFunc(abxArray[target-1].x, calcSkip)
			} else {
				if skipV > (nowTensuu + int64(abxArray[target-1].b)) {
					tennsuuMap[abxArray[target-1].x] = calcSkip
					stageFunc(abxArray[target-1].x, calcSkip)
				}
			}
		}

	}
	stageFunc(1, 0)
	fmt.Println(tennsuuMap[n])
}

func DSuperTakahashiBrosReadLine(rdr *bufio.Reader) string {
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
