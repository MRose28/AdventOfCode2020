package utility

import (
	"io/ioutil"
	"log"
	"strings"
)

//Given a path and given the path leads to a txt-file this will return the content as a string.
func inputToString(path string) string  {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	return strings.TrimSuffix(string(content), "\n")
}

//Day 6
func InputDay6() string {
	return inputToString("assets/day6.txt")
}

//Day 7
func InputDay7() string {
	return inputToString("assets/day7.txt")
}

//Day 8
func InputDay8() string {
	return inputToString("assets/day8.txt")
}

//Day 9
func InputDay9() string {
	return inputToString("assets/day9.txt")
}

//Day 10
func InputDay10() string {
	return inputToString("assets/day10.txt")
}