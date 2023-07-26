package internal

func duplicatePredicate(templateText string, readedText string) bool {
	return templateText == readedText
}

func anagrammasPredicate(templateText string, readedText string) bool {
	return templateText == reverseSwapRunes(readedText)
}
