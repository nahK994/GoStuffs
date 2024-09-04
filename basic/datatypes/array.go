package datatypes

import (
	"fmt"
	"reflect"
)

func Array() {
	// Nested structure
	type Person struct {
		Name string
		Age  int
		Tags []string
	}

	person1 := Person{Name: "Alice", Age: 30, Tags: []string{"golang", "developer"}}
	person2 := Person{Name: "Alice", Age: 30, Tags: []string{"golang", "developer"}}

	fmt.Println("Structs equal:", reflect.DeepEqual(person1, person2)) // true

	// Modify one element in the slice within the struct
	person2.Tags[0] = "python"
	fmt.Println("Structs equal after modification:", reflect.DeepEqual(person1, person2)) // false

	arr1 := [5]int{1, 2, 3, 4}
	arr2 := [5]int{1, 2, 4, 3}
	fmt.Println(arr1, arr2, arr1 == arr2) // false
}
