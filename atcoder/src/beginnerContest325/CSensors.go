package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CSensorsMain() {
	var h, w int
	fmt.Scan(&h, &w)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var hwSliceSlice = make([][]string, h)
	for i := 0; i < h; i++ {
		wStrs := strings.Split(CSensorsReadLine(rdr), "")
		hwSliceSlice[i] = wStrs
	}

	hwMap := map[string]int{}
	var judgeFunc func(inputH int, inputW int, censorKey int)
	judgeFunc = func(inputH int, inputW int, censorKey int) {
		// 上
		if inputH > 0 {
			targetH := inputH - 1
			targetW := inputW
			mapKey := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
			_, ok := hwMap[mapKey]
			if hwSliceSlice[targetH][targetW] == "#" && !ok {
				hwMap[mapKey] = censorKey
				judgeFunc(targetH, targetW, censorKey)
			}
		}
		// 左
		if inputW > 0 {
			targetH := inputH
			targetW := inputW - 1
			mapKey := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
			_, ok := hwMap[mapKey]
			if hwSliceSlice[targetH][targetW] == "#" && !ok {
				hwMap[mapKey] = censorKey
				judgeFunc(targetH, targetW, censorKey)
			}
		}
		// 下
		if inputH < h-1 {
			targetH := inputH + 1
			targetW := inputW
			mapKey := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
			_, ok := hwMap[mapKey]
			if hwSliceSlice[targetH][targetW] == "#" && !ok {
				hwMap[mapKey] = censorKey
				judgeFunc(targetH, targetW, censorKey)
			}
		}
		// 右
		if inputW < w-1 {
			targetH := inputH
			targetW := inputW + 1
			mapKey := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
			_, ok := hwMap[mapKey]
			if hwSliceSlice[targetH][targetW] == "#" && !ok {
				hwMap[mapKey] = censorKey
				judgeFunc(targetH, targetW, censorKey)
			}
		}

		// 上左
		if inputH > 0 && inputW > 0 {
			targetH := inputH - 1
			targetW := inputW - 1
			mapKey := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
			_, ok := hwMap[mapKey]
			if hwSliceSlice[targetH][targetW] == "#" && !ok {
				hwMap[mapKey] = censorKey
				judgeFunc(targetH, targetW, censorKey)
			}
		}
		// 上右
		if inputH > 0 && inputW < w-1 {
			targetH := inputH - 1
			targetW := inputW + 1
			mapKey := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
			_, ok := hwMap[mapKey]
			if hwSliceSlice[targetH][targetW] == "#" && !ok {
				hwMap[mapKey] = censorKey
				judgeFunc(targetH, targetW, censorKey)
			}
		}
		// 下左
		if inputH < h-1 && inputW > 0 {
			targetH := inputH + 1
			targetW := inputW - 1
			mapKey := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
			_, ok := hwMap[mapKey]
			if hwSliceSlice[targetH][targetW] == "#" && !ok {
				hwMap[mapKey] = censorKey
				judgeFunc(targetH, targetW, censorKey)
			}
		}
		// 下左
		if inputH < h-1 && inputW < w-1 {
			targetH := inputH + 1
			targetW := inputW + 1
			mapKey := strconv.FormatInt(int64(targetH), 10) + "-" + strconv.FormatInt(int64(targetW), 10)
			_, ok := hwMap[mapKey]
			if hwSliceSlice[targetH][targetW] == "#" && !ok {
				hwMap[mapKey] = censorKey
				judgeFunc(targetH, targetW, censorKey)
			}
		}
	}

	censorKey := 1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			mapKey := strconv.FormatInt(int64(i), 10) + "-" + strconv.FormatInt(int64(j), 10)
			_, ok := hwMap[mapKey]
			if hwSliceSlice[i][j] == "#" && !ok {
				hwMap[mapKey] = censorKey
				judgeFunc(i, j, censorKey)
				censorKey = censorKey + 1
			}
		}
	}

	resultSet := make(map[int]struct{})
	for _, v := range hwMap {
		resultSet[v] = struct{}{}
	}
	fmt.Println(len(resultSet))

}

func CSensorsReadLine(rdr *bufio.Reader) string {
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
