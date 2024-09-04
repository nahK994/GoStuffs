package datatypes

import "fmt"

func String() {
	str1 := "asdf"
	str2 := "asdf"
	fmt.Println(str1, str2, str1 == str2)

	str3 := "1232321"
	fmt.Println(str3)
	for i := 0; i < len(str3); i++ {
		if str3[i] != str3[len(str3)-1-i] {
			fmt.Println("Not palindrome")
			break
		}

		if i >= len(str3)-1-i {
			fmt.Println("Palindrome")
			break
		}
	}
}
