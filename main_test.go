package main

import (
	"io/ioutil"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	testCases := map[string]string{
		"hello": "testCases/expectedOutput1.txt",
		"Hello": "testCases/expectedOutput2.txt",
		"Hello\n\nThere": "testCases/expectedOutput3.txt",
		// Add more test cases as needed
	}

	for input, expectedFilePath := range testCases {
		t.Run(input, func(t *testing.T) {
			expectedContent, err := ioutil.ReadFile(expectedFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", expectedFilePath, input, err)
			}

			// Remove trailing newline character from expected content
			expectedContentStr := string(expectedContent)

			result := AsciiArt(input)
			if result != expectedContentStr {
				t.Errorf("For input '%s', expected:\n%s\n but got:\n%s", input, expectedContentStr, result)
			}
		})
	}


}
