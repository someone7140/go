package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type CBeltConveyorRdrXy struct {
	x int
	y int
}

func CBeltConveyorMain() {
	var h, w int
	fmt.Scan(&h, &w)

	hwArray := make([][]string, h)
	goneMap := map[string]bool{}
	result := "-1"

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < h; i++ {
		gLine := CBeltConveyorRdr(rdr)
		gArray := make([]string, w)
		for j, c := range gLine {
			gMoji := string([]rune{c})
			gArray[j] = gMoji
		}
		hwArray[i] = gArray
	}

	nowPos := CBeltConveyorRdrXy{x: 0, y: 0}
	for {
		nowPosStr := strconv.FormatInt(int64(nowPos.x), 10) + " " + strconv.FormatInt(int64(nowPos.y), 10)
		_, ok := goneMap[nowPosStr]
		if ok {
			break
		} else {
			goneMap[nowPosStr] = true
		}

		nowG := hwArray[nowPos.y][nowPos.x]

		if nowG == "U" {
			if nowPos.y != 0 {
				nowPos = CBeltConveyorRdrXy{
					x: nowPos.x,
					y: nowPos.y - 1,
				}
			} else {
				result = strconv.FormatInt(int64(nowPos.y+1), 10) + " " + strconv.FormatInt(int64(nowPos.x+1), 10)
				break
			}
		}
		if nowG == "D" {
			if nowPos.y != h-1 {
				nowPos = CBeltConveyorRdrXy{
					x: nowPos.x,
					y: nowPos.y + 1,
				}
			} else {
				result = strconv.FormatInt(int64(nowPos.y+1), 10) + " " + strconv.FormatInt(int64(nowPos.x+1), 10)
				break
			}
		}
		if nowG == "L" {
			if nowPos.x != 0 {
				nowPos = CBeltConveyorRdrXy{
					x: nowPos.x - 1,
					y: nowPos.y,
				}
			} else {
				result = strconv.FormatInt(int64(nowPos.y+1), 10) + " " + strconv.FormatInt(int64(nowPos.x+1), 10)
				break
			}
		}
		if nowG == "R" {
			if nowPos.x != w-1 {
				nowPos = CBeltConveyorRdrXy{
					x: nowPos.x + 1,
					y: nowPos.y,
				}
			} else {
				result = strconv.FormatInt(int64(nowPos.y+1), 10) + " " + strconv.FormatInt(int64(nowPos.x+1), 10)
				break
			}
		}
	}

	fmt.Println(result)
}

func CBeltConveyorRdr(rdr *bufio.Reader) string {
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
