package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	s "strings"
)

func main() {
	args := os.Args
	var fileContent string

	fileContent = getFileContent(args[1])
	if args[2] == "-w" {
		fmt.Println(countWords(fileContent))
	} else if args[2] == "-c" {
		fmt.Println(countCharacters(fileContent))
	} else if args[2] == "-l" {
		fmt.Println(countLines(fileContent))
	} else if args[2] == "-ml" {
		fmt.Println(largestText(fileContent))
		fmt.Println(largestLineLength(fileContent))
	} else if args[2] == "--help" {
		fmt.Println("-w \t get number of words in the file")
		fmt.Println("-c \t get number of characters in the file")
		fmt.Println("-l \t get number of lines in the file")
		fmt.Println("-ml \t get the Laragest Line in the file")
	} else {
		fmt.Println("unknown Command")
	}

}

func getFileContent(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		return string(file)
	}
	return ""
}

func countCharacters(content string) int {
	var lines []string = s.Split(content, "\n")
	var c int = 0
	for i := 0; i < len(lines); i++ {
		c += len(lines[i])
	}
	return c - (len(lines) - 1)
}

// Get Line Count
func countLines(content string) int {
	return len(s.Split(content, "\n"))
}

// Get Word Count
func countWords(value string) int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return len(results)
}

// Get largest Line Text
func largestText(content string) string {
	lines := s.Split(content, "\n")
	maximumCharCount := len(s.Split(lines[0], ""))
	var maxLine string = lines[0]

	i := 1
	for i < len(lines) {
		if maximumCharCount < len(s.Split(lines[i], "")) {
			maxLine = lines[i]
		}
		i++
	}
	return maxLine
}

// Get largest Line Text
func largestLineLength(content string) int {
	return countCharacters(largestText(content))
}
