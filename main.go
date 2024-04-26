package main

import (
	"fmt"
	"os"
	"AsciiArt/functions"
)


func main() {
	//Ensure that the number of arguments entered is 2
	if len(os.Args) != 2 {
		fmt.Println("The number of arguments is not correct")
		return
	}

	input := os.Args[1]

	fmt.Print(functions.AsciiArt(input))
	
}
