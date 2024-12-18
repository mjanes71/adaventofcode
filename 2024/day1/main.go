// order both lists smallest to largest
// compare the first num from list 1 with first num from list 2
// etc etc
// add all the distances up to get one final number

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
    first, second, err := processTextFile("data.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("First List  length:", len(first))
    fmt.Println("Second List length:", len(second))

	firstSorted := sortList(first)
	secondSorted := sortList(second)

	answer := compare(firstSorted, secondSorted)
	fmt.Println("First part Answer:", int(answer))

	similarity := calculateSimilarity(first, second)
	fmt.Println("Similarity Answer:", similarity)
}

func processTextFile(filePath string) ([]string, []string, error) {
    var firstList, secondList []string

    file, err := os.Open(filePath)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) != "" { // Check if the line has text
            parts := strings.SplitN(line, "   ", 2) // Split the line into two parts using 3 spaces as the delimiter
            if len(parts) == 2 {
                firstList = append(firstList, strings.TrimSpace(parts[0]))
                secondList = append(secondList, strings.TrimSpace(parts[1]))
            }
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, nil, err
    }

    return firstList, secondList, nil
}

// Helper function to sort a list of strings
func sortList(list []string) []string {
    sortedList := make([]string, len(list))
    copy(sortedList, list)
    sort.Strings(sortedList)
    return sortedList
}

// Helper function to calculate the absolute values
func compare(list1, list2 []string) float64 {
	var total float64
	for i := 0; i < len(list1); i++ {
		num1, _ := strconv.ParseFloat(list1[i], 64)
		num2, _ := strconv.ParseFloat(list2[i], 64)
		total += math.Abs(num1 - num2)
	}
	return total
}

func countOccurrences(list []string, target string) int {
    count := 0
    for _, item := range list {
        if item == target {
            count++
        }
    }
    return count
}

func calculateSimilarity(list1, list2 []string) int {
	var total int
	for i := 0; i < len(list1); i++ {
		currentNum, _ := strconv.Atoi(list1[i])
		total += countOccurrences(list2, list1[i]) * currentNum
	}
	return total
}
