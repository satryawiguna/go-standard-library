package main

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s UserSlice) Swap(i, j int) {
	// var name string = s[i].name

	// s[i].name = s[j].name
	// s[j].name = name

	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Satrya", 40},
		{"Wiguna", 39},
		{"Ganjar", 50},
		{"Pranowo", 49},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
