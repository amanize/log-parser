package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"log-parser/arguments"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

func parseFile(mode string, path string) error {

	var (
		file           []byte
		scannedText    []string
		sortedText     []string
		duplicates     []string
		duplicateLines map[string][]int = make(map[string][]int)
		err            error
	)
	if file, err = os.ReadFile(path); err != nil {
		return err
	}
	reader := bytes.NewReader(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		scannedText = append(scannedText, scanner.Text())
	}
	sortedText = make([]string, len(scannedText))
	if mode == arguments.Duplicates {
		copy(sortedText, scannedText)
		sort.Strings(sortedText)
		duplicates = findAllDuplicates(sortedText)
		fmt.Println("Search for duplicates")
		duplicateLines = getDuplicateLines(scannedText, duplicates)
		printResult(duplicateLines)
	}
	if mode == arguments.Anagrammas {
		fmt.Println("Search for anagrammas")
		anagrammas := findAllAnagrammas(scannedText)
		printResult(anagrammas)
	}
	return nil
}

func printResult(result map[string][]int) {
	for keys, group := range result {
		if len(group) > 1 {
			fmt.Println(keys, ":", arrayToString(group, ","))
		}
	}
}

func getDuplicateLines(scannedText []string, duplicates []string) map[string][]int {
	duplicateLines := make(map[string][]int)
	for i, str := range scannedText {
		if slices.Contains(duplicates, str) {
			duplicateLines[str] = append(duplicateLines[str], i+1)
		}
	}
	return duplicateLines
}

func findAllDuplicates(sortedText []string) []string {
	result := make([]string, 0)
	for i, srt := range sortedText {
		if i+1 >= len(sortedText) {
			break
		}
		if srt == sortedText[i+1] {
			result = append(result, srt)
		}
	}
	return removeDuplicate(result)
}

func findAllAnagrammas(sortedText []string) map[string][]int {
	anagrams := make(map[string][]int)
	for i, t := range sortedText {
		sortedLine := sortString(t)
		anagrams[sortedLine] = append(anagrams[sortedLine], i+1)
	}
	return anagrams
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
