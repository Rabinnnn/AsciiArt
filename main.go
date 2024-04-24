package main

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(input string) string {
	//Read the contents of standard.txt 
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	//split the contents of file by line
	fileLines := strings.Split(string(file), "\n")

	var result strings.Builder

	//split the input string by \n to handle cases where the argument has newline character
	words := strings.Split(input, "\\n")

	for _, word := range words{
		for i := 0; i < 8; i++ {
			for _, char := range word {
				index := int(char - ' ') //find the difference between the ascii value of the letter and that of space
				if index >= 0 && index < len(fileLines){
					result.WriteString(fileLines[i+(index*9)+1])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String() //return the ASCII art as a string
}

func main() {
	//Ensure that the number of arguments entered is 2
	if len(os.Args) != 2 {
		fmt.Println("The number of arguments is not correct")
		return
	}

	input := os.Args[1]

	fmt.Println(AsciiArt(input))
}
