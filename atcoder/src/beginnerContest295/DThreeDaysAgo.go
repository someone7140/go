package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DThreeDaysAgoMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := DThreeDaysAgoReadLine(rdr)
	var kisuCountSliceSlice [][]int
	var result int64
	sMap := map[int]int{}
	for _, c := range s {
		sMoji := string([]rune{c})
		sInt, _ := strconv.Atoi(sMoji)
		var newKisuCountSliceSlice [][]int
		kisuCountSlice := []int{sInt}
		v, ok := sMap[sInt]
		if ok {
			sMap[sInt] = v + 1
		} else {
			sMap[sInt] = 1
		}
		newKisuCountSliceSlice = append(newKisuCountSliceSlice, kisuCountSlice)
		for _, kisuCountSlice := range kisuCountSliceSlice {
			newKisuCountSlice := []int{}
			sIntFlag := false
			for _, kisuu := range kisuCountSlice {
				if kisuu != sInt {
					newKisuCountSlice = append(newKisuCountSlice, kisuu)
				} else {
					sIntFlag = true
				}
			}
			if len(newKisuCountSlice) == 0 {
				result = result + 1
			} else {
				if sIntFlag {
					newKisuCountSliceSlice = append(newKisuCountSliceSlice, newKisuCountSlice)
				} else {
					newKisuCountSlice = append(newKisuCountSlice, sInt)
					newKisuCountSliceSlice = append(newKisuCountSliceSlice, newKisuCountSlice)
				}
			}
		}
		kisuCountSliceSlice = newKisuCountSliceSlice
	}
	plusFlag := true
	for _, v := range sMap {
		if v%2 != 0 {
			plusFlag = false
			break
		}
	}
	if plusFlag {
		result = result + 1
	}
	fmt.Println(result)

}

func DThreeDaysAgoReadLine(rdr *bufio.Reader) string {
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
