package main

import (
	"encoding/csv"
	"os"
)

func main() {
	var writer *csv.Writer = csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Satrya", "Wiguna", "Gianyar"})
	_ = writer.Write([]string{"Erna", "Widhiastuti", "Klungkung"})
	_ = writer.Write([]string{"Arsha", "Kaela", "Germany"})
	_ = writer.Write([]string{"Riana", "Laskmidewi", "Norwegia"})

	writer.Flush()
}
