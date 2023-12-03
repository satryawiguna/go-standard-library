package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("helloworld/hello.go"))
	fmt.Println(path.Ext("helloworld/hello.go"))
	fmt.Println(path.Join("helloworld", "satrya", "hello.go"))

}
