package main

import (
	"math/rand"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// hermitdave_frequencywords_en_50k.txt
// hermitdave_frequencywords_es_50k.txt
// raw_dict_en.txt
// raw_dict_es.txt

func getRandomItem(words []string) string {
	randomNumber := rand.Intn(len(words))

	return words[randomNumber]
}

func generateRandomSymbol() string {
	return "@"
}

func generatePassword(words []string, symbols []string, wordCount int) string {
	password := []string{}

	for wordCounter := 0; wordCounter < wordCount; wordCounter++ {
		randomWord := getRandomItem(words)
		capitalisedRandomWord := cases.Title(language.English).String(randomWord)

		password = append(password, capitalisedRandomWord)
	}

	randomSymbol := getRandomItem(symbols)
	password = append(password, randomSymbol+randomSymbol)

	return strings.Join(password, "-")
}

func getWords() []string {
	filePath := "assets/raw_dict_en.txt"

	file, err := os.ReadFile(filePath)

	if err != nil {
		os.Stderr.WriteString("error opening file")
		return nil
	}

	if err != nil {
		os.Stderr.WriteString("error reading file")
		return nil
	}

	return strings.Split(string(file), "\n")
}

func getSymbols() []string {
	filePath := "assets/symbols.txt"

	file, err := os.ReadFile(filePath)

	if err != nil {
		os.Stderr.WriteString("error opening file")
		return nil
	}

	if err != nil {
		os.Stderr.WriteString("error reading file")
		return nil
	}

	return strings.Split(string(file), "\n")
}

func main() {

	words := getWords()
	symbols := getSymbols()

	os.Stdout.WriteString(generatePassword(words, symbols, 3))
}
