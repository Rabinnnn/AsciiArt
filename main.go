package main

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(word string) string {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	fileLines := strings.Split(string(file), "\n")
	var result strings.Builder

	for i := 0; i < 8; i++ {
		for _, char := range word {
			index := int(char - ' ')
			if index >= 0 && index < len(fileLines) {
				result.WriteString(fileLines[i+(index*9)+1])
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("The number of arguments is not correct")
		return
	}

	word := os.Args[1]
	fmt.Println(AsciiArt(word))
}
