package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxCubesPerColorMap = map[string]int{"red": 12, "green": 13, "blue": 14}

func main() {
	sumOfGameIds := 0
	f, err := os.ReadFile("input.txt")
	check(err)
	separatedRows := strings.Split(string(f), "\n")

	for _, row := range separatedRows {
		var impossibleGame bool
		cubeNumberMap := make(map[string]int)
		separatedInput := strings.FieldsFunc(row, split)
		for i := 1; i < len(separatedInput); i++ {
			cubeCounts := strings.Split(separatedInput[i], " ")
			number, err := strconv.Atoi(cubeCounts[1])
			check(err)
			if cubeNumberMap[cubeCounts[2]] < number || cubeNumberMap[cubeCounts[2]] == 0 {
				cubeNumberMap[cubeCounts[2]] = number
			}
		}

		for cube := range maxCubesPerColorMap {
			if cubeNumberMap[cube] > maxCubesPerColorMap[cube] {
				impossibleGame = true
				break
			}
		}

		if !impossibleGame {
			game := strings.Split(separatedInput[0], " ")
			gameId, err := strconv.Atoi(game[1])
			check(err)
			sumOfGameIds += gameId
		}
	}

	fmt.Println(sumOfGameIds)
}

func split(r rune) bool {
	return r == ';' || r == ',' || r == ':'
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
