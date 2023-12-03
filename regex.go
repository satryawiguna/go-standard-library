package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regexOne *regexp.Regexp = regexp.MustCompile(`^s.*a`)

	fmt.Println(regexOne.MatchString("satrya"))
	fmt.Println(regexOne.FindAllString("Tai anjing satrya wiguna", 5))
}
