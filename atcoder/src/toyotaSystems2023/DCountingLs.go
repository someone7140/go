package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	rowKeyMap := map[int][]int{}
	var column0Array []int
	var column1Array []int
	var column2Array []int

	var rows []int
	for i := 0; i < n; i++ {
		rowFindFlag := false
		sStrs := strings.Split(DCountingLsReadLine(rdr), "")
		for j := 0; j < 3; j++ {
			if sStrs[j] == "o" {
				if !rowFindFlag {
					rows = append(rows, i)
					rowFindFlag = true
				}
				// 行に追加
				rows, ok := rowKeyMap[i]
				if !ok {
					rowKeyMap[i] = []int{j}
				} else {
					rowKeyMap[i] = append(rows, j)
				}
				// 列に追加
				if j == 0 {
					column0Array = append(column0Array, i)
				} else if j == 1 {
					column1Array = append(column1Array, i)
				} else if j == 2 {
					column2Array = append(column2Array, i)
				}
			}
		}
	}
	result := int64(0)

	var calcFunc2 func(rowKey int, columns []int)
	calcFunc2 = func(rowKey int, columns []int) {
		lenColumn := len(columns)
		if lenColumn > 0 && columns[0] <= rowKey {
			lenColumn = lenColumn - 1
		}
		result = result + int64(lenColumn)
	}

	var calcFunc func(rowKey int, column1 int, column2 int)
	calcFunc = func(rowKey int, column1 int, column2 int) {
		if column1 == 0 {
			calcFunc2(rowKey, column0Array)
		}
		if column1 == 1 {
			calcFunc2(rowKey, column1Array)
		}
		if column1 == 2 {
			calcFunc2(rowKey, column2Array)
		}
		if column2 == 0 {
			calcFunc2(rowKey, column0Array)
		}
		if column2 == 1 {
			calcFunc2(rowKey, column1Array)
		}
		if column2 == 2 {
			calcFunc2(rowKey, column2Array)
		}
	}

	// 行で回す
	for _, rowKey := range rows {
		rowValue := rowKeyMap[rowKey]
		rowValueLen := len(rowValue)
		if rowValueLen == 2 {
			calcFunc(rowKey, rowValue[0], rowValue[1])
		} else if rowValueLen == 3 {
			calcFunc(rowKey, rowValue[0], rowValue[1])
			calcFunc(rowKey, rowValue[0], rowValue[2])
			calcFunc(rowKey, rowValue[1], rowValue[2])
		}
		// カラム情報を削除
		lenColumn0Array := len(column0Array)
		if lenColumn0Array > 0 {
			if column0Array[0] == rowKey {
				if lenColumn0Array == 1 {
					column0Array = []int{}
				} else {
					column0Array = column0Array[1:]
				}
			}
		}
		lenColumn1Array := len(column1Array)
		if lenColumn1Array > 0 {
			if lenColumn1Array == 1 {
				column1Array = []int{}
			} else {
				column1Array = column1Array[1:]
			}
		}
		lenColumn2Array := len(column2Array)
		if lenColumn2Array > 0 {
			if lenColumn2Array == 1 {
				column2Array = []int{}
			} else {
				column2Array = column2Array[1:]
			}
		}
	}
	fmt.Println(result)

}

func DCountingLsReadLine(rdr *bufio.Reader) string {
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
