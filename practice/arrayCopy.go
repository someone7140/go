package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	sliceCopy := func(in, out interface{}) {
		buf := new(bytes.Buffer)
		gob.NewEncoder(buf).Encode(in)
		gob.NewDecoder(buf).Decode(out)
	}

	a := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{11, 12, 13}, {14, 15, 16}, {17, 18, 19}},
		{{21, 22, 23}, {24, 25, 26}, {27, 28, 29}},
	}

	var b [][][]int
	sliceCopy(a, &b)
	// bの値だけ変更する
	b[0][0][0] = 121

	fmt.Println(a)
	fmt.Println(b)
}
