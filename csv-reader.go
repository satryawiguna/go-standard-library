package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	var csvString string = "satrya, wiguna, gianyar\n" +
		"joni, artha, denpasar\n" +
		"ariston, dharmayudha, denpasar"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
