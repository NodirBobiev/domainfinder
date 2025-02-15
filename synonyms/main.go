package main

import (
	"bufio"
	"domain/thesaurus"
	"fmt"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("NINJA_API_KEY")
	thesaurus := &thesaurus.Ninja{APIKey: apiKey}
	// thesaurus := &thesaurus.ChatSynonyms{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for "+word, err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for " + word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
