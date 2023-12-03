package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("helloworld/hello.go"))
	fmt.Println(filepath.Base("helloworld/hello.go"))
	fmt.Println(filepath.Ext("helloworld/hello.go"))
	fmt.Println(filepath.IsAbs("helloworld/hello.go"))
	fmt.Println(filepath.IsLocal("helloworld/hello.go"))
	fmt.Println(filepath.Join("helloworld", "satrya", "hello.go"))
}
