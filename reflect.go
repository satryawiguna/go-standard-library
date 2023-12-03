package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type Person struct {
	Name string
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name", valueType)

	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)

		fmt.Println(valueField.Name)
		fmt.Println(valueField.Type)
	}
}

func main() {
	readField(Sample{"Satrya"})
	readField(Person{"Wiguna"})
}
