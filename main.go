package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	folder, keywords, top := getParameters()
	findJSON(folder, keywords, top)
}

// getParameters gets all user input
func getParameters() (string, []string, int) {
	input := flag.String("input", "", "Folder having JSON files")
	keywords := flag.String("keywords", "", "Keywords to be found in JSON files")
	top := flag.Int("top", 1, "Top k files to be found")
	flag.Parse()

	if *input == "" {
		log.Fatal("No input folder was given.")
	}
	if *keywords == "" {
		log.Fatal("keywords are not specified.")
	}
	return *input, strings.Split(*keywords, ","), *top
}