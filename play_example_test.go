package play_test

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/wilhelm-murdoch/go-play"
)

func ExampleClient_Fetch_string() {
	code := []byte(`package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}`)

	client := play.NewClient(play.ConfigDefault)
	share, err := client.Fetch(code)
	if err != nil {
		log.Fatal(err)
	}

	pattern := regexp.MustCompile(playUrlPattern)

	fmt.Println(pattern.FindStringIndex(share) != nil)

	// Output:
	// true
}

func ExampleClient_Fetch_file() {
	client := play.NewClient(play.ConfigDefault)

	code, err := os.ReadFile("play.go")
	if err != nil {
		log.Fatal(err)
	}

	share, err := client.Fetch(code)
	if err != nil {
		log.Fatal(err)
	}

	pattern := regexp.MustCompile(playUrlPattern)

	fmt.Println(pattern.FindStringIndex(share) != nil)

	// Output:
	// true
}

func ExampleClient_FetchGroup_file() {
	var (
		bytes  [][]byte
		files  = []string{"play.go", "config.go"}
		client = play.NewClient(play.ConfigDefault)
	)

	for _, file := range files {
		code, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		bytes = append(bytes, code)
	}

	shares, err := client.FetchGroup(bytes)
	if err != nil {
		log.Fatal(err)
	}

	pattern := regexp.MustCompile(playUrlPattern)

	for _, share := range shares {
		fmt.Println(pattern.FindStringIndex(share) != nil)
	}

	// Output:
	// true
	// true
}
