package utility

import (
	"io/ioutil"
	"log"
	"strings"
)

//Given a path and given the path leads to a txt-file this will return the content as a string.
func InputToString(path string) string  {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	return strings.TrimSuffix(string(content), "\n")
}

//Day 10
func Input2019Day1() string {
	return InputToString("2019/assets/day1.txt")
}

//Day 6
func Input2020Day6() string {
	return InputToString("2020/assets/day6.txt")
}

//Day 7
func Input2020Day7() string {
	return InputToString("2020/assets/day7.txt")
}

//Day 8
func Input2020Day8() string {
	return InputToString("2020/assets/day8.txt")
}

//Day 9
func Input2020Day9() string {
	return InputToString("2020/assets/day9.txt")
}

//Day 10
func Input2020Day10() string {
	return InputToString("2020/assets/day10.txt")
}

//Day 11
func Input2020Day11() string {
	return InputToString("2020/assets/day11.txt")
}

//Day 12
func Input2020Day12() string {
	return InputToString("2020/assets/day12.txt")
}

//Day 13
func Input2020Day13() string {
	return InputToString("2020/assets/day13.txt")
}

//Day 14
func Input2020Day14() string {
	return InputToString("2020/assets/day14.txt")
}

//Day 15
func Input2020Day15() string {
	return InputToString("2020/assets/day15.txt")
}

//Day 16
func Input2020Day16() string {
	return InputToString("2020/assets/day15.txt")
}

//Day 17
func Input2020Day17() string {
	return InputToString("2020/assets/day17.txt")
}

//Day 21
func Input2020Day21() string {
	return InputToString("2020/assets/day21.txt")
}

//Day 22
func Input2020Day22() string {
	return InputToString("2020/assets/day22.txt")
}
