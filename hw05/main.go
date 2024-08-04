package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextEditor struct {
	TextLines []string
	WordIndex map[string][]int
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		WordIndex: make(map[string][]int),
	}
}

func (te *TextEditor) IndexTextByWord() {
	te.WordIndex = make(map[string][]int) 

	for i, line := range te.TextLines {
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			if _, exists := te.WordIndex[word]; !exists {
				te.WordIndex[word] = []int{}
			}
			te.WordIndex[word] = append(te.WordIndex[word], i)
		}
	}
}

func (te *TextEditor) SearchAllLinesByWord(word string) []int {
	word = strings.ToLower(word)
	return te.WordIndex[word]
}

func main() {
	editor := NewTextEditor()

	editor.TextLines = []string{
		"Black tea",
		"Green tea",
		"Matcha",
		"Cafe latte",
		"Latte macchiato",
	}

	editor.IndexTextByWord()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word to search: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	lineNumbers := editor.SearchAllLinesByWord(word)

	fmt.Printf("Lines containing the word '%s':\n", word)
	if len(lineNumbers) == 0 {
		fmt.Println("No lines found containing the word.")
	} else {
		for _, lineNum := range lineNumbers {
			fmt.Printf("Line %d: %s\n", lineNum, editor.TextLines[lineNum])
		}
	}
}
