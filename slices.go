package main

import (
	"fmt"
	"slices"
)

func main() {
	var dayName = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	monthName := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	number := []int{10, 20, 30, 40, 50}

	fmt.Println(slices.Max(number))
	fmt.Println(slices.Min(number))
	fmt.Println(slices.Contains(monthName, "July"))
	fmt.Println(slices.Index(dayName, "Wednesday"))
}
