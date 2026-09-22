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



func CheckValidWord(guess string, words []string ) bool {
	strRune := []rune(guess)
	if len(strRune) != 5 {
		fmt.Println(Red+"Your guess must be exactly 5 letters."+Reset)
		return false
	}

	for _, r := range strRune {
		if r < 'a' || r > 'z' {
			fmt.Println(Red+"Your guess must only conatin lowercase letters."+Reset)
			return false
		}
	}
	
	allInOne := strings.Join(words, " ")
	if !strings.Contains(allInOne, guess) {
		fmt.Println(Red+"Word not in the list. Please enter a vallid word."+Reset)
		return false
	}
	return true
}

func DeleteFromRemainig(ch string){
	for i, rl := range RemainingLetters {
		if rl == ch {
			RemainingLetters = append(RemainingLetters[:i], RemainingLetters[i+1:]...)
		}
	}
}

func ColorLetters(guess string, secret string) string {
	feedback:= ""
	guessArr := []rune(guess)
	for i, ch := range guessArr {
		inSecret := strings.Contains(secret, string(ch) )
		if inSecret == false {
			DeleteFromRemainig( string(ch) )
		}
		if inSecret == true {
			if string(ch) == string(secret[i]) {
				feedback += Green + strings.ToUpper( string(ch) ) + Reset
			} else {
				feedback += Yellow + strings.ToUpper( string(ch) ) + Reset
			}
		} else {
			feedback += strings.ToUpper( string(ch) )
		}
	}
	return feedback
}


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