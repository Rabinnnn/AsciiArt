package functions

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

	if input == "\\n"{
		fmt.Println()
		os.Exit(0)
	} else if input == ""{
		os.Exit(0)
	}  

	//split the contents of file by line
	fileLines := strings.Split(string(file), "\n")
	var result strings.Builder
	var words []string
	
	words = strings.Split(input, "\\n")

	for _, word := range words{
		if word != ""{			 
			for i := 0; i < 8; i++ {
				for _, char := range word {
					index := int(char - ' ') //find the difference between the ascii value of the letter and that of space
					if index >= 0 && index < len(fileLines){
						result.WriteString(fileLines[i+(index*9)+1])
					}
				}
				result.WriteString("\n")	
			}
		}else {
			result.WriteString("\n")
		}	
	}
	return result.String() //return the ASCII art as a string	
}