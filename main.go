package main

import (
	"fmt"
	"os"

	"github.com/siddhant-vij/Load-Testing-Utility/internal/reporter"
	"github.com/siddhant-vij/Load-Testing-Utility/internal/tester"
)

func main() {
	words, err := tester.LoadWords("resources/search_words.csv")
	if err != nil {
		fmt.Printf("Failed to load words: %v", err)
		os.Exit(1)
	}

	results := tester.RunTests(words, 200)
	reporter.ReportResults(results)
}
