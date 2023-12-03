package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotfoundError   = errors.New("Notfound Error")
)

func main() {
	var err error = GetById("")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Ini error validation")
		} else if errors.Is(err, NotfoundError) {
			fmt.Println("Ini error notfound")
		} else {
			fmt.Println("Error tak terdeteksi")
		}
	}
}

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Satrya Wiguna" {
		return NotfoundError
	}

	return nil
}
