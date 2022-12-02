package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main(){
	file, err := os.Open("/Users/meganracheljanes/Documents/sandbox/advent-of-code/2022/day1/day1_data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    maxCal := 0
	secondMostCals := 0
	thirdMostCals := 0
	currentCal := 0
	var allThemCals []int

    for scanner.Scan() {
		if scanner.Text() == "" {
			allThemCals = append(allThemCals, currentCal)
			switch {
			case currentCal > maxCal:
				maxCal = currentCal
			case currentCal > secondMostCals:
				secondMostCals = currentCal
			case currentCal > thirdMostCals:
				thirdMostCals = currentCal
			}
			fmt.Println(currentCal)
			currentCal = 0
		}else {
			calOnLine, _ := strconv.Atoi(scanner.Text())
			currentCal = currentCal + calOnLine
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	sort.Sort(sort.Reverse(sort.IntSlice(allThemCals)))

	fmt.Printf("Maxcal is %v and second is %v and third is %v", allThemCals[0], allThemCals[1], allThemCals[2])
	fmt.Printf("Total cals is %v", allThemCals[0]+allThemCals[1]+allThemCals[2] )

}
