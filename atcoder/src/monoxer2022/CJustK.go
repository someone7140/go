package main

import (
	"bufio"
	"fmt"
	"os"
)

var result = 0
var n, k int
var sMapArray = []map[string]int{}

func main() {
	fmt.Scan(&n, &k)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < n; i++ {
		sMap := map[string]int{}
		s := CJustKReadLine(rdr)
		for _, c := range s {
			sMoji := string([]rune{c})
			v, ok := sMap[sMoji]
			if ok {
				sMap[sMoji] = v + 1
			} else {
				sMap[sMoji] = 1
			}
		}
		sMapArray = append(sMapArray, sMap)
	}

	hantei(0, []map[string]int{})
	fmt.Println(result)
}

func hantei(index int, kaisuuMapArray []map[string]int) {
	if index < n {
		newKaisuuMapArray := []map[string]int{}
		for _, m := range kaisuuMapArray {
			newKaisuuMapArray = append(newKaisuuMapArray, m)
			newMap := hantei2(index, m)
			newKaisuuMapArray = append(newKaisuuMapArray, newMap)
		}
		initialMap := hantei2(index, map[string]int{})
		newKaisuuMapArray = append(newKaisuuMapArray, initialMap)
		hantei(index+1, newKaisuuMapArray)
	}
}

func hantei2(index int, kaisuuMap map[string]int) map[string]int {
	// 判定
	sMapIndex := sMapArray[index]
	newKaisuuMap := map[string]int{}

	for key, v := range kaisuuMap {
		newKaisuuMap[key] = v
	}

	for key, v := range sMapIndex {
		kaisuuMapGet, ok := kaisuuMap[key]
		newKaisuu := v
		if ok {
			newKaisuu = v + kaisuuMapGet
		}
		newKaisuuMap[key] = newKaisuu
	}

	newKaisuuMap2 := map[int]int{}
	for _, v := range newKaisuuMap {
		v2, ok := newKaisuuMap2[v]
		if ok {
			newKaisuuMap2[v] = v2 + 1
		} else {
			newKaisuuMap2[v] = 1
		}
	}
	for key, v := range newKaisuuMap2 {
		if key == k {
			if result < v {
				result = v
			}
		}
	}
	return newKaisuuMap
}
func CJustKReadLine(rdr *bufio.Reader) string {
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
