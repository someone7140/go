package main

import (
	"fmt"
	"reflect"
	"time"
)

type ReflectSampleStruct struct {
	IntSample    int
	StringSample string
	DateSample   time.Time
}

func ReflectSampleMain() {
	reflectSample := ReflectSampleStruct{
		IntSample:    100,
		StringSample: "sample",
		DateSample:   time.Now(),
	}
	reflectValue := reflect.ValueOf(reflectSample)
	reflectType := reflect.TypeOf(reflectSample)
	fieldLen := reflectType.NumField()
	for i := 0; i < fieldLen; i++ {
		field := reflectType.Field(i)
		value := reflectValue.FieldByName(field.Name)
		// 一度Interface型で取得する
		valueInterface := value.Interface()
		switch v := valueInterface.(type) {
		case int:
			fmt.Println("【int】")
			fmt.Println(v)
		case string:
			fmt.Println("【string】")
			fmt.Println(v)
		case time.Time:
			fmt.Println("【time.Time】")
			fmt.Printf("%d/%d/%d", v.Year(), v.Month(), v.Day())
		}
	}

}
