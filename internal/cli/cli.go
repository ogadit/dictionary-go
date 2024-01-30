package cli

import (
	"fmt"
	"strings"

	"osamagadit.com/dict/internal/dictionary"
)

// Prints the first N definitions
func printNDefinitions(r *dictionary.Response, n int) {
	for i, meaning := range r.Meanings {
		fmt.Printf("%d. %s\n", i+1, meaning.PartOfSpeech)
		for j, definition := range meaning.Definitions {
			if j < n-1 {
				fmt.Printf("  - %s\n", definition.Text)
				if definition.Example != "" {
					fmt.Printf("  Example: \"%s\"\n\n", definition.Example)
				}
			}
		}
	}
}

// Prints the word titlecased, and the phonectic in brackets along with meanings
func Print(r *dictionary.Response) {
	word := strings.ToUpper(string(r.Word[0])) + r.Word[1:]
	if phonetic, err := r.GetPhonetic(); err == nil {
		fmt.Printf("%s (%s)\n", word, phonetic)
	} else {
		fmt.Println(word)
	}
	printNDefinitions(r, 3)
}
