package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(scanner * bufio.Scanner) string {
	if !scanner.Scan() {
		if scanner.Err() != nil {
			fmt.Println("Error reading input:", scanner.Err())
		}
		os.Exit(0)
	}
	text := strings.TrimSpace(scanner.Text())
	return text
}

func CheckArguments(args []string) int {
	if len(args) != 2 {
		fmt.Println("Please provide a number as command line argument")
		os.Exit(0)
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid command-line argument. Please launch with a valid number.")
		os.Exit(0)
	}
	return id
}

func GetWords(id int) (string, []string) {
	//fileName := "valid-wordle-words.txt" // for local
	fileName := "wordle-words.txt" // for tests
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(0)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	if !fileScanner.Scan() {
		if fileScanner.Err() != nil {
			fmt.Println("Error reading file with word's list:", fileScanner.Err())
		}
		os.Exit(0)
	}

	var words[] string
	for fileScanner.Scan() {
		words = append(words, fileScanner.Text() )

	}

	if id <= 0 || id >= len(words) {
		fmt.Println(Red+"Error number out of range of word's list."+Reset)
		os.Exit(0)
	} 

	return words[id-1], words
}
