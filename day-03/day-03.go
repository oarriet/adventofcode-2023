package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"log"
)

func main() {
	fileLines, err := utils.ReadFile("day-03/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	matrix := buildMatrix(fileLines)
	//fmt.Println(matrix)

	partNumbers := make([]PartNumber, 0)

	for i, row := range matrix {

		var n int
		numberIsComplete := false
		isThereANumber := false

		start := -1
		end := -1

		for j, char := range row {
			//get the whole number

			if isNumber(char) {
				if start < 0 {
					start = j
				}
				n = appendNumber(n, getNumber(char))
				isThereANumber = true
			} else {
				if isThereANumber {
					numberIsComplete = true
					end = j - 1
				}

				if numberIsComplete {
					partNumbers = append(partNumbers, PartNumber{
						number: n,
						beginning: Coordinate{
							i: i,
							j: start,
						},
						end: Coordinate{
							i: i,
							j: end,
						},
					})
				}

				n = 0
				numberIsComplete = false
				isThereANumber = false

				start = -1
				end = -1
			}
		}

		//in case the number is the last thing in the row
		if isThereANumber {
			numberIsComplete = true
			end = len(row) - 1
		}
		if numberIsComplete {
			partNumbers = append(partNumbers, PartNumber{
				number: n,
				beginning: Coordinate{
					i: i,
					j: start,
				},
				end: Coordinate{
					i: i,
					j: end,
				},
			})
		}
	}

	gearMap := make(map[string][]int)
	for _, partNumber := range partNumbers {
		hasGear, i, j := checkSubMatrixForGear(matrix, partNumber.beginning, partNumber.end)
		if hasGear {
			gearMap[fmt.Sprintf("%d-%d", i, j)] = append(gearMap[fmt.Sprintf("%d-%d", i, j)], partNumber.number)
		}
	}

	sum := 0

	for k, adjNumbers := range gearMap {
		if len(adjNumbers) == 2 {
			fmt.Printf("adding i-j: %d for numbers: %d * %d\n", k, adjNumbers[0], adjNumbers[1])
			sum += adjNumbers[0] * adjNumbers[1]
		}
	}

	fmt.Println(sum)
}

type PartNumber struct {
	number int

	beginning Coordinate
	end       Coordinate
}

type Coordinate struct {
	i, j int
}

func checkSubMatrixForGear(m [][]rune, beginning, end Coordinate) (bool, int, int) {
	iBeginning := 0
	if beginning.i > 0 {
		iBeginning = beginning.i - 1
	}
	iEnd := end.i
	if end.i < len(m)-1 {
		iEnd = end.i + 1
	}
	jBeginning := 0
	if beginning.j > 0 {
		jBeginning = beginning.j - 1
	}
	jEnd := end.j
	if end.j < len(m)-1 {
		jEnd = end.j + 1
	}

	for i := iBeginning; i <= iEnd; i++ {
		for j := jBeginning; j <= jEnd; j++ {
			//fmt.Printf("i:%d, j:%d -> %d\n", i, j, m[i][j])
			if m[i][j] == 42 {
				return true, i, j
			}
		}
	}

	return false, 0, 0
}

// checkSubMatrix we will check the subMatrix including row before, left, right, row after
func checkSubMatrix(m [][]rune, beginning, end Coordinate) bool {
	iBeginning := 0
	if beginning.i > 0 {
		iBeginning = beginning.i - 1
	}
	iEnd := end.i
	if end.i < len(m)-1 {
		iEnd = end.i + 1
	}
	jBeginning := 0
	if beginning.j > 0 {
		jBeginning = beginning.j - 1
	}
	jEnd := end.j
	if end.j < len(m)-1 {
		jEnd = end.j + 1
	}

	for i := iBeginning; i <= iEnd; i++ {
		for j := jBeginning; j <= jEnd; j++ {
			//fmt.Printf("i:%d, j:%d -> %d\n", i, j, m[i][j])
			if m[i][j] < 46 || m[i][j] > 57 {
				return true
			}
			if m[i][j] == 47 {
				return true
			}
		}
	}

	return false
}

func buildMatrix(fileLines []string) [][]rune {

	matrix := make([][]rune, len(fileLines))
	for i, line := range fileLines {
		row := make([]rune, len(line))
		for j, char := range line {
			row[j] = char
		}
		matrix[i] = row
	}

	return matrix
}

func isNumber(r rune) bool {
	return r >= 48 && r <= 57
}

func getNumber(r rune) int {
	return int(r - 48)
}
func appendNumber(n int, a int) int {
	return n*10 + a
}
