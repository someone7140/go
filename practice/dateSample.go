package main

import (
	"fmt"
	"time"
)

func main() {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	// JST今日日付
	currentTimeJst, _ := time.ParseInLocation("20060102", time.Now().In(jst).Format("20060102"), jst)
	fmt.Println(currentTimeJst)
	// JST昨日日付
	yesterDayTimeJst := currentTimeJst.AddDate(0, 0, -1)
	fmt.Println(yesterDayTimeJst)
	// UTC時間での日本の今日日付
	currentTimeUtc := currentTimeJst.UTC()
	fmt.Println(currentTimeUtc)
	// UTC時間での日本の昨日日付
	yesterDayTimeUtc := yesterDayTimeJst.UTC()
	fmt.Println(yesterDayTimeUtc)
}
