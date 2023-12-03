package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Nama saya adalah Satrya Wiguna", "Satrya"))
	fmt.Println(strings.Split("Satrya Wiguna", ","))
	fmt.Println(strings.ToLower("SATRYA WIGUNA"))
	fmt.Println(strings.ToUpper("satrya wiguna"))
	fmt.Println(strings.Trim("          satrya wiguna", " "))
	fmt.Println(strings.ReplaceAll("Satrya Wiguna", "Satrya", "Raden"))

}
