package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"net"
	"log"
	"context"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	
	pb "github.com/solomonbroadbent/password.solthe.dev/project/generated"
	// "github.com/solomonbroadbent/password.solthe.dev/project/generated"
	// "generated/password-generator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

func getAPassword() string {

	words := getWords()
	symbols := getSymbols()

	password := generatePassword(words, symbols, 3)

	os.Stdout.WriteString(password)

	return password;
}


type server struct {
	pb.UnimplementedPasswordGeneratorServer
}

func (s *server) GeneratePassword(ctx context.Context, req *pb.PasswordRequest) (*pb.PasswordResponse, error) {
	// Simple implementation: Always return the same password for testing
	return &pb.PasswordResponse{
		Password: getAPassword(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPasswordGeneratorServer(s, &server{})

	// Enable gRPC server reflection
	reflection.Register(s)

	log.Println("gRPC server is running on port 80")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

