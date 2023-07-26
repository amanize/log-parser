package internal

import (
	"fmt"
	"log-parser/arguments"
	"os"
	"strconv"
	"strings"
)

const Version = "0.0.1"

func Start(args *arguments.Arguments) {
	fmt.Println("log-parser", Version)
	switch args.Mode {
	case arguments.Duplicates:
		if err := findDuplicates(args.Path); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	case arguments.Anagrammas:
		if err := findAnagrammas(args.Path); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	}
}

func findDuplicates(path string) error {
	var (
		result map[string][]int
		err    error
	)
	if result, err = parseFile(path, duplicatePredicate, duplicateTransformer); err != nil {
		return fmt.Errorf("error reading file. Reason: %s", err)

	}
	printResult(duplicateMode, result)
	return nil
}

func findAnagrammas(path string) error {
	var (
		result map[string][]int
		err    error
	)
	if result, err = parseFile(path, anagrammasPredicate, anagrammasTransformer); err != nil {
		return fmt.Errorf("error reading file. Reason: %s", err)

	}
	printResult(anagrammasMode, result)
	return nil
}

func printResult(title string, result map[string][]int) {
	lineNumberPrinter := func(numbers []int) string {
		var numbersString []string
		for _, number := range numbers {
			numbersString = append(numbersString, strconv.Itoa(number))
		}
		return strings.Join(numbersString, ", ")
	}

	for text, lines := range result {
		fmt.Fprintf(os.Stdout, "Record: %s, %s on lines: %s\n", text, title, lineNumberPrinter(lines))
	}

}
