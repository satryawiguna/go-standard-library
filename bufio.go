package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var inputOne *strings.Reader = strings.NewReader("Ki Djoko Bodo saat ini")
	// readerOne := strings.NewReader("Ki Djoko Bodo saat ini")

	readerOne := bufio.NewReader(inputOne)

	for {
		readOneLine, _, err := readerOne.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(readOneLine))
	}

	writeOne := bufio.NewWriter(os.Stdout)

	writeOne.WriteString("Satrya Wiguna")
	writeOne.WriteString("Keren and Cool")
	writeOne.Flush()
}
