package datatypes

import (
	"fmt"
)

func intToFloat(a int) {
	fmt.Println(a, "from int to float is", float32(a))
}

func floatToInt(a float32) {
	fmt.Println(a, "from float to int is", int(a))
}

func stringToByteArray(str string) {
	fmt.Println(str, "from string to []byte", []byte(str))
}

func byteArrayToSting(b []byte) {
	fmt.Println(b, "from []byte to string", string(b))
}

func interfaceToInt(a interface{}) {
	aa, ok := a.(int)
	if !ok {
		fmt.Println(a, "cannot be parsed to int")
		return
	}
	fmt.Println(a, "from interface{} to int", aa)
}

func checkDatatype(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println(a, "-> int")
	case float32:
		fmt.Println(a, "-> float")
	case string:
		fmt.Println(a, "-> string")
	default:
		fmt.Println(a, "-> Unknown datatype")
	}
}

func Typecasting() {
	intToFloat(4)
	floatToInt(5.34)
	stringToByteArray("Hello")
	byteArrayToSting([]byte{72, 101, 108, 108, 111})
	interfaceToInt("asd")
	interfaceToInt(321)
	checkDatatype(23)
	checkDatatype("asdf")
}
