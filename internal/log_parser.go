package internal

import (
	"fmt"
	"log-parser/arguments"
	"os"
	"strconv"
	"strings"
)

const Version = "0.0.2"

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

	if err := parseFile(duplicateMode, path, duplicatePredicate, duplicateTransformer); err != nil {
		return fmt.Errorf("error reading file. Reason: %s", err)

	}
	return nil
}

func findAnagrammas(path string) error {

	if err := parseFile(anagrammasMode, path, anagrammasPredicate, anagrammasTransformer); err != nil {
		return fmt.Errorf("error reading file. Reason: %s", err)
	}
	return nil
}

func printResult(title string, result SearchResult) {
	lineNumberPrinter := func(numbers []int) string {
		var numbersString []string
		for _, number := range numbers {
			numbersString = append(numbersString, strconv.Itoa(number))
		}
		return strings.Join(numbersString, ", ")
	}

	fmt.Fprintf(os.Stdout, "Record: %s, line:%d, %s on lines: %s\n", result.template, result.line, title, lineNumberPrinter(result.duplicates))

}
