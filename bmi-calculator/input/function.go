package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input(text string) float64 {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print(text + " ")
	inputText, _ := reader.ReadString('\n')
	prcessedInputText := strings.Replace(inputText, "\n", "", -1)
	inputNumber, _ := strconv.ParseFloat(prcessedInputText, 64)
	return inputNumber
}
