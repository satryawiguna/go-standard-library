package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(10)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)

		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
