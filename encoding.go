package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var name string = "My name is Satrya Wiguna"

	var encodeOne string = base64.StdEncoding.EncodeToString([]byte(name))

	fmt.Println(encodeOne)

	decodeOne, err := base64.StdEncoding.DecodeString(encodeOne)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decodeOne))
	}
}
