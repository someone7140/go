package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var k int
	fmt.Scan(&k)

	var result int64
	count := 0
	ketaKosuuMap := map[string]int{}
	for {
		if count <= 10 {
			result = int64(count)
		} else {
			resultStr := strconv.FormatInt(int64(result), 10)
			keta := len(resultStr)
			if keta <= 2 {
				tempResultStr := strconv.FormatInt(int64(result+1), 10)
				tempResultStrArray := strings.Split(tempResultStr, "")
				futaketaMe, _ := strconv.Atoi(tempResultStrArray[0])
				hitoketaMe, _ := strconv.Atoi(tempResultStrArray[1])
				if hitoketaMe >= futaketaMe {
					ketaKosuuMap["2-"+tempResultStrArray[0]] = futaketaMe
					if futaketaMe != 9 {
						result = 210
					} else {
						nextFutaketa := strconv.FormatInt(int64(futaketaMe+1), 10)
						tempResult2, _ := strconv.Atoi(nextFutaketa + "0")
						result = int64(tempResult2)
					}
				} else {
					result = result + 1
				}
			}
		}
		count = count + 1
		if count == k {
			break
		}
	}
	fmt.Println(result)
}
