package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
        Reset = "\033[0m"
        Red = "\033[31m"
        Green = "\033[32m"
		Yellow = "\033[33m"
        Gray = "\033[37m"
)

var RemainingLetters = []string{
	"a", "b", "c", "d", "e", "f", "g", 
	"h", "i", "j", "k", "l", "m","n", 
	"o", "p", "q", "r", "s", "t", "u", 
	"v", "w", "x", "y", "z"}


func main() {
	idSecret := CheckArguments(os.Args)

	scanner := bufio.NewScanner(os.Stdin)	
	fmt.Printf("Enter your username: ")
	username := getInput(scanner)

	wordSecret, validWords := GetWords(idSecret)



	fmt.Printf("Welcome to Wordle! Guess the 5-letter word.\n")
	
	var wordGuess string
	winSatuts := "loss"
	numberOfAttempts := 0
	idAttemps := 6
	for i:=1; i<=idAttemps; i++ {
		fmt.Printf("Enter your guess:  ")
		wordGuess = getInput(scanner)

		is_valid := CheckValidWord(wordGuess, validWords)
		if !is_valid {
			i--
			continue
		}

		numberOfAttempts++

		if wordGuess == wordSecret {
			fmt.Println("Congratulations! You've guessed the word correctly.")
			winSatuts = "win"
			break
		}

		wordFeedback := ColorLetters(wordGuess, wordSecret)

		fmt.Printf("Feedback: %s\n", wordFeedback)
		remainingStr := strings.ToUpper(strings.Join(RemainingLetters, " ") )
		fmt.Printf("Remaining letters: %s \n", remainingStr)
		fmt.Printf("Attempts remaining:  %d\n", idAttemps-i)

	}

if winSatuts == "loss" {
    fmt.Printf("Game over. The correct word was: %s\n", wordSecret)
}

	
	// Open file, create it if it does not.
	fileName := "stats.csv"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
    if err != nil {
        fmt.Println(Red+"Error opening/creating file"+Reset, err)
        return
    }
    defer file.Close()

	gameStat := username + ","
	gameStat += wordSecret + "," 
	gameStat += strconv.Itoa(numberOfAttempts) + ","
	gameStat += winSatuts + "\n"
	fmt.Fprint(file, gameStat)



	fmt.Printf("Do you want to see your stats? (yes/no): ")
	statAnswer := getInput(scanner)
	
	if statAnswer == "yes" {
		GetStats(username, fileName)
	}
	
	pressEnter(scanner)
}

func pressEnter(scanner * bufio.Scanner) {
	fmt.Printf("Press Enter to exit...\n")
	for {
		getInput(scanner)
		os.Exit(0)
	}
}

func GetStats(user string, statFile string ) {
	    // Open the file in read-only mode.
        file, err := os.Open(statFile)
        if err != nil {
                fmt.Println(Red+"Error opening file:"+Reset, err)
          }
        defer file.Close()

		fileScanner := bufio.NewScanner(file)
		var statistics []string

		for fileScanner.Scan() {
    		statistics = append(statistics, fileScanner.Text())
		}

		if fileScanner.Err() != nil {
    		fmt.Println(Red+"Error reading statistics file:"+Reset, fileScanner.Err())
    		os.Exit(0)
		}

		var gamesNumber, wonNumber int
		var avgAttemps, temp float64

		for _, n := range statistics{
			userLine := strings.Split(n, ",")
			if userLine[0] == user {
				if userLine[3] == "win" {
					wonNumber++
				}
				gamesNumber++
				temp, _ = strconv.ParseFloat( userLine[2], 64 )
				avgAttemps += temp
			}
		}

	if gamesNumber > 0 {
		avgAttemps /= float64(gamesNumber)
	}

		fmt.Printf("Stats for %s:\n", user)
		fmt.Printf("Games played: %d\n", gamesNumber)
		fmt.Printf("Games won: %d\n", wonNumber)
		fmt.Printf("Average attempts per game: %.2f\n", avgAttemps)
		
        //return statistics

}