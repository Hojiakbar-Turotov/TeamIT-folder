package main

import (
	"fmt"
	"os"
	cv "strconv"
)

func main() {
	/*	for n := 0; n <= 9; n++ {
		fmt.Println(n)
	}*/
	var x, _ = cv.Atoi(os.Args[1])
	var c = os.Args[2]
	var y, _ = cv.Atoi(os.Args[3])

	if c == "*" {
		fmt.Println(x * y)
	} else if c == "+" {
		fmt.Println(x + y)
	} else if c == "-" {
		fmt.Println(x - y)
	} else if c == "/" {
		fmt.Println(x / y)
	}
}
