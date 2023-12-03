package main

import (
	"fmt"
	"strconv"
)

func main() {
	var boolOne bool
	var err error

	boolOne, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolOne)
	} else {
		fmt.Println("Error", err.Error())
	}

	var floatOne float64
	if err == nil {
		fmt.Println(floatOne)
	} else {
		fmt.Println("Error", err.Error())
	}

	floatOne, err = strconv.ParseFloat("9.350", 64)

	binary := strconv.FormatInt(199, 2)
	fmt.Println("Binary: ", binary)

	fmt.Println(strconv.ParseInt("23", 8, 8))
	fmt.Println(strconv.Atoi("1.24"))
	fmt.Println(strconv.Itoa(23))
	fmt.Println(strconv.FormatBool(false))
}
