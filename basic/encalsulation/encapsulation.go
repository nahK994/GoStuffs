package encalsulation

import (
	"basic/encalsulation/engine"
	"fmt"
)

type Car struct {
	engine.Engine // anonymous field (embedded struct)
	model         string
}

func Encalsulation() {
	car := Car{
		Engine: engine.Engine{},
		model:  "Mustang",
	}

	// Accessing the horsepower through a public method
	car.SetHorsepower(121)
	fmt.Println("model:", car.model, ", horsepower:", car.Horsepower()) // Works fine

	// Trying to access private field directly (this will cause a compile-time error)
	// fmt.Println("Horsepower:", car.horsepower) // Error: cannot refer to unexported field or method horsepower
}
