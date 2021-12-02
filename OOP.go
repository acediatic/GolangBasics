package main

import (
	"fmt"
	"reflect"
)

// uses composition rather than inheritance
// uses 'has-a' rater than 'is-a'

type Animal struct {
	Name string
	Age  int
}

// bird contains an aniaml
// however, bird is not an animal, and the type (Bird) will reflect that.
// its syntactic sugar to deference the aniaml fields.
type Bird struct {
	Animal   // this says embed the animal struct into the bird struct (not a named field)
	SpeedKPH float32
}

// tags are used to specify the fields of a struct
type taggedAnimal struct {
	Name   string `json:"name"`
	Origin string `required:"t  rue"`
}

func oop() {
	t := reflect.TypeOf(taggedAnimal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
