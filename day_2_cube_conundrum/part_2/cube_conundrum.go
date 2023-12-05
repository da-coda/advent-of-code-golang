package main

import (
	"advent-of-code-golang/common"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var checksum int
	common.ReadFile(os.Args[1], func(s string) {
		g := parseGame(s)
		minDiceSet := findMinDiceSet(g)
		checksum += minDiceSet.power()
	})
	fmt.Println(checksum)
}

func findMinDiceSet(g game) set {
	return g.sets.min()
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

func (s sets) min() set {
	var r []int
	var g []int
	var b []int
	for _, set := range s {
		r = append(r, set.red)
		g = append(g, set.green)
		b = append(b, set.blue)
	}

	minSet := set{}
	minSet.red = slices.Max(r)
	minSet.green = slices.Max(g)
	minSet.blue = slices.Max(b)
	return minSet
}

func (s set) power() int {
	mul := 1
	for _, i := range []int{s.red, s.blue, s.green} {
		mul *= i
	}
	return mul
}
