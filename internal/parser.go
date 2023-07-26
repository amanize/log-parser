package internal

import (
	"bufio"
	"bytes"
	"os"
)

type Predicate = func(templateText string, readedText string) bool

// Tranformer for templates that will be added to already checked templates.
// For example: anagrammas provides to templates for one string.
// Source: asd, anagramma: dsa, checked templates: asd, dsa.
type TemplateTransformer = func(templateText string) []string

func parseFile(title string, path string, predicate Predicate, templateTransformer TemplateTransformer) error {
	const (
		maxWorks   = 10
		maxWorkers = 3
	)
	var (
		file []byte
		err  error
	)
	if file, err = os.ReadFile(path); err != nil {
		return err
	}
	works := make(chan SearchWork, maxWorks)
	results := make(chan SearchResult, maxWorks)
	scanner := bufio.NewScanner(bytes.NewReader(file))
	go findWorker(file, works, predicate, templateTransformer, results)
	scanFile(scanner, works)
	close(works)
	for result := range results {
		printResult(title, result)
	}
	return nil
}

func scanFile(scanner *bufio.Scanner, works chan<- SearchWork) {
	index := 1
	for scanner.Scan() {
		template := scanner.Text()
		works <- SearchWork{
			template: template,
			line:     index,
		}
		index += 1
	}

}
