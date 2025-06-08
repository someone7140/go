package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DDominoCoveringXORMain() {
	var h, w int
	fmt.Scan(&h, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var nowResult int64 = -1
	aListList := make([][]int64, h)
	for i := 0; i < h; i++ {
		aList := make([]int64, w)
		aStrList := strings.Split(DSecurity2Rdr(rdr), " ")
		for j := 0; j < w; j++ {
			a, _ := strconv.ParseInt(aStrList[j], 10, 64)
			if nowResult < 0 {
				nowResult = a
			} else {
				nowResult = nowResult ^ a
			}
			aList[j] = a
		}
		aListList[i] = aList
	}

	result := nowResult
	aDominoSet := make(map[string]struct{})

	var loopFunc func(targetH int, targetW int, tempResult int64)
	loopFunc = func(targetH int, targetW int, tempResult int64) {
		if targetH == h-1 && targetW == w-1 {
			// 終了
		} else {
			key := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
			_, ok := aDominoSet[key]
			// すでにドミノがある場合
			if ok {
				// 下に行く
				if targetH < h-1 {
					loopFunc(targetH+1, targetW, tempResult)
				}
				// 右に行く
				if targetW < w-1 {
					loopFunc(targetH, targetW+1, tempResult)
				}
			} else {
				// 置かない場合
				// 下に行く
				if targetH < h-1 {
					loopFunc(targetH+1, targetW, tempResult)
				}
				// 右に行く
				if targetW < w-1 {
					loopFunc(targetH, targetW+1, tempResult)
				}

				// 置く場合

				// 下に置く
				if targetH < h-1 {
					aDominoSet[key] = struct{}{}
					// 該当の値でXOR
					tempResult2 := tempResult ^ aListList[targetH][targetW]

					key2 := strconv.FormatInt(int64(targetH+1), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
					aDominoSet[key2] = struct{}{}
					tempResult2 = tempResult2 ^ aListList[targetH+1][targetW]
					if tempResult2 > result {
						result = tempResult2
					}
					// 下に行く
					if targetH < h-1 {
						loopFunc(targetH+1, targetW, tempResult2)
					}
					// 右に行く
					if targetW < w-1 {
						loopFunc(targetH, targetW+1, tempResult2)
					}

					// 戻す
					delete(aDominoSet, key)
					delete(aDominoSet, key2)

				}

				// 右に置く
				if targetW < w-1 {
					aDominoSet[key] = struct{}{}
					// 該当の値でXOR
					tempResult2 := tempResult ^ aListList[targetH][targetW]

					key2 := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW+1), 10)
					aDominoSet[key2] = struct{}{}
					tempResult2 = tempResult2 ^ aListList[targetH][targetW+1]
					if tempResult2 > result {
						result = tempResult2
					}
					// 下に行く
					if targetH < h-1 {
						loopFunc(targetH+1, targetW, tempResult2)
					}
					// 右に行く
					if targetW < w-1 {
						loopFunc(targetH, targetW+1, tempResult2)
					}

					// 戻す
					delete(aDominoSet, key)
					delete(aDominoSet, key2)
				}
			}
		}

	}

	loopFunc(0, 0, nowResult)
	fmt.Println(result)
}

func DSecurity2Rdr(rdr *bufio.Reader) string {
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
