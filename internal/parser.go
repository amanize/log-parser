package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

type Predicate = func(templateText string, readedText string) bool

// Tranformer for templates that will be added to already checked templates.
// For example: anagrammas provides to templates for one string.
// Source: asd, anagramma: dsa, checked templates: asd, dsa.
type TemplateTransformer = func(templateText string) []string

func parseFile(title string, path string, predicate Predicate, templateTransformer TemplateTransformer) error {

	var (
		file           []byte
		scannedText    []string
		scannedRunes   [][]int
		sortedRunes    [][]int
		sortedText     []string
		duplicates     []string
		duplicateLines map[string][]int = make(map[string][]int)
		anagrammas []string
		anagrammasLines map[string][]int = make(map[string][]int)
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

	copy(sortedText, scannedText)

	sort.Strings(sortedText)
	duplicates = findAllDuplicates(sortedText)
	for i, str := range scannedText {
		for _, dpl := range duplicates {
			if str == dpl {
				duplicateLines[str] = append(duplicateLines[str], i)
			}

		}
	}

	sortedBytes = make([][]int, len(scannedBytes))
	copy(sortedBytes, scannedBytes)
	sortInt(sortedBytes)
	intDuplicate:=findAllDuplicates(sortedBytes)
	fmt.Println(duplicateLines)
	return nil
}

func btoi(b []byte) []int {
	blen := len(b)
	ints := make([]int, blen)
	for i := 1; i < blen; i++ {
		ints[blen-i] = int(b[i])
	}
	return ints
}

func itob(i []int) []byte{
	ilen:= len(i)
	bytes := make([]byte, ilen)
	for b:=1; b<ilen; b++{
		bytes[ilen-b] = byte(i[])
	}

}

func sortInt(ints [][]int) {
	for _, i := range ints {
		sort.Ints(i)
	}
}

func findAllDuplicates[T comparable](sortedFile []T) []T {
	duplicates := make([]T, 0)
	for i, srt := range sortedFile {
		if i+1 >= len(sortedFile) {
			break
		}
		if srt == sortedFile[i+1] {
			duplicates = append(duplicates, srt)
		}
	}
	return duplicates
}


func Anagram(s string, t string) bool {
	// initializing the variables
	string1 := len(s)
	string2 := len(t)
	fmt.Println("Letter 1 =", s, "\nLetter 2 =", t)
	fmt.Println("Is it an Anagram ?")
	if string1 != string2 {
	   return false
	}
	
	// create and initialize a map anagramMap
	// Using make () function
	anagramMap := make(map[string]int)
	
	// As we already know that make() function
	// always returns a map which is initialized
	// Iterating map using for range loop
	// Traverse the first string and increase the count of each character in the map
	for i := 0; i < string1; i++ {
	   anagramMap[string(s[i])]++
	}
	
	// Traverse the second string and decrease the count of each character in the map
	for i := 0; i < string2; i++ {
	   anagramMap[string(t[i])]--
	}
	
	// Traverse the first string again and if for any character the count
	// is non-zero in the map then return false
	for i := 0; i < string1; i++ {
	   if anagramMap[string(s[i])] != 0 { // if this condition satisfies return false
		  return false
	   }
	}
	// In the end return true
	return true
 }