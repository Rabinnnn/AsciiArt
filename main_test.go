package main

import (
	"io/ioutil"
	"testing"
	"AsciiArt/functions"
	
)

func TestAsciiArt(t *testing.T) {
	//create a map containing the input text as key and the path to its expected output as value
	testCases := map[string]string{
		"hello": "testCases/expectedOutput1.txt",
		"Hello": "testCases/expectedOutput2.txt",
		"Hello\\nThere": "testCases/expectedOutput3.txt",
	}

	for input, expectedFilePath := range testCases {
		//subtest to ensure that the right expected output filepath is used
		t.Run(input, func(t *testing.T) {
			expectedContent, err := ioutil.ReadFile(expectedFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", expectedFilePath, input, err)
			}

			//convert content read from the file to string
			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(input)
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", input, expectedContentStr, result)
			}
		})
	}


}
