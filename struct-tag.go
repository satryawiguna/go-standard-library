package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"200"`
}

type Person struct {
	Name string `required:"true" max:"200"`
	Address string `required:"true" max:"100"`
	Email string `required:"true" max:"50"`
}

func ReadField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name", valueType)

	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)

		fmt.Println(valueField.Name)
		fmt.Println(valueField.Type)
	}
}

func IsValid(data any) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" {
			return reflect.
		}
	}
}

func main() {
	ReadField(Sample{"Satrya"})
	ReadField(Person{"Wiguna"})
}
