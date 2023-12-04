package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"log"
	"strings"
)

func main() {
	games := populateData()

	//Part One
	//testGame := TestGame{
	//	Red:   12,
	//	Blue:  14,
	//	Green: 13,
	//}

	//sum := 0
	//for _, g := range games {
	//	if isTestGameViable(g, testGame) {
	//		sum += g.Id
	//	}
	//}

	//Part Two
	sum := 0
	for _, g := range games {
		g.FindFewest()

		fmt.Printf("%+v\n", g)

		sum += g.MultiplyFewest()
	}

	fmt.Println(sum)

}

func isTestGameViable(g Game, t TestGame) bool {

	for _, s := range g.Sets {
		if s.Red > t.Red {
			return false
		}
		if s.Blue > t.Blue {
			return false
		}
		if s.Green > t.Green {
			return false
		}
	}

	return true
}

type Set struct {
	Red, Blue, Green int
}

type Game struct {
	Id     int
	Sets   []Set
	Fewest Set
}

func (g *Game) FindFewest() {
	for _, s := range g.Sets {
		if g.Fewest.Red < s.Red {
			g.Fewest.Red = s.Red
		}
		if g.Fewest.Blue < s.Blue {
			g.Fewest.Blue = s.Blue
		}
		if g.Fewest.Green < s.Green {
			g.Fewest.Green = s.Green
		}
	}
}

func (g *Game) MultiplyFewest() int {
	return g.Fewest.Red * g.Fewest.Blue * g.Fewest.Green
}

type TestGame struct {
	Red, Blue, Green int
}

func populateData() []Game {
	fileLines, err := utils.ReadFile("day-02/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	//iterate and print over fileLines
	var games []Game
	for _, line := range fileLines {
		fmt.Println(line)

		var game Game

		//getting the id
		_, err = fmt.Sscanf(line, "Game %d:", &game.Id)
		if err != nil {
			log.Fatal()
		}

		//removing the "Game <N>:" part
		line = strings.TrimPrefix(line, fmt.Sprintf("Game %d:", game.Id))

		//fmt.Println(line)

		//split ;
		sets := strings.Split(line, ";")
		for _, set := range sets {
			//fmt.Println(set)

			cubes := strings.Split(set, ",")
			var gameSet Set
			for _, cube := range cubes {
				//fmt.Println(cube)

				if strings.Contains(cube, "red") {
					fmt.Sscanf(cube, "%d red:", &gameSet.Red)
				}
				if strings.Contains(cube, "blue") {
					fmt.Sscanf(cube, "%d blue:", &gameSet.Blue)
				}
				if strings.Contains(cube, "green") {
					fmt.Sscanf(cube, "%d green:", &gameSet.Green)
				}
			}
			game.Sets = append(game.Sets, gameSet)
		}

		fmt.Printf("%+v\n", game)

		games = append(games, game)
	}

	return games
}
