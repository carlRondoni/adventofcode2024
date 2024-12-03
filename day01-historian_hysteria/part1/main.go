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
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicf("error on get file content: %s", err)
	}
	defer file.Close()
	
	var leftList, rightList []int
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])

			leftList = append(leftList, num1)
			rightList = append(rightList, num2)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	sort.Ints(leftList)
	sort.Ints(rightList)
	
	sum := 0
	for i, pos := range leftList {
		diff := pos - rightList[i]
		if diff < 0 {
			diff = -diff
		}
		
		sum += diff
	}
	
	log.Printf("total position: %v", sum)
}
