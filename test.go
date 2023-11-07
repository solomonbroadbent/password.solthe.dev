package main

import (
	"math/rand"
	"os"
	"strconv"
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

func generateRandomNumber() string {
	return strconv.Itoa(rand.Intn(9))
}

func generatePassword(words []string, symbols []string, wordCount int) string {
	password := ""

	for wordCounter := 0; wordCounter < wordCount; wordCounter++ {
		randomWord := getRandomItem(words)
		capitalisedRandomWord := cases.Title(language.English).String(randomWord)

		password += capitalisedRandomWord + "-"
	}

	password += generateRandomNumber() + generateRandomNumber()

	randomSymbol := getRandomItem(symbols)
	password += randomSymbol + randomSymbol

	return password
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
