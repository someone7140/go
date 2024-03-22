package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DStringBagsMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	t := DStringBagsReadLine(rdr)
	tLen := len(t)
	n, _ := strconv.Atoi(DStringBagsReadLine(rdr))

	sArrayArray := make([][]string, n)
	for i := 0; i < n; i++ {
		asStrArray := strings.Split(DStringBagsReadLine(rdr), " ")
		a, _ := strconv.Atoi(asStrArray[0])
		var sArray = make([]string, a)
		for j := 1; j <= a; j++ {
			sArray[j-1] = asStrArray[j]
		}
		sArrayArray[i] = sArray
	}

	resultArrayMap := make([]map[string]int, n)

	var judgeMoji func(moji string, fistFlag bool) bool
	judgeMoji = func(moji string, fistFlag bool) bool {
		mojiLen := len(moji)
		if mojiLen == tLen {
			return moji == t
		} else if mojiLen > tLen {
			return false
		} else {
			if fistFlag {
				return strings.HasPrefix(t, moji)
			}
			return true
		}
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			sMap := map[string]int{}

			for _, v := range sArrayArray[0] {
				if judgeMoji(v, true) {
					sMap[v] = 1
				}
			}
			resultArrayMap[0] = sMap
		} else {
			beforeMap := resultArrayMap[i-1]
			newSMap := map[string]int{}

			// まずは普通に追加する
			for _, v := range sArrayArray[i] {
				if judgeMoji(v, false) {
					newSMap[v] = 1
				}
			}

			// 前のも追加する
			for k, v := range beforeMap {

				// 新しいのと文字列結合
				for k2, v2 := range newSMap {
					combine := k + k2
					combineV := v + v2
					v20, ok20 := newSMap[combine]
					v30, ok30 := beforeMap[combine]
					if judgeMoji(combine, true) {
						if !ok20 && !ok30 {
							newSMap[combine] = combineV
						} else if ok20 && ok30 {
							if v20 > combineV && v30 > combineV {
								newSMap[combine] = combineV
							}
						} else if ok20 && v20 > combineV {
							newSMap[combine] = combineV
						} else if ok30 && v30 > combineV {
							newSMap[combine] = combineV

						}
					}

				}

			}

			for k, v := range beforeMap {
				// 前のをそのまま追加
				v2, ok := newSMap[k]
				if judgeMoji(k, true) {
					if !ok {
						newSMap[k] = v
					} else if v2 > v {
						newSMap[k] = v
					}
				}
			}

			resultArrayMap[i] = newSMap
		}

	}

	result := -1
	for k, v := range resultArrayMap[n-1] {
		if k == t {
			if result == -1 {
				result = v
			} else if result > v {
				result = v
			}
		}
	}

	fmt.Println(result)
}

func DStringBagsReadLine(rdr *bufio.Reader) string {
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
