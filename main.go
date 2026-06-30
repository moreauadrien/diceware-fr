package main

import (
	"crypto/rand"
	_ "embed"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

//go:embed list.txt
var wordList string

var words []string

func init() {
	words = strings.Split(strings.TrimSpace(wordList), "\n")
}

func main() {
	size := flag.Int("size", 6, "number of words")
	separator := flag.String("separator", "-", "word separator")
	flag.Parse()

	if *size < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [--size N] [--separator SEP]\n", os.Args[0])
		os.Exit(1)
	}

	passphrase := make([]string, *size)
	max := big.NewInt(int64(len(words)))
	for i := range passphrase {
		idx, err := rand.Int(rand.Reader, max)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		passphrase[i] = words[idx.Int64()]
	}

	fmt.Println(strings.Join(passphrase, *separator))
}
