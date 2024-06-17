package main

import "fmt"

func main(){

	menu := map[string]float64{
		"black tea": 3.75,
		"green tea": 3.75,
		"white tea": 3.75,
		"herbal tea": 4,
		"latte": 4.5,
		"capuccino": 4.25,
		"ice latte": 4.5,
	}

	fmt.Println("What are you looking for?")

	var input string
	fmt.Scan(&input)

	if menu[input] == 0 {
		fmt.Printf("Sorry, we dont have %v on the menu.", input)
	} else {
		fmt.Printf("One %v costs %v USD", input, menu[input])
	}

}