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
			totalPointsCaseOne += 1
			totalPointsCaseTwo += 1
		case outcome == "C Y":
			totalPointsCaseOne += 2
			totalPointsCaseTwo += 6
		case outcome == "A Z":
			totalPointsCaseOne += 3
			totalPointsCaseTwo += 8
		case outcome == "A X":
			totalPointsCaseOne += 4
			totalPointsCaseTwo += 3
		case outcome == "B Y":
			totalPointsCaseOne += 5
			totalPointsCaseTwo += 5
		case outcome == "C Z":
			totalPointsCaseOne += 6
			totalPointsCaseTwo += 7
		case outcome == "C X":
			totalPointsCaseOne += 7
			totalPointsCaseTwo += 2
		case outcome == "A Y":
			totalPointsCaseOne += 8
			totalPointsCaseTwo += 4
		case outcome == "B Z":
			totalPointsCaseOne += 9
			totalPointsCaseTwo += 9
		}
		
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	fmt.Printf("Total points is %v", totalPointsCaseOne)
	fmt.Printf("Total points is %v", totalPointsCaseTwo)
}
