package datatypes

import (
	"fmt"
	"reflect"
)

func Map() {
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}

	fmt.Println("Maps equal:", reflect.DeepEqual(map1, map2)) // true

	// Different order in maps, still equal
	map3 := map[string]int{"b": 2, "a": 1}
	fmt.Println("Maps with different order equal:", reflect.DeepEqual(map1, map3)) // true
}
