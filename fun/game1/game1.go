package game1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func inputNumber(inputText string) (int, error) {
	fmt.Print(inputText)
	input, _ := reader.ReadString('\n')
	formattedInput := strings.Replace(input, "\n", "", -1)
	number, err := strconv.ParseInt(formattedInput, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(number), nil
}

func welcome() {
	fmt.Println("Let's play a game. Input a number")
	fmt.Println("1) To get factorial of the number")
	fmt.Println("2) Print the table of the number")
	fmt.Println("3) To get summation of all numbers from 1 to the number")
	fmt.Println("4) To get summation of some comma separated numbers")
}

func printFactorial() {
	number, err1 := inputNumber("Input number: ")
	if err1 != nil {
		fmt.Println("invalid input")
	}

	ans := 1
	for i := 1; i <= number; i++ {
		ans *= i
	}
	fmt.Println("Factorial is ", ans)
}

func printTable() {
	number, err1 := inputNumber("Input number: ")
	if err1 != nil {
		fmt.Println("invalid input")
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}

func printSeriesSummation() {
	number, err1 := inputNumber("Input number: ")
	if err1 != nil {
		fmt.Println("invalid input")
	}

	ans := 0
	for i := 1; i <= number; i++ {
		ans += i
	}
	fmt.Println("summation is: ", ans)
}

func printListSummation() {
	fmt.Println("Give us space separated list of numbers")
	inputString, _ := reader.ReadString('\n')
	formattedString := strings.Replace(inputString, "\n", "", -1)
	numbers := strings.Split(formattedString, " ")

	var sum float32 = 0
	for _, value := range numbers {
		n, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("%v isn't a valid number\n", n)
			continue
		}
		sum += float32(n)
	}
	fmt.Println("summation is: ", sum)
}

func Play() {
	welcome()
	option, err := inputNumber("Choose option: ")
	if err != nil {
		fmt.Println("invalid input")
	}

	if option == 1 {
		printFactorial()
	} else if option == 2 {
		printTable()
	} else if option == 3 {
		printSeriesSummation()
	} else {
		printListSummation()
	}
}
