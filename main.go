package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, "2 *")
			fmt.Print("1")
			fmt.Print(i, "44 *")
			fmt.Print("-")
			fmt.Print("new")

		}
		fmt.Println()
	}
	fmt.Println("qq")
}
