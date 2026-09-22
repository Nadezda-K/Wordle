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
        Blue = "\033[34m"
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



	fmt.Printf(" Welcome to Wordle! Guess the 5-letter word.\n")
	
	var wordGuess string
	winSatuts := "loss"
	numberOfAttempts := 1
	idAttemps := 6
	for i:=1; i<=idAttemps; i++ {
		fmt.Printf("Enter your guess:")
		wordGuess = getInput(scanner)

		if wordGuess == wordSecret {
			fmt.Println("Congratulations! You've guessed the word correctrly")
			winSatuts = "win"
			break
		}

		is_valid := CheckValidWord(wordGuess, validWords)
		if !is_valid {
			i--
		}

		wordFeedback := ColorLetters(wordGuess, wordSecret)

		if is_valid {
			fmt.Printf("%s\n",wordGuess)
			fmt.Printf("Feedback: %s\n", wordFeedback)
			remainingStr := strings.ToUpper(strings.Join(RemainingLetters, " ") )
			fmt.Printf("Remaining letters: %s\n", remainingStr)
			fmt.Printf("Attemps remaining: %d\n", idAttemps-i)
		}
		numberOfAttempts++
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
	fmt.Fprintf(file, gameStat)



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
		if !fileScanner.Scan() {
			if fileScanner.Err() != nil {
				fmt.Println(Red+"Error reading file with word's list:"+Reset, fileScanner.Err())
			}
			os.Exit(0)
		}

		var statistics []string
		for fileScanner.Scan() {
			statistics = append(statistics, fileScanner.Text() )
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
			avgAttemps /= float64(gamesNumber)
		}

		fmt.Printf("Stats for %s:\n", user)
		fmt.Printf("Games played: %d:\n", gamesNumber)
		fmt.Printf("Games won: %d:\n", wonNumber)
		fmt.Printf("Average attempts per game: %f:\n", avgAttemps)
		
        //return statistics

}