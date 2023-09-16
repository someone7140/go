package main

import (
	"bufio"
	"fmt"
	"os"
)

func ALegendaryPlayersMain() {

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := ALegendaryPlayersReadLine(rdr)

	scoreMap := map[string]string{}
	scoreMap["tourist"] = "3858"
	scoreMap["ksun48"] = "3679"
	scoreMap["Benq"] = "3658"
	scoreMap["Um_nik"] = "3648"
	scoreMap["apiad"] = "3638"
	scoreMap["Stonefeang"] = "3630"
	scoreMap["ecnerwala"] = "3613"
	scoreMap["mnbvmar"] = "3555"
	scoreMap["newbiedmy"] = "3516"
	scoreMap["semiexp"] = "3481"

	fmt.Println(scoreMap[s])
}

func ALegendaryPlayersReadLine(rdr *bufio.Reader) string {
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
