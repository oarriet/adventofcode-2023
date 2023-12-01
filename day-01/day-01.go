package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//read contents from file, copied from stackoverflow
	file, err := os.Open("day-01/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int

	for scanner.Scan() {
		fmt.Printf("line:[%s], pair:[%d]\n", scanner.Text(), getNumberPair(scanner.Text()))
		total += getNumberPair(scanner.Text())
	}

	fmt.Printf("total:[%d]\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var lettersToNumbersMap = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "4",
	"five":  "5e",
	"six":   "6",
	"seven": "7n",
	"eight": "e8t",
	"nine":  "9e",
}

func getNumberPair(s string) int {

	firstI, firstNum := findFirstNumberInLetters(s)
	if firstI >= 0 {
		s = replaceAt(firstI, s, firstNum, lettersToNumbersMap[firstNum])
	}

	lastI, lastNum := findLastNumberInLetters(s)
	if lastI >= 0 {
		s = replaceAt(lastI, s, lastNum, lettersToNumbersMap[lastNum])
	}

	var result uint8

	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			result = s[i] - 48
			break
		}
	}

	result = result * 10

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 48 && s[i] <= 57 {
			result = result + (s[i] - 48)
			break
		}
	}

	return int(result)
}

func findFirstNumberInLetters(s string) (index int, numberInLetters string) {
	index = -1
	numberInLetters = ""

	for k, _ := range lettersToNumbersMap {
		i := strings.Index(s, k)
		if i == -1 {
			continue
		}
		if index == -1 { //first time we found one
			index = i
			numberInLetters = k
		} else {
			if i < index {
				index = i
				numberInLetters = k
			}
		}
	}

	return index, numberInLetters
}

func findLastNumberInLetters(s string) (index int, numberInLetters string) {
	index = -1
	numberInLetters = ""

	for k, _ := range lettersToNumbersMap {
		i := strings.LastIndex(s, k)
		if i == -1 {
			continue
		}
		if index == -1 { //first time we found one
			index = i
			numberInLetters = k
		} else {
			if i > index {
				index = i
				numberInLetters = k
			}
		}
	}

	return index, numberInLetters
}

func replaceAt(i int, s string, old string, new string) string {
	if i == -1 {
		return s
	}
	return s[:i] + new + s[i+len(old):]
}
