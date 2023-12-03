package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2023, time.December, 20, 7, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2023-12-20T15:04:00Z")
	fmt.Println(parse.Local())

	time := now.Local()
	fmt.Println("Day:", time.Day())
	fmt.Println("Month:", time.Month())

}
