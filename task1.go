package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	names := make([]string, 0, 3)

	myfile, err := os.Open("menu.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer myfile.Close()

	scanner := bufio.NewScanner(myfile)
	for scanner.Scan() {
		Data := strings.Split(scanner.Text(), "\n")
		names = append(names, Data...)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Println("What are you looking for?")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	i := sort.Search(len(names), func(i int) bool { return input.Text() <= names[i] })
	if i < len(names) && names[i] == input.Text() {
	fmt.Printf("Found %s at index %d in %v.\n", input.Text(), i, names)
	} else {
	fmt.Printf("Did not find %s in %v.\n", input.Text(), names)
	}

}
