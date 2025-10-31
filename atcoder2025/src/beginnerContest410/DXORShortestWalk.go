package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DXORShortestWalkRoute struct {
	id    int
	start int
	end   int
	w     int
}

func DXORShortestWalkMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	routeStartMap := map[int][]DXORShortestWalkRoute{}
	for i := 0; i < m; i++ {
		abwStrList := strings.Split(DXORShortestWalkRdr(rdr), " ")
		start, _ := strconv.Atoi(abwStrList[0])
		end, _ := strconv.Atoi(abwStrList[1])
		w, _ := strconv.Atoi(abwStrList[2])

		routeMapList, ok := routeStartMap[start]
		if ok {
			routeStartMap[start] = append(routeMapList, DXORShortestWalkRoute{
				id:    i,
				start: start,
				end:   end,
				w:     w,
			})
		} else {
			routeStartMap[start] = []DXORShortestWalkRoute{DXORShortestWalkRoute{
				id:    i,
				start: start,
				end:   end,
				w:     w,
			}}
		}
	}

	result := -1

	var loopFunc func(targetStart int, tempResult int)
	visitedMapIdSet := map[int]struct{}{}
	loopFunc = func(targetStart int, tempResult int) {
		routeList, ok := routeStartMap[targetStart]
		if ok {
			for _, route := range routeList {
				_, ok2 := visitedMapIdSet[route.id]
				if !ok2 {
					visitedMapIdSet[route.id] = struct{}{}
					var tempResult2 int
					if tempResult == -1 {
						tempResult2 = route.w
					} else {
						tempResult2 = route.w ^ tempResult
					}

					if route.end == n && (result == -1 || tempResult2 < result) {
						result = tempResult2
					}

					loopFunc(route.end, tempResult2)
					delete(visitedMapIdSet, route.id)
				}
			}

		}
	}
	loopFunc(1, -1)
	fmt.Println(result)
}

func DXORShortestWalkRdr(rdr *bufio.Reader) string {
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
