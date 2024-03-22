package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DTilingAb struct {
	a int
	b int
}

func main() {
	var n, h, w int
	fmt.Scan(&n, &h, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var abArray = make([]DTilingAb, n)
	for i := 0; i < n; i++ {
		ab := strings.Split(DTilingReadLine(rdr), " ")
		a, _ := strconv.Atoi(ab[0])
		b, _ := strconv.Atoi(ab[1])
		abArray[i] = DTilingAb{
			a, b,
		}
	}

	tileMenseki := h * w
	var abArrayArray [][]int
	var mensekiFunc func(abList []int, menseki int, nextIndex int)
	mensekiFunc = func(abList []int, menseki int, nextIndex int) {
		copyAbList := make([]int, len(abList))
		copy(copyAbList, abList)
		if nextIndex < n {
			// 1個追加
			next := abArray[nextIndex]
			tsuika := next.a * next.b
			nextArray := append(abList, nextIndex)
			if menseki+tsuika == tileMenseki {
				abArrayArray = append(abArrayArray, nextArray)
			}
			mensekiFunc(nextArray, menseki+tsuika, nextIndex+1)
			// そのまま渡す
			mensekiFunc(copyAbList, menseki, nextIndex+1)

		}
	}
	mensekiFunc([]int{}, 0, 0)

	sliceCopy := func(in, out interface{}) {
		buf := new(bytes.Buffer)
		gob.NewEncoder(buf).Encode(in)
		gob.NewDecoder(buf).Decode(out)
	}

	result := "No"
	var haichiFunc func(tileIndexList []int, abUmetaList [][]int, nextIndex int)
	haichiFunc = func(tileIndexList []int, abUmetaList [][]int, nextIndex int) {
		if result == "Yes" {
			return
		}
		fmt.Println(abUmetaList)

		tileLen := len(tileIndexList)
		ab := abArray[tileIndexList[nextIndex]]
		kuurannAru := false
		for i1 := 0; i1 < h; i1++ {
			for j1 := 0; j1 < w; j1++ {
				if abUmetaList[i1][j1] != 1 {
					kuurannAru = true
					var abUmetaListCopy [][]int
					sliceCopy(abUmetaList, &abUmetaListCopy)
					// そこを起点に塗りつぶせるか
					yZan := ab.b
					umeresult := true
					for i2 := i1; i2 < h; i2++ {
						if yZan != 0 && umeresult {
							xZan := ab.a
							for j2 := j1; j2 < w; j2++ {
								if abUmetaListCopy[i2][j2] == 1 {
									break
								}
								if xZan != 0 {
									abUmetaListCopy[i2][j2] = 1
									xZan = xZan - 1
								}
							}
							if xZan == 0 {
								yZan = yZan - 1
							} else {
								umeresult = false
							}
						}
					}
					if yZan == 0 && umeresult && nextIndex != tileLen-1 {
						haichiFunc(tileIndexList, abUmetaListCopy, nextIndex+1)
					}

					// 逆にする
					var abUmetaListCopy2 [][]int
					sliceCopy(abUmetaList, &abUmetaListCopy2)
					yZan2 := ab.a
					umeresult2 := true
					for i3 := i1; i3 < h; i3++ {
						if yZan != 0 && umeresult {
							xZan2 := ab.b
							for j3 := j1; j3 < w; j3++ {
								if abUmetaListCopy[i3][j3] == 1 {
									break
								}
								if xZan2 != 0 {
									abUmetaListCopy2[i3][j3] = 1
									xZan2 = xZan2 - 1
								}
							}
							if xZan2 == 0 {
								yZan2 = yZan2 - 1
							} else {
								umeresult2 = false
							}
						}
					}
					if yZan2 == 0 && umeresult2 && nextIndex != tileLen-1 {
						haichiFunc(tileIndexList, abUmetaListCopy, nextIndex+1)
					}
				}
			}
		}
		if !kuurannAru {
			result = "Yes"
		}
	}

	for _, v := range abArrayArray {
		var abUmetaList = make([][]int, h)
		for i := 0; i < h; i++ {
			var abUmetaListRow = make([]int, w)
			abUmetaList[i] = abUmetaListRow
		}
		haichiFunc(v, abUmetaList, 0)
	}

	fmt.Println(result)
}

func DTilingReadLine(rdr *bufio.Reader) string {
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
