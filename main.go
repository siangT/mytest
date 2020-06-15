package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, "2 *")
			fmt.Print("step5")

		}
		fmt.Println()
	}
	fmt.Println("qq")
}
