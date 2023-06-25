package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func CIdealSheetMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	// Aのシートを読み込み
	ahw := strings.Split(CIdealSheetReadLine(rdr), " ")
	aH, _ := strconv.Atoi(ahw[0])
	aMap := map[int][]int{}
	aIndex := 0
	for i := 0; i < aH; i++ {
		row := CIdealSheetReadLine(rdr)
		find := strings.Index(row, "#")
		if find == -1 && aIndex == 0 {
			return
		}
		var aColumn []int
		for i2, c := range row {
			cell := string([]rune{c})
			if cell == "#" {
				aColumn = append(aColumn, i2)
			}
		}
		aMap[aIndex] = aColumn
		aIndex = aIndex + 1
	}

	// Bのシートを読み込み
	bhw := strings.Split(CIdealSheetReadLine(rdr), " ")
	bH, _ := strconv.Atoi(bhw[0])
	bMap := map[int][]int{}
	bIndex := 0
	for i := 0; i < bH; i++ {
		row := CIdealSheetReadLine(rdr)
		find := strings.Index(row, "#")
		if find == -1 && bIndex == 0 {
			return
		}
		var bColumn []int
		for i2, c := range row {
			cell := string([]rune{c})
			if cell == "#" {
				bColumn = append(bColumn, i2)
			}
		}
		bMap[bIndex] = bColumn
		bIndex = bIndex + 1
	}

	// Xのシートを読み込み
	xhw := strings.Split(CIdealSheetReadLine(rdr), " ")
	xH, _ := strconv.Atoi(xhw[0])
	xW, _ := strconv.Atoi(xhw[0])
	xMap := map[int][]int{}
	xIndex := 0
	for i := 0; i < xH; i++ {
		row := CIdealSheetReadLine(rdr)
		find := strings.Index(row, "#")
		if find == -1 && xIndex == 0 {
			continue
		}
		if find == -1 && i == xH-1 {
			continue
		}
		var xColumn []int
		for i2, c := range row {
			cell := string([]rune{c})
			if cell == "#" {
				xColumn = append(xColumn, i2)
			}
		}
		xMap[xIndex] = xColumn
		xIndex = xIndex + 1
	}

	result := "No"

	// AからBをマージ箇所
	for i := 0; i < xH; i++ {
		for j := 0; j < xW; j++ {
			mergeMap := map[int][]int{}
			for k, v := range aMap {
				mergeMap[k] = v
			}
			for k, bRow := range bMap {
				tempSet := make(map[int]struct{})
				// aの値を取得
				aArray, okA := mergeMap[k+i]
				if okA {
					for _, vA := range aArray {
						tempSet[vA] = struct{}{}
					}
				}

				// bの値を取得
				for _, vB := range bRow {
					tempSet[vB+j] = struct{}{}
				}
				mergeArray := []int{}
				for vSet := range tempSet {
					mergeArray = append(mergeArray, vSet)
				}
				sort.Ints(mergeArray)
				mergeMap[k+i] = mergeArray
			}
			tempResult := "Yes"
			for k, v := range xMap {
				mergeRow, ok2 := mergeMap[k]
				if !ok2 {
					tempResult = "No"
					break
				} else {
					if !reflect.DeepEqual(v, mergeRow) {
						tempResult = "No"
						break
					}
				}
			}
			if tempResult == "Yes" {
				result = "Yes"
				break
			}
		}
		if result == "Yes" {
			break
		}
	}

	if result == "Yes" {
		fmt.Println("Yes")
	} else {
		result = "No"
		// BからAをマージ箇所
		for i := 0; i < xH; i++ {
			for j := 0; j < xW; j++ {
				mergeMap := map[int][]int{}
				for k, v := range bMap {
					mergeMap[k] = v
				}
				for k, aRow := range aMap {
					tempSet := make(map[int]struct{})
					// bの値を取得
					bArray, okB := mergeMap[k+i]
					if okB {
						for _, vB := range bArray {
							tempSet[vB] = struct{}{}
						}
					}

					// aの値を取得
					for _, vA := range aRow {
						tempSet[vA+j] = struct{}{}
					}
					mergeArray := []int{}
					for vSet := range tempSet {
						mergeArray = append(mergeArray, vSet)
					}
					sort.Ints(mergeArray)
					mergeMap[k+i] = mergeArray
				}
				tempResult := "Yes"
				for k, v := range xMap {
					mergeRow, ok2 := mergeMap[k]
					if !ok2 {
						tempResult = "No"
						break
					} else {
						if !reflect.DeepEqual(v, mergeRow) {
							tempResult = "No"
							break
						}
					}
				}
				if tempResult == "Yes" {
					result = "Yes"
					break
				}
			}
			if result == "Yes" {
				break
			}
		}
		fmt.Println(result)
	}
}

func CIdealSheetReadLine(rdr *bufio.Reader) string {
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
