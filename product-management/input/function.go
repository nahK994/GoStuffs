package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Get(promtText string) string {
	fmt.Print(promtText)
	reader := bufio.NewReader(os.Stdin)
	inputText, _ := reader.ReadString('\n')
	formattedString := strings.Replace(inputText, "\n", "", -1)
	return formattedString
}
