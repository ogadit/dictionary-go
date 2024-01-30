package main

import (
	"log"
	"os"
	"strings"

	"osamagadit.com/dict/internal/cli"
	"osamagadit.com/dict/internal/dictionary"
)

func main() {
	word := strings.Join(os.Args[1:], " ")

	def, err := dictionary.FetchResponse(word)
	if err != nil {
		log.Fatalf("Error getting a response\n %v", err)
	}

	cli.Print(&def)
}
