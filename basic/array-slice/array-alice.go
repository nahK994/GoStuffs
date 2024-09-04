package arrayslice

import (
	"fmt"
	"reflect"
)

func Arrayslice() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	fmt.Println("Slices equal:", reflect.DeepEqual(slice1, slice2)) // true

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}

	fmt.Println("Maps equal:", reflect.DeepEqual(map1, map2)) // true

	// Different order in maps, still equal
	map3 := map[string]int{"b": 2, "a": 1}
	fmt.Println("Maps with different order equal:", reflect.DeepEqual(map1, map3)) // true

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
	arr2 := [5]int{1, 2, 3, 4}
	fmt.Println(arr1, arr2, arr1 == arr2)
}
