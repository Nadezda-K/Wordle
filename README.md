# Wordle

A Go learning project that recreates a word-guessing game in the terminal.

## Task

Build a game where a player guesses a five-letter secret word in up to six valid attempts, receives feedback, and can review their saved statistics.

## What the program does

- Reads a secret-word index from the command line and loads `wordle-words.txt`.
- Asks for a username and accepts five-letter lowercase guesses from the word list.
- Marks letters green for a matching position, yellow for a letter found elsewhere, and gray/white for a letter absent from the secret word.
- Shows remaining attempts and removes absent letters from the displayed alphabet.
- Reveals the answer after a loss and appends the username, secret word, attempt count, and result to `stats.csv`.
- Optionally displays games played, games won, and average attempts for that username.

## Run

With Go installed, run from this folder so the word list can be found:

```bash
go run main.go game.go io.go 10
```

The current implementation uses the argument directly as a slice index while rejecting zero. Use an index from `1` to one less than the number of words; `10` selects the eleventh entry. Letter feedback checks presence rather than repeated-letter counts.

## Learning focus

Command-line arguments, input validation, slices, string processing, ANSI terminal colors, file reading, and CSV-style data storage.
