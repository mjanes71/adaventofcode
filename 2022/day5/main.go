package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stack1 = []string{"S", "M", "R", "N", "W", "J", "V", "T"}
var stack2 = []string{"B", "W", "D", "J", "Q", "P", "C", "V"}
var stack3 = []string{"B", "J", "F", "H", "D", "R", "P"}
var stack4 = []string{"F", "R", "P", "B", "M", "N", "D"}
var stack5 = []string{"H", "V", "R", "P", "T", "B"}
var stack6 = []string{"C", "B", "P", "T"}
var stack7 = []string{"B", "J", "R", "P", "L"}
var stack8 = []string{"N", "C", "S", "L", "T", "Z", "B", "W"}
var stack9 = []string{"L", "S", "G"}
var stacks = map[int][]string{1:stack1, 2:stack2, 3:stack3, 4:stack4, 5:stack5, 6:stack6, 7:stack7, 8:stack8, 9:stack9}
var stacksPt2 = map[int][]string{1:stack1, 2:stack2, 3:stack3, 4:stack4, 5:stack5, 6:stack6, 7:stack7, 8:stack8, 9:stack9}

// var stack1 = []string{"Z", "N"}
// var stack2 = []string{"M", "C", "D"}
// var stack3 = []string{"P"}
// var stacks = map[int][]string{1:stack1, 2:stack2, 3:stack3}


func main() {
	//read a line from the file
	file, err := os.Open("day5_data.txt")
	if err != nil {
		log.Fatalf("Try opening this file with the secret password next time.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		//trim up the line to get the 3 coordinates
		move, from, to := handleString(line)
		//pop and push onto the proper arrays
		stacks[from], stacks[to] = moveCrates9000(stacks[from], stacks[to], move)
		stacksPt2[from], stacksPt2[to] = moveCrates9001(stacksPt2[from], stacksPt2[to], move)
	}
	//grab the last item from every slice
	fmt.Println("Part 1 answer: ")
	for key, element := range stacks {
		fmt.Printf("\nElement %v is on top of stack %v", element[len(element)-1], key)
	}

	fmt.Println("\nPart 2 answer: ")
	for key, element := range stacksPt2 {
		fmt.Printf("\nElement %v is on top of stack %v", element[len(element)-1], key)
	}

}

func handleString(s string) (int, int, int) {
	//remove move_
	s = strings.Replace(s, "move ", "", -1)
	//replace from and to with ,
	s = strings.Replace(s, " from ", ",", -1)
	s = strings.Replace(s, " to ", ",", -1)
	//split on , to make a slice
	sliceAndDice := strings.Split(s, ",")
	//get all 3 values from each index
	move, _ := strconv.Atoi(sliceAndDice[0])
	from, _ := strconv.Atoi(sliceAndDice[1])
	to, _ := strconv.Atoi(sliceAndDice[2])

	return move, from, to
}

func moveCrates9000 (fromStack, toStack []string, move int) ([]string, []string) {
	//take len-1 from the "from" stack and append it to the "to" stack 
	//as many times as move dictates
	
	for i := 0; i < move; i++ {
		lastOnFrom := fromStack[len(fromStack)-1]
		toStack =  append(toStack, lastOnFrom)
		fromStack = fromStack[:len(fromStack)-1]
	}
	return fromStack, toStack

}

func moveCrates9001 (fromStack, toStack []string, move int) ([]string, []string) {
	//take len-1 from the "from" stack and append it to the "to" stack 
	//as many times as move dictates
	
	chunkOfFrom := fromStack[len(fromStack)-move:]
	toStack =  append(toStack, chunkOfFrom...)
	fromStack = fromStack[:len(fromStack)-move]
	return fromStack, toStack

}
