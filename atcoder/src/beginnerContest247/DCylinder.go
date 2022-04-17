package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ball struct {
	bangou int64
	kosuu  int64
}

func DCylinderMain() {
	var q int
	fmt.Scan(&q)

	var resultSlice []string
	var ballSlice []Ball

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < q; i++ {
		queryStr := DCylinderReadLine(rdr)
		queryStrArray := strings.Split(queryStr, " ")
		if len(queryStrArray) == 3 {
			kosuu, _ := strconv.Atoi(queryStrArray[2])
			addBangou, _ := strconv.Atoi(queryStrArray[1])
			ballSlice = append(ballSlice, Ball{
				kosuu: int64(kosuu), bangou: int64(addBangou),
			})
		} else {
			var goukei int64
			goukei = 0
			removeKosuuTemp, _ := strconv.Atoi(queryStrArray[1])
			removeKosuu := int64(removeKosuuTemp)
			for {
				if removeKosuu <= 0 {
					break
				}
				ball := ballSlice[0]
				if ball.kosuu == removeKosuu {
					goukei = goukei + (ball.kosuu * ball.bangou)
					ballSlice = removeBall(ballSlice, 0)
					removeKosuu = 0
				} else if ball.kosuu < removeKosuu {
					removeKosuu = removeKosuu - ball.kosuu
					goukei = goukei + (ball.kosuu * ball.bangou)
					ballSlice = removeBall(ballSlice, 0)
				} else {
					goukei = goukei + (removeKosuu * ball.bangou)
					ballSlice[0] = Ball{
						kosuu: ball.kosuu - removeKosuu, bangou: ball.bangou,
					}
					removeKosuu = 0
				}
			}
			resultSlice = append(resultSlice, strconv.FormatInt(goukei, 10))
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func removeBall(slice []Ball, s int) []Ball {
	return slice[1:]
}

func DCylinderReadLine(rdr *bufio.Reader) string {
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
