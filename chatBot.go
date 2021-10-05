package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isEscapeText(text string) bool {
	if text == "q" || text == "bye" || text == "quit" || text == "exit" {
		return true
	}
	return false
}

func main() {

	data := map[string]string{
		"hi":          "Hello",
		"hello":       "Hi",
		"how are you": "i'm fine and you?",
	}

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for !isEscapeText(text) {
		fmt.Print("Me: ")
		scanner.Scan()
		text = scanner.Text()
		text = strings.ToLower(text)
		text = strings.TrimSpace(text)
		if isEscapeText(text) {
			fmt.Println("Bot: Have a Great Day")
			continue
		}
		if res, found := data[text]; found {
			fmt.Println("Bot: ", res)
		} else {
			fmt.Println("I didn't understand you! ")
		}
	}
}
