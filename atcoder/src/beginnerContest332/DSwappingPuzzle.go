package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func DSwappingPuzzleMain() {
	var h, w int
	fmt.Scan(&h, &w)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var aArrayArray [][]int
	var bArrayArray [][]int

	// aの配列を格納
	for i := 0; i < h; i++ {
		aArray := make([]int, w)
		aArrayStrs := strings.Split(DSwappingPuzzleReadLine(rdr), " ")
		for j, aMoji := range aArrayStrs {
			a, _ := strconv.Atoi(aMoji)
			aArray[j] = a
		}
		aArrayArray = append(aArrayArray, aArray)
	}
	// bの配列を格納
	for i := 0; i < h; i++ {
		bArray := make([]int, w)
		bArrayStrs := strings.Split(DSwappingPuzzleReadLine(rdr), " ")
		for j, bMoji := range bArrayStrs {
			b, _ := strconv.Atoi(bMoji)
			bArray[j] = b
		}
		bArrayArray = append(bArrayArray, bArray)
	}

	makeCopy := func(nums []int) []int {
		return append([]int{}, nums...)
	}

	factorial := func(n int) int {
		ret := 1
		for i := 2; i <= n; i++ {
			ret *= i
		}
		return ret
	}

	permute := func(nums []int, ret *[][]int) {
		*ret = append(*ret, makeCopy(nums))

		n := len(nums)
		p := make([]int, n+1)
		for i := 0; i < n+1; i++ {
			p[i] = i
		}
		for i := 1; i < n; {
			p[i]--
			j := 0
			if i%2 == 1 {
				j = p[i]
			}

			nums[i], nums[j] = nums[j], nums[i]
			*ret = append(*ret, makeCopy(nums))
			for i = 1; p[i] == 0; i++ {
				p[i] = i
			}
		}
	}

	Permute := func(nums []int) [][]int {
		n := factorial(len(nums))
		ret := make([][]int, 0, n)
		permute(nums, &ret)
		return ret
	}

	// 行の連番愛列
	hRenbans := make([]int, h)
	for i, _ := range hRenbans {
		hRenbans[i] = i
	}
	// 行の全パターン
	hRenbanspermute := Permute(hRenbans)

	// 列の連番愛列
	wRenbans := make([]int, w)
	for i, _ := range wRenbans {
		wRenbans[i] = i
	}
	// 行の全パターン
	wRenbanspermute := Permute(wRenbans)
	result := 1000000000000000000

	sliceCopy := func(in, out interface{}) {
		buf := new(bytes.Buffer)
		gob.NewEncoder(buf).Encode(in)
		gob.NewDecoder(buf).Decode(out)
	}

	for _, hRenban := range hRenbanspermute {
		tempResult := 0
		var tempHArrayArray [][]int
		var tempWArrayArrayEmpty [][]int
		for _, index := range hRenban {
			var tempArray = make([]int, w)
			copy(tempArray, aArrayArray[index])
			tempHArrayArray = append(tempHArrayArray, tempArray)
			tempWArrayArrayEmpty = append(tempWArrayArrayEmpty, make([]int, w))
		}

		for _, wRenban := range wRenbanspermute {
			var tempWArrayArray [][]int
			sliceCopy(tempWArrayArrayEmpty, &tempWArrayArray)
			for j, index := range wRenban {
				for i := 0; i < h; i++ {
					tempWArrayArray[i][j] = tempHArrayArray[i][index]
				}
			}

			if reflect.DeepEqual(bArrayArray, tempWArrayArray) {
				fmt.Println("aaaaaaa")
				if result > tempResult {
					result = tempResult
				}
			}
		}

	}

	fmt.Println(result)

}

func DSwappingPuzzleReadLine(rdr *bufio.Reader) string {
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
