package internal

import (
	"bufio"
	"bytes"
)

type SearchWork struct {
	template string
	line     int
}

type SearchResult struct {
	template   string
	line       int
	duplicates []int
}

func findWorker(file []byte, works <-chan SearchWork, predicate Predicate, templateTransformer TemplateTransformer, results chan<- SearchResult) {
	for work := range works {
		parseResult, err := find(file, work, predicate)
		if err != nil {
			return
		}
		if len(parseResult) > 0 {
			results <- SearchResult{
				template:   work.template,
				line:       work.line,
				duplicates: parseResult,
			}
		}
	}
	close(results)
}

func find(file []byte, sw SearchWork, predicate Predicate) (results []int, err error) {
	var (
		index int = 1
	)
	scanner := bufio.NewScanner(bytes.NewReader(file))
	for scanner.Scan() {
		rdText := scanner.Text()
		if index != sw.line && predicate(sw.template, rdText) {
			results = append(results, index)
		}
		index += 1
	}
	return
}
