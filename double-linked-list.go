package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Satrya")
	data.PushFront("Wiguna")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	var dataDua *list.List = list.New()

	dataDua.PushBack("Adi Sustika")
	dataDua.PushBack("Satrya Wiguna")
	dataDua.PushBack("Erna Widhiastuti")

	var head *list.Element = dataDua.Front()

	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

}
