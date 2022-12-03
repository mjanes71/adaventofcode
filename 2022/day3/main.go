package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const LOWERCASE_ABC =  "abcdefghijklmnopqrstuvwxyz"
const UPPERCASE_ABC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"


func main(){
	file, err := os.Open("day3_data.txt")
    if err != nil {
        log.Fatalf("We'll never tell you what's in the sack.")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

	//total of all duplicate items in the ruck
	ruckTotalPart1 := 0
	ruckTotalPart2 := 0
	groupCounter := 0
	var groupSack1 map[string]int
	var groupSack2 map[string]int

	//create maps to figure out item values
	ruckItemValues := makeItemValueMap()

    for scanner.Scan() {
		groupCounter += 1
		line := strings.TrimSpace(scanner.Text())
		
		if groupCounter == 3 {
			//logic for part 2
			//do stuff with the 3 sacks to figure out the overlap
			for _, item := range sortString(line) {
				_, itsIn1 := groupSack1[item]
				_, itsIn2 := groupSack2[item]
				if itsIn1 && itsIn2 {
					ruckTotalPart2 += ruckItemValues[item]
					break
				}
			}
			//add overlap value to part2 total
			//reset the group counter to 0
			groupCounter = 0
		}
		//to keep track of grouping for part 2
		switch {
		case groupCounter == 1:
			groupSack1 = makeAMap(line)
		case groupCounter == 2:
			groupSack2 = makeAMap(line)
		}

		//back to part 1 after that short commercial break

		ruck := strings.Split(line, "")
		ruckCenter := len(ruck)/2
		
		//sort the string
		sortedRuck := sortString(line)

		//split the string

		compartmentOnePt1 := ruck[:ruckCenter]
		compartmentTwoPt1 := ruck[ruckCenter:]

		duplicateItemPart1 := findDuplicateforTwo(compartmentOnePt1, compartmentTwoPt1, sortedRuck)

		//get value of duplicate item and add it to the ruckTotal

		ruckTotalPart1 += ruckItemValues[duplicateItemPart1] 	
	}
    err = scanner.Err()
	if err != nil {
        log.Fatalf("The scanner is broke. It's probably related to the printer.")
    }

	fmt.Printf("\nThe ruck total for part1 is %v", ruckTotalPart1)
	fmt.Printf("\nThe ruck total for part2 is %v", ruckTotalPart2)
	
}

func sortString(w string) []string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return s
}
func makeItemValueMap() map[string]int {
	ruckItemValues := make(map[string]int)
		for i, item := range sortString(LOWERCASE_ABC) {
			ruckItemValues[item] = i + 1	
		}
		for i, item := range sortString(UPPERCASE_ABC) {
			ruckItemValues[item] = i + 27 	
		}
		return ruckItemValues
}

func makeAMap(s string) map[string]int {
	theMap := make(map[string]int)
	for _, item := range sortString(s) {
		theMap[item] = 0
	}
	return theMap
}

func findDuplicateforTwo(s1, s2, combined []string) string {
	//create maps to represent each compartment
	mapCompartOne := make(map[string]int)
	mapCompartTwo := make(map[string]int) 

	for _, item := range s1 {
		mapCompartOne[item] = 0
	}

	for _, item := range s2 {
		mapCompartTwo[item] = 0
	}
	//find duplicate by identifying duplicates in the ruck and 
	//checking to see if those duplicates are in both of the 
	//compartments
	var prevItem string

	duplicateItem := ""

	for _, item := range combined {
		if duplicateItem == "" {
			if prevItem == item {
				_, itemInCompartOne := mapCompartOne[item]
				_, itemInCompartTwo := mapCompartTwo[item]

				if itemInCompartOne && itemInCompartTwo {
					duplicateItem = item
				}
				prevItem = item
			}
			prevItem = item
		} else {
			break
		}
		
	}
	return duplicateItem
}
