package main

import (
	"bmi-calculator/calculator"
	"bmi-calculator/input"
	"bmi-calculator/introduction"
	"fmt"
)

func main() {
	introduction.Header()
	height := input.Input("Give us you height (m)")
	weight := input.Input("Give us your weight (kg)")

	fmt.Printf("%0.2f\n", calculator.Calculate(height, weight))
}
