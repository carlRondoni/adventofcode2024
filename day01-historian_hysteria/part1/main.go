package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
	"fmt"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panicf("error on get file content: %s", err)
	}
	defer file.Close()
	
	// separarlo en
	// left list
	// right list
	var leftList, rightList []int
	
	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split each line into two numbers
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			// Convert strings to integers
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])

			// Append to respective lists
			leftList = append(leftList, num1)
			rightList = append(rightList, num2)
		}
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	// sort each of them
	sort.Ints(leftList)
	sort.Ints(rightList)
	
	sum := 0
	for i, pos := range leftList {
		// x postion on the left - x position on the right = diff
		diff := pos - rightList[i]
		if diff < 0 {
			diff = -diff
		}
		
		// total distance = sum stored diffs
		sum += diff
	}
	
	log.Printf("total position: %v", sum)
	
}
