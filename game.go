package main

import (
	"fmt"
	"strings"
)

func CheckValidWord(guess string, words []string ) bool {
	strRune := []rune(guess)
	if len(strRune) != 5 {
		fmt.Printf("Your guess must be exactly 5 letters long.\n")
		return false
	}

	for _, r := range strRune {
		if r < 'a' || r > 'z' {
			fmt.Printf("Your guess must only contain lowercase letters.\n")
			return false
		}
	}
	
	// allInOne := strings.Join(words, " ")
	// if !strings.Contains(allInOne, guess) {
	// 	fmt.Printf("Word not in list. Please enter a valid word.\n")
	// 	return false
	// }
	// return true

	wordExists := false
	
	for _, word := range words {
		if word == guess {
			wordExists = true
			break
		}
	}

	if !wordExists {
    	fmt.Printf("Word not in list. Please enter a valid word.\n")
    	return false
	}
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
			feedback += Gray + strings.ToUpper( string(ch) ) + Reset
		}
	}
	return feedback
}
