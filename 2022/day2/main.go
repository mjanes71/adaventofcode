package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	file, err := os.Open("/Users/meganracheljanes/Documents/sandbox/advent-of-code/2022/day2/day2_data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	totalPointsCaseOne := 0
	totalPointsCaseTwo := 0

    for scanner.Scan() {
		outcome := strings.TrimSpace(scanner.Text())
		
		switch {
		case outcome == "B X":
			totalPointsCaseOne = totalPointsCaseOne + 1
			totalPointsCaseTwo = totalPointsCaseTwo + 1
		case outcome == "C Y":
			totalPointsCaseOne = totalPointsCaseOne + 2
			totalPointsCaseTwo = totalPointsCaseTwo + 6
		case outcome == "A Z":
			totalPointsCaseOne = totalPointsCaseOne + 3
			totalPointsCaseTwo = totalPointsCaseTwo + 8
		case outcome == "A X":
			totalPointsCaseOne = totalPointsCaseOne + 4
			totalPointsCaseTwo = totalPointsCaseTwo + 3
		case outcome == "B Y":
			totalPointsCaseOne = totalPointsCaseOne + 5
			totalPointsCaseTwo = totalPointsCaseTwo + 5
		case outcome == "C Z":
			totalPointsCaseOne = totalPointsCaseOne + 6
			totalPointsCaseTwo = totalPointsCaseTwo + 7
		case outcome == "C X":
			totalPointsCaseOne = totalPointsCaseOne + 7
			totalPointsCaseTwo = totalPointsCaseTwo + 2
		case outcome == "A Y":
			totalPointsCaseOne = totalPointsCaseOne + 8
			totalPointsCaseTwo = totalPointsCaseTwo + 4
		case outcome == "B Z":
			totalPointsCaseOne = totalPointsCaseOne + 9
			totalPointsCaseTwo = totalPointsCaseTwo + 9
		}
		
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	fmt.Printf("Total points is %v", totalPointsCaseOne)
	fmt.Printf("Total points is %v", totalPointsCaseTwo)
}
