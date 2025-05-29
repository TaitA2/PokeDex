package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("PokeDex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		userCommand := cleanUserInput[0]
		fmt.Println("Your command was:", userCommand)

	}
}

func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	var cleanWords []string
	for _, word := range words {
		if word != "" {
			cleanWords = append(cleanWords, strings.ToLower(word))
	}
}
	return cleanWords
}
