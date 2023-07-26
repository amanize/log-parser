package internal

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

type Predicate = func(templateText string, readedText string) bool

// Tranformer for templates that will be added to already checked templates.
// For example: anagrammas provides to templates for one string.
// Source: asd, anagramma: dsa, checked templates: asd, dsa.
type TemplateTransformer = func(templateText string) []string

func parseFile(path string, predicate Predicate, templateTransformer TemplateTransformer) (map[string][]int, error) {
	var (
		index            int = 1
		file             *os.File
		parseResult      []int
		checkedTemplates []string
		result           map[string][]int = make(map[string][]int)
		err              error
	)
	if file, err = os.Open(path); err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		template := scanner.Text()
		// Already find this template
		if slices.Contains(checkedTemplates, template) {
			index += 1
			continue
		}
		if parseResult, err = find(path, index, template, predicate); err != nil {
			return nil, err
		}
		if len(parseResult) > 0 {
			checkedTemplates = append(checkedTemplates, templateTransformer(template)...)
			result[fmt.Sprintf("%s, line:%d", template, index)] = parseResult
		}
		index += 1
	}
	return result, nil
}

func find(path string, i int, templVal string, predicate Predicate) (results []int, err error) {
	var (
		index int = 1
		file  *os.File
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rdText := scanner.Text()
		if index != i && predicate(templVal, rdText) {
			results = append(results, index)
		}
		index += 1
	}
	return
}
