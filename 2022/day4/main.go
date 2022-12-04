package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	//read a line from the file
	file, err := os.Open("day4_data.txt")
    if err != nil {
        log.Fatalf("Try opening this file with the secret password next time.")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

	overlapCount := 0
	partialOverlapCount := 0

	for scanner.Scan() {
		//replace hyphen with commma so string is just csv
		line := strings.TrimSpace(scanner.Text())
		line = strings.Replace(line,"-",",",-1)
		//create slice so we can grab values from indexes
		mySlice := strings.Split(line,",")
		//high and low values
		aLow,_ := strconv.Atoi(mySlice[0])
		aHigh,_ := strconv.Atoi(mySlice[1])
		bLow,_ := strconv.Atoi(mySlice[2])
		bHigh,_ := strconv.Atoi(mySlice[3])

		//find if the ranges indicate an overlap
		isOverlap := compareForCompleteOverlap(aLow, aHigh, bLow, bHigh)
		isPartialOverlap := compareForPartialOverlap(aLow, aHigh, bLow, bHigh)

		if isOverlap {
			overlapCount ++
		}else if isPartialOverlap {
			partialOverlapCount ++
		}

	}
	fmt.Printf("\nThe overlap count is %v", overlapCount)
	fmt.Printf("\nThe partial overlap count is %v", partialOverlapCount)
	fmt.Printf("\nThe total overlap count is %v", overlapCount + partialOverlapCount)

}

func compareForCompleteOverlap(aLow, aHigh, bLow, bHigh int) bool {
	theyOverlap := false

	switch {
	case aLow >= bLow && aHigh <= bHigh:
		theyOverlap = true
	case aLow <= bLow && aHigh >= bHigh:
		theyOverlap = true
	}

	return theyOverlap
}

func compareForPartialOverlap(aLow, aHigh, bLow, bHigh int) bool {
	theyOverlap := false

	switch {
	case aHigh == bLow:
		theyOverlap = true
	case aHigh > bLow:
		if aLow <= bLow {
			theyOverlap = true
		}else if bHigh >= aLow {
			theyOverlap = true
		}
	}

	return theyOverlap
}