package main

import (
	"bufio"
	"fmt"
	"os"
//	"strconv"
	"strings"
)

const (
        Reset = "\033[0m"
        Red = "\033[31m"
        Green = "\033[32m"
		Yellow = "\033[33m"
        Blue = "\033[34m"
)

var RemainingLetters = []string{
	"a", "b", "c", "d", "e", "f", "g", 
	"h", "i", "j", "k", "l", "m","n", 
	"o", "p", "q", "r", "s", "t", "u", 
	"v", "w", "x", "y", "z"}


func main() {
	idSecret := CheckArguments(os.Args)
	wordSecret, validWords := GetWords(idSecret)
	fmt.Println(wordSecret)

	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Enter your username:")

	username := getInput(scanner)
	fmt.Println("Username is", username)


	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
	
	var wordGuess string
	idAttemps := 2
	for i:=1; i<=idAttemps; i++ {
		fmt.Printf("Enter your guess: ")
		wordGuess = getInput(scanner)
		fmt.Printf("%s\n",wordGuess)

		if wordGuess == wordSecret {
			fmt.Println("Congratulations! You've guessed the word correctrly")
			break
		}

		is_valid := CheckValidWord(wordGuess, validWords)
		if !is_valid {
			i--
		}

		wordFeedback := ColorLetters(wordGuess, wordSecret)

		if is_valid {
			fmt.Printf("Feedback: %s\n", wordFeedback)
			remainingStr := strings.ToUpper(strings.Join(RemainingLetters, " ") )
			fmt.Printf("Remaining letters: %s\n", remainingStr)
			fmt.Printf("Attemps remaining: %d\n", idAttemps-i)
		}		
	}


	fmt.Printf("Do you want to see your stats? (yes/no):")
	statAnswer := getInput(scanner)
	fmt.Printf("%s\n",statAnswer)

}