package main

import (
	"advent-of-code-golang/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var gameList games
	common.ReadFile(os.Args[1], func(s string) {
		gameList = append(gameList, parseGame(s))
	})
	gameList = filterGames(gameList, 12, 14, 13)
	checksum := reduceGames(gameList)
	fmt.Println(checksum)
}

func filterGames(gameList games, maxRed, maxBlue, maxGreen int) games {
	var filteredList games
	for _, g := range gameList {
		isValid := true
		for _, set := range g.sets {
			if set.green > maxGreen || set.blue > maxBlue || set.red > maxRed {
				isValid = false
				break
			}
		}
		if isValid {
			filteredList = append(filteredList, g)
		}
	}
	return filteredList
}

func reduceGames(gameList games) int {
	acc := 0
	for _, g := range gameList {
		acc += g.id
	}
	return acc
}

func parseGame(gameStr string) game {
	g := game{}
	gameId, sets := splitGameAndSets(gameStr)
	g.id = gameId
	g.sets = parseSets(sets)
	return g
}

func splitGameAndSets(game string) (gameId int, sets string) {
	parts := strings.Split(game, ": ")
	gameId, err := strconv.Atoi(strings.ReplaceAll(parts[0], "Game ", ""))
	if err != nil {
		panic(err.Error())
	}
	sets = parts[1]
	return
}

func parseSets(setsStr string) (s sets) {
	setParts := strings.Split(setsStr, "; ")
	for _, part := range setParts {
		s = append(s, parseSet(part))
	}
	return
}

func parseSet(setStr string) (s set) {
	cubes := strings.Split(setStr, ", ")
	for _, cube := range cubes {
		cubeParts := strings.Split(cube, " ")
		switch cubeParts[1] {
		case "red":
			s.red, _ = strconv.Atoi(cubeParts[0])
		case "blue":
			s.blue, _ = strconv.Atoi(cubeParts[0])
		case "green":
			s.green, _ = strconv.Atoi(cubeParts[0])
		}
	}
	return
}

type games []game
type sets []set

type game struct {
	id   int
	sets sets
}

type set struct {
	red   int
	green int
	blue  int
}
