package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BCountDistinctIntegersMain() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	aStr := scanner.Text()
	aArray := strings.Split(aStr, " ")

	set := make(map[string]struct{})
	for _, v := range aArray {
		set[v] = struct{}{}
	}

	fmt.Println(len(set))
}
