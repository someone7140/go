package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func DLRUDInstructionsMain() {
	var h, w, r, c int
	fmt.Scan(&h, &w, &r, &c)
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var hToWMap = map[int][]int{}
	var wToHMap = map[int][]int{}
	for i := 0; i < n; i++ {
		rcStrArray := strings.Split(DLRUDInstructionsReadLine(rdr), " ")
		r, _ := strconv.Atoi(rcStrArray[0])
		c, _ := strconv.Atoi(rcStrArray[1])

		wArray, okH := hToWMap[r]
		if okH {
			wArray = append(wArray, c)
			sort.Ints(wArray)
			hToWMap[r] = wArray
		} else {
			hToWMap[r] = []int{c}
		}

		hArray, okW := wToHMap[c]
		if okW {
			hArray = append(hArray, r)
			sort.Ints(hArray)
			wToHMap[r] = hArray
		} else {
			wToHMap[r] = []int{r}
		}
	}

	q, _ := strconv.Atoi(DLRUDInstructionsReadLine(rdr))
	nowH := r
	nowW := c
	var resultSlice []string
	for i := 0; i < q; i++ {
		dlStrArray := strings.Split(DLRUDInstructionsReadLine(rdr), " ")
		d := dlStrArray[0]
		l, _ := strconv.Atoi(dlStrArray[1])

		if d == "L" || d == "R" {
			tempW := nowW
			if d == "L" {
				tempW = tempW - l
			} else {
				tempW = tempW + l
			}
			wArray, ok := hToWMap[nowH]
			if !ok {
				if tempW < 1 {
					nowW = 1
				} else if tempW > w {
					nowW = w
				} else {
					nowW = tempW
				}
			} else {
				wArrayLen := len(wArray)
				minW := wArray[0]
				maxW := wArray[wArrayLen-1]
				if minW > nowW {
					if d == "L" {
						if tempW < 1 {
							nowW = 1
						} else {
							nowW = tempW
						}
					} else {
						if tempW < minW {
							nowW = tempW
						} else {
							nowW = minW - 1
						}
					}
				} else if maxW < nowW {
					if d == "R" {
						if tempW > w {
							nowW = w
						} else {
							nowW = tempW
						}
					} else {
						if tempW > maxW {
							nowW = tempW
						} else {
							nowW = maxW + 1
						}
					}
				} else {

				}

			}
		} else {

		}
		resultSlice = append(resultSlice, strconv.FormatInt(int64(nowH), 10)+" "+strconv.FormatInt(int64(nowW), 10))
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func DLRUDInstructionsReadLine(rdr *bufio.Reader) string {
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
