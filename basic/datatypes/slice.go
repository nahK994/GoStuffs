package datatypes

import (
	"fmt"
	"reflect"
)

func Slice() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	fmt.Println("Slices equal:", reflect.DeepEqual(slice1, slice2)) // true
}
