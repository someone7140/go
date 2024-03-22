package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func COneTimeSwapMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := COneTimeSwapReadLine(rdr)
	sMap := map[string][]int{}

	for i, c := range s {
		sMoji := string([]rune{c})
		indexList, ok := sMap[sMoji]
		if !ok {
			sMap[sMoji] = []int{i}
		} else {
			sMap[sMoji] = append(indexList, i)
		}
	}
	var result int64
	keys := reflect.ValueOf(sMap).MapKeys()
	lenLey := len(keys)
	ownMoveFlag := false

	for i := 0; i < lenLey; i++ {
		key1 := keys[i].String()
		valList1 := sMap[key1]
		valList1Len := len(valList1)
		if len(valList1) > 1 && !ownMoveFlag {
			result = result + 1
			ownMoveFlag = true
		}
		if i < lenLey-1 {
			for j := i + 1; j < lenLey; j++ {
				key2 := keys[j].String()
				valList2 := sMap[key2]
				valList2Len := len(valList2)
				result = result + int64(valList1Len*valList2Len)
			}
		}

	}

	fmt.Println(result)
}

func COneTimeSwapReadLine(rdr *bufio.Reader) string {
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
