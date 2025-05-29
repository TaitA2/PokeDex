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
		fmt.Println("Your command was:",cleanInput(userInput)[0])

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
