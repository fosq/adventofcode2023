package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sum := 0

	f, err := os.ReadFile("input.txt")
	check(err)

	separatedInput := strings.Split(string(f), "\n")
	expr := `(?i)(\d)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`
	revExpr := `(?i)(\d)|(eno)|(owt)|(eerht)|(ruof)|(evif)|(xis)|(neves)|(thgie)|(enin)`
	r, _ := regexp.Compile(expr)
	revR, _ := regexp.Compile(revExpr)

	for _, row := range separatedInput {
		reversedRow := reverseString(row)
		match := r.FindAllString(row, -1)
		matchReverse := revR.FindAllString(reversedRow, -1)
		nbrArr := convertStringNumbersToNumbers(match)
		reversedNbrArr := convertStringNumbersToNumbers(matchReverse)

		stringNumber := nbrArr[0] + reversedNbrArr[0]
		number, err := strconv.Atoi(stringNumber)
		check(err)

		fmt.Println(row)
		fmt.Println(nbrArr, reversedNbrArr)
		fmt.Println(number)

		sum += number
	}

	fmt.Println(sum)
}

func convertStringNumbersToNumbers(stringArr []string) []string {
	for i, str := range stringArr {
		switch str {
		case "one", "eno":
			stringArr[i] = "1"
		case "two", "owt":
			stringArr[i] = "2"
		case "three", "eerht":
			stringArr[i] = "3"
		case "four", "ruof":
			stringArr[i] = "4"
		case "five", "evif":
			stringArr[i] = "5"
		case "six", "xis":
			stringArr[i] = "6"
		case "seven", "neves":
			stringArr[i] = "7"
		case "eight", "thgie":
			stringArr[i] = "8"
		case "nine", "enin":
			stringArr[i] = "9"
		}
	}

	return stringArr
}

func reverseString(str string) string {
	var reversedString string
	for i := range str {
		reversedString += string(str[len(str)-1-i])
	}

	return reversedString
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
