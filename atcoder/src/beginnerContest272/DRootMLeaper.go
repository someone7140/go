package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var nSliceSlice [][]int
var sqrtSliceSlice [][]int
var sqrtSliceLen int
var n, m int

func main() {

	fmt.Scan(&n, &m)

	nSliceSlice = make([][]int, n)

	for i := 0; i < n; i++ {
		var nSlice = make([]int, n)
		for i := 0; i < n; i++ {
			nSlice[i] = -1
		}
		nSliceSlice[i] = nSlice
	}
	nSliceSlice[0][0] = 0

	sqrtM := int(math.Sqrt(float64(m))) + 1
	for i := 0; i <= sqrtM; i++ {
		for j := i; j <= sqrtM; j++ {
			ij := i*i + j*j
			if ij > m {
				break
			} else if ij == m {
				var sqrtSlice = make([]int, 2)
				sqrtSlice[0] = i
				sqrtSlice[1] = j
				sqrtSliceSlice = append(sqrtSliceSlice, sqrtSlice)
			}
		}
	}
	sqrtSliceLen = len(sqrtSliceSlice)
	if sqrtSliceLen > 0 {
		tansaku(0, 0, 1)
	}

	for i := 0; i < n; i++ {
		resultSlice := make([]string, n)

		for j := 0; j < n; j++ {
			resultSlice[j] = strconv.FormatInt(int64(nSliceSlice[i][j]), 10)
		}
		fmt.Println(strings.Join(resultSlice, " "))
	}

}

func tansaku(yoko int, tate int, count int) {
	for k := 0; k < sqrtSliceLen; k++ {
		i := sqrtSliceSlice[k][0]
		j := sqrtSliceSlice[k][1]

		if yoko+i < n && tate+j < n {
			next := nSliceSlice[yoko+i][tate+j]
			if next == -1 || next > count {
				nSliceSlice[yoko+i][tate+j] = count
				tansaku(yoko+i, tate+j, count+1)
			}
		}
		if yoko-i > -1 && tate-j > -1 {
			next := nSliceSlice[yoko-i][tate-j]
			if next == -1 || next > count {
				nSliceSlice[yoko-i][tate-j] = count
				tansaku(yoko-i, tate-j, count+1)
			}
		}
		if yoko-i > -1 && tate+j < n {
			next := nSliceSlice[yoko-i][tate+j]
			if next == -1 || next > count {
				nSliceSlice[yoko-i][tate+j] = count
				tansaku(yoko-i, tate+j, count+1)
			}
		}
		if yoko+i < n && tate-j > 0 {
			next := nSliceSlice[yoko+i][tate-j]
			if next == -1 || next > count {
				nSliceSlice[yoko+i][tate-j] = count
				tansaku(yoko+i, tate-j, count+1)
			}
		}

		if i != j {
			if yoko+j < n && tate+i < n {
				next := nSliceSlice[yoko+j][tate+i]
				if next == -1 || next > count {
					nSliceSlice[yoko+j][tate+i] = count
					tansaku(yoko+j, tate+i, count+1)
				}
			}
			if yoko-j > -1 && tate-i > -1 {
				next := nSliceSlice[yoko-j][tate-i]
				if next == -1 || next > count {
					nSliceSlice[yoko-j][tate-i] = count
					tansaku(yoko-j, tate-i, count+1)
				}
			}
			if yoko-j > -1 && tate+i < n {
				next := nSliceSlice[yoko-j][tate+i]
				if next == -1 || next > count {
					nSliceSlice[yoko-j][tate+i] = count
					tansaku(yoko-j, tate+i, count+1)
				}
			}
			if yoko+j < n && tate-i > 0 {
				next := nSliceSlice[yoko+j][tate-i]
				if next == -1 || next > count {
					nSliceSlice[yoko+j][tate-i] = count
					tansaku(yoko+j, tate-i, count+1)
				}
			}
		}
	}
}
