package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main() {
    phoneRegex := regexp.MustCompile(`\b(\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4})\b`)

    file, err := os.Open("contacts.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        matches := phoneRegex.FindAllString(line, -1)
        if matches != nil {
            fmt.Println("Found phone numbers:", matches)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}
