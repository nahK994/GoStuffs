package main
import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("Hello World!!!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Number of bytes written:", n)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}