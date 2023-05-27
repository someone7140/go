package main

import (
	"bufio"
	"fmt"
	"os"
)

func DShiftvsCapsLockMain() {
	var x, y, z int64
	fmt.Scan(&x, &y, &z)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := DShiftvsCapsLockReadLine(rdr)
	sLen := len(s)
	sArrayArray := make([][]int64, sLen)
	var result int64
	result = 0

	for i, c := range s {
		sWord := string([]rune{c})
		sArray := make([]int64, 2)
		if i == 0 {
			if sWord == "a" {
				// 普通に入れる
				nonCapstime1 := x
				// CapsLock → Shift
				capsTime1 := z + y
				// CapsLock → Shift → CapsLock
				nonCapstime2 := z + y + z
				sArray[1] = capsTime1
				if nonCapstime1 < nonCapstime2 {
					sArray[0] = nonCapstime1
				} else {
					sArray[0] = nonCapstime2
				}
			} else {
				// shift
				nonCapstime1 := y
				// CapsLock → 普通
				capsTime1 := z + x

				sArray[0] = nonCapstime1
				sArray[1] = capsTime1
			}
			sArrayArray[i] = sArray
		} else {
			beforeArray := sArrayArray[i-1]
			var minNonCaps int64
			var minCaps int64

			if sWord == "a" {
				// 前がCaps→Caps
				// shift
				minCaps = y + beforeArray[1]
				// CapsLock → 普通 → CapsLock
				tempCapstime1 := z + x + z + beforeArray[1]
				if tempCapstime1 < minCaps {
					minCaps = tempCapstime1
				}

				// 前がCaps→noCaps
				// CapsLock → 普通
				minNonCaps = z + x + beforeArray[1]
				// shift ⇨ CapsLock
				tempNoCapstime1 := y + z + beforeArray[1]
				if tempNoCapstime1 < minNonCaps {
					minNonCaps = tempNoCapstime1
				}

				// 前がnoCaps→Caps
				// 普通 → CapsLock
				tempCapstime2 := z + x + beforeArray[0]
				if tempCapstime2 < minCaps {
					minCaps = tempCapstime2
				}
				// CapsLock → shift
				tempCapstime3 := z + y + beforeArray[0]
				if tempCapstime3 < minCaps {
					minCaps = tempCapstime3
				}

				// 前がnoCaps→noCaps
				// 普通
				tempNoCapstime2 := x + beforeArray[0]
				if tempNoCapstime2 < minNonCaps {
					minNonCaps = tempNoCapstime2
				}
				// CapsLock → shift → CapsLock
				tempNoCapstime3 := z + y + z + beforeArray[0]
				if tempNoCapstime3 < minNonCaps {
					minNonCaps = tempNoCapstime3
				}
			} else {
				// 前がCaps→Caps
				// 普通
				minCaps = x + beforeArray[1]
				// CapsLock → shift → CapsLock
				tempCapstime1 := z + y + z + beforeArray[1]
				if tempCapstime1 < minCaps {
					minCaps = tempCapstime1
				}

				// 前がCaps→noCaps
				// 普通 → CapsLock
				minNonCaps = z + x + beforeArray[1]
				// CapsLock → shift
				tempNoCapstime1 := z + y + beforeArray[1]
				if tempNoCapstime1 < minNonCaps {
					minNonCaps = tempNoCapstime1
				}

				// 前がnoCaps→Caps
				// CapsLock → 普通
				tempCapstime2 := z + x + beforeArray[0]
				if tempCapstime2 < minCaps {
					minCaps = tempCapstime2
				}
				// shift ⇨ CapsLock
				tempCapstime3 := y + z + beforeArray[0]
				if tempCapstime3 < minCaps {
					minCaps = tempCapstime3
				}

				// 前がnoCaps→noCaps
				// shift
				tempNoCapstime2 := y + beforeArray[0]
				if tempNoCapstime2 < minNonCaps {
					minNonCaps = tempNoCapstime2
				}
				// CapsLock → 普通 → CapsLock
				tempNoCapstime3 := z + x + z + beforeArray[0]
				if tempNoCapstime3 < minNonCaps {
					minNonCaps = tempNoCapstime3
				}

			}

			sArray[0] = minNonCaps
			sArray[1] = minCaps
			sArrayArray[i] = sArray
		}

		if i == sLen-1 {
			resultArray := sArrayArray[i]
			if resultArray[0] < resultArray[1] {
				result = resultArray[0]
			} else {
				result = resultArray[1]
			}
		}

	}

	fmt.Println(result)
}

func DShiftvsCapsLockReadLine(rdr *bufio.Reader) string {
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
