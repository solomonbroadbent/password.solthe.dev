package main

import (
	"embed"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// this seems really hacky below. probably need to go to s3 solution
// that means vendor lock in? unless use a standardised s3 env variables thing maybe?
// so maybe docker self hosting instead? FUCK THAT HACK AWAY RN

//go:embed assets/*.txt
var CONTENT embed.FS

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

	file, err := CONTENT.ReadFile(filePath)

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

	file, err := CONTENT.ReadFile(filePath)

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

type Response struct {
	Body string `json:"body"`
	// StatusCode int    `json:"statusCode,omitempty"`
}

type Request struct {
	Name string `json:"name"`
}

func Main(request Request) Response {

	words := getWords()
	symbols := getSymbols()

	password := generatePassword(words, symbols, 3)

	return Response{
		Body: password,
	}
}
