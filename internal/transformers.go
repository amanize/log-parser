package internal

func duplicateTransformer(template string) []string {
	// No transformation needed.
	return []string{template}
}

func anagrammasTransformer(template string) []string {
	// Need straight and reverse template.
	return []string{template, reverseSwapRunes(template)}
}
