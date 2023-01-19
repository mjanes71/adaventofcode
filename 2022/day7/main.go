package main

import (
	"bufio"
	//"fmt"
	"log"
	"os"
	//"strconv"
	//"strings"
)

func main(){
	//read a line from the file
	file, err := os.Open("day4_data.txt")
    if err != nil {
        log.Fatalf("Try opening this file with the secret password next time.")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//get line from file
		//line := strings.TrimSpace(scanner.Text())
		//parse line
		//$ lines are commands
		//dir lines are directories
		//lines starting with ints are files with sizes
		//gotta parse the lines and map out the file structure
		//maybe a map?
	}

}

func decideLineType(line string) string {
	lineType := "file"
	switch {
	case string(line[0]) == "$":
		lineType = "command"
	case string(line[0:3]) == "dir":
		lineType = "directory"
	}
	return lineType
}

//takes a map and a line
//strips line to grab directory we are changing into
//check map to see if a key or value exists with that name

func handleCommand()
