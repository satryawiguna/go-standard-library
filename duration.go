package main

import (
	"fmt"
	"time"
)

func main() {
	var durationOne time.Duration = time.Second * 100
	durationTwo := time.Millisecond * 10

	durationThree := durationOne - durationTwo

	fmt.Println(durationThree)
}
