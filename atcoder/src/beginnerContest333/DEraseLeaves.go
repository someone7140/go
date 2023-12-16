package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func DEraseLeavesMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	u := DEraseLeavesNewUnionFind(n + 1)
	var oneConnectList []int
	for i := 0; i < n-1; i++ {
		vStrs := strings.Split(DEraseLeavesReadLine(rdr), " ")
		v1, _ := strconv.Atoi(vStrs[0])
		v2, _ := strconv.Atoi(vStrs[1])

		if v1 == 1 || v2 == 1 {
			if v1 == 1 {
				oneConnectList = append(oneConnectList, v2)
			} else {
				oneConnectList = append(oneConnectList, v1)
			}
		} else {
			u.DEraseLeavesUnite(v1, v2)
		}

	}

	if len(oneConnectList) == 1 {
		fmt.Println(1)
	} else {
		var results []int
		for _, k := range oneConnectList {
			tempResult := u.size(k)
			results = append(results, tempResult)
		}
		sort.Ints(results)
		result := 0
		lenResults := len(results)
		for i := 0; i < lenResults-1; i++ {
			result = result + results[i]
		}

		fmt.Println(result + 1)
	}
}

func DEraseLeavesReadLine(rdr *bufio.Reader) string {
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

type DEraseLeavesUnionFind struct {
	par []int
}

/* コンストラクタ */
func DEraseLeavesNewUnionFind(N int) *DEraseLeavesUnionFind {
	u := new(DEraseLeavesUnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

/* xの所属するグループを返す */
func (u DEraseLeavesUnionFind) DEraseLeavesRoot(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.DEraseLeavesRoot(u.par[x])
	return u.par[x]
}

/* xの所属するグループ と yの所属するグループ を合体する */
func (u DEraseLeavesUnionFind) DEraseLeavesUnite(x, y int) {
	xr := u.DEraseLeavesRoot(x)
	yr := u.DEraseLeavesRoot(y)
	if xr == yr {
		return
	}
	u.par[yr] += u.par[xr]
	u.par[xr] = yr
}

/* xとyが同じグループに所属するかどうかを返す */
func (u DEraseLeavesUnionFind) DEraseLeavesSame(x, y int) bool {
	if u.DEraseLeavesRoot(x) == u.DEraseLeavesRoot(y) {
		return true
	}
	return false
}

/* xの所属するグループの木の大きさを返す */
func (u DEraseLeavesUnionFind) size(x int) int {
	return -u.par[u.DEraseLeavesRoot(x)]
}
