package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CFlappingTakahashiTlu struct {
	time  int64
	lower int64
	upper int64
}

func main() {
	var t int
	fmt.Scan(&t)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var resultSlice []string

	for i := 0; i < t; i++ {
		nhArray := strings.Split(CFlappingTakahashiRdr(rdr), " ")
		n, _ := strconv.Atoi(nhArray[0])
		h, _ := strconv.ParseInt(nhArray[1], 10, 64)
		time := int64(0)
		var tluArray []CFlappingTakahashiTlu
		for j := 0; j < n; j++ {
			tluInputs := strings.Split(CFlappingTakahashiRdr(rdr), " ")
			t, _ := strconv.ParseInt(tluInputs[0], 10, 64)
			l, _ := strconv.ParseInt(tluInputs[1], 10, 64)
			u, _ := strconv.ParseInt(tluInputs[2], 10, 64)
			tluArray = append(tluArray, CFlappingTakahashiTlu{
				time:  t,
				lower: l,
				upper: u,
			})
		}

		result := "Yes"
		minH := h
		maxH := h
		for j := 0; j < n; j++ {
			tlu := tluArray[j]
			timeSabun := tlu.time - time
			if minH > tlu.upper {
				if timeSabun < (minH - tlu.upper) {
					result = "No"
					break
				}
				maxH = tlu.upper
				timeH := minH - timeSabun
				if timeH < tlu.lower {
					minH = tlu.lower
				} else {
					minH = timeH
				}
			} else if maxH < tlu.lower {
				if timeSabun < (tlu.lower - maxH) {
					result = "No"
					break
				}
				minH = tlu.lower
				timeH := maxH + timeSabun
				if timeH > tlu.upper {
					maxH = tlu.upper
				} else {
					maxH = timeH
				}
			} else {
				if tlu.upper > maxH {
					timeH := maxH + timeSabun
					if timeH > tlu.upper {
						maxH = tlu.upper
					} else {
						maxH = timeH
					}
				} else {
					maxH = tlu.upper
				}

				if tlu.lower <= minH {
					timeH := minH - timeSabun
					if timeH > tlu.lower {
						minH = timeH
					} else {
						minH = tlu.lower
					}
				} else {
					minH = tlu.lower
				}
			}
			time = tlu.time
			// fmt.Println("time:" + strconv.FormatInt(time, 10))
			// fmt.Println("minH:" + strconv.FormatInt(minH, 10))
			// fmt.Println("maxH:" + strconv.FormatInt(maxH, 10))
		}
		resultSlice = append(resultSlice, result)
	}
	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CFlappingTakahashiRdr(rdr *bufio.Reader) string {
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
