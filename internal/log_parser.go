package internal

import (
	"fmt"
	"log-parser/arguments"
	"os"
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
