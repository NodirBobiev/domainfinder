package main

import (
	"bufio"
	"domain/whois"
	"fmt"
	"log"
	"os"
	"time"
)

var marks = map[bool]string{true: "✅", false: "❌"}

func main() {
	apiKey := os.Getenv("NINJA_API_KEY")
	whoisServer := whois.Ninja{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := whoisServer.Exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(marks[!exist])
		time.Sleep(1 * time.Second)
	}
}
