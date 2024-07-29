package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main() {
    regexStartVowelEndConsonant := regexp.MustCompile(`\b[AEIOUaeiou][a-zA-Z]*[^AEIOUaeiou\s]\b`)

    file, err := os.Open("poem.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()
	
	fmt.Println("Words that start with vowels and end with consonants:")

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		matchesVowelConsonant := regexStartVowelEndConsonant.FindAllString(line, -1)
        if matchesVowelConsonant != nil {
            fmt.Println(matchesVowelConsonant)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}
