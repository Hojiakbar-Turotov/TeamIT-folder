package main

import (
	"fmt"
	// "os"
)

func main() {
	/*	fmt.Println("Hello Rustam")
		fmt.Println("Hello Temur")
		fmt.Println("Hello Shohruh")
		fmt.Println("Hello Hojiakbar")
		fmt.Println("Hello Dilshoda")
		fmt.Println("Hello Zafar")
		fmt.Println("Hello Abduazim") */

	var users []string = []string{
		"Rustam", "Temur", "Shohruh", "Hojiakbar", "Dilshoda", "Zafar", "Abduazim",
	}

	for _, u := range users {
		// fmt.Println("Hello " + u + "!")
		if len(u) > 3 {
			fmt.Printf("Hello %s!\n", u)
		}
		fmt.Printf("keyingisi \n")
		if u == "Temur" {
			continue
		}
		fmt.Printf("Hello %s!\n", u)
		fmt.Printf("keyingisi \n")
		if u == "Temur" {
			break
		}
		fmt.Printf("Hello %s!\n", u)

	}

	/*age := 20
	if age >= 20 {
		fmt.Println(true)
	}
	*/
	/*	fmt.Println(20 > 10)
		fmt.Println(20 < 10)
		fmt.Println(20 == 10)
		fmt.Println(20 >= 10)
		fmt.Println("A" == "B")
		fmt.Println("A" <= "B")
	*/
	/*if os.Args[1] == "ism" || os.Args[1] == "Hojiakbar" {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("ism hato")
	}*/
}
