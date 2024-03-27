package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
	otherWord + "hq",
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[random.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
