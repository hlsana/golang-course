package main

import "fmt"

const (
	Left    = "left"
	Right   = "right"
	Knife   = "knife"
	Lighter = "lighter"
	Yes     = "yes"
	No      = "no"
)

type Character struct {
	Name       string
	HasKnife   bool
	HasLighter bool
	HasRope    bool
	IsInjured  bool
	IsAlive    bool
	Choice
}

type Choice struct {
	ItemChoice      string
	DirectionChoice string
	SearchCamp      string
}

func main() {

	Player1 := Character{
		IsAlive: true,
	}

	println("You have woken up in a cave and remember nothing, but your name. What is it?")

	fmt.Scan(&Player1.Name)

	fmt.Println("Hi", Player1.Name, "!")

	fmt.Println("You find a bag by your side with a knife and a lighter. You can take only one object with you. What do you choose?")

	for {

		fmt.Scan(&Player1.ItemChoice)

		switch Player1.ItemChoice {
		case Knife:
			Player1.HasKnife = true
		case Lighter:
			Player1.HasLighter = true
		default:
			println("Try again!")
		}
		if Player1.ItemChoice == Knife || Player1.ItemChoice == Lighter {
			break
		}
	}
	fmt.Println("You grab the", Player1.ItemChoice, "and decide to get out of the cave.")

	if !Player1.HasLighter {
		fmt.Println("You don't have a light source, so you stumble in the dark until you see a passage to the left.")
		Player1.DirectionChoice = Left
	} else {
		fmt.Println("The light coming from the lighter illuminates your path ahead. Soon you see two passageways, to your left and right. Which one do you choose?")
		fmt.Scan(&Player1.DirectionChoice)
	}

	for {
		switch Player1.DirectionChoice {
		case Left:
		case Right:
		default:
			println("Try again!")
			fmt.Scan(&Player1.DirectionChoice)
		}
		if Player1.DirectionChoice == Left || Player1.DirectionChoice == Right {
			break
		}
	}

	switch Player1.DirectionChoice {
	case Left:
		fmt.Println("The passage gets you out of the cave, and you see an abandoned campsite. You decide to go there and search for something useful.")
		fmt.Println("When you approach the campsite, a wild dog suddenly attacks you.")
		if Player1.HasKnife {
			fmt.Println("Luckily, you have a knife, so you fight back, and the dog runs away scared.")
		} else {
			fmt.Println("Unfortunately, you don't have a knife to protect yourself. You still manage to scare the dog away, but it bites you painfully before running away.")
			Player1.IsInjured = true
		}

		for {
			fmt.Println("More dogs might come and attack you if you don't leave. Do you still want to search the campsite?")
			fmt.Scan(&Player1.SearchCamp)

			switch Player1.SearchCamp {
			case Yes:
				fmt.Println("You look around and find a rope. It might be useful!")
				Player1.HasRope = true
			case No:
				fmt.Println("You decide to be cautious and keep moving.")
			default:
				println("Try again!")
			}
			if Player1.SearchCamp == Yes || Player1.SearchCamp == No {
				break
			}
		}

		fmt.Println("You walk for a long time until you finally see a town, but it lays down in the valley. You would need to get down the steep rocky hill you are on.")
		if Player1.HasRope {
			fmt.Println("The rope was a lucky find indeed! You tie it to a nearby tree and begin your descent.")
			if !Player1.IsInjured {
				fmt.Println("You get down safely and walk to the town with relief. You are finally safe.")
			} else {
				fmt.Println("Your injury makes you slip, so you fall down and die. Better luck next time!")
				Player1.IsAlive = false
			}
		} else {
			fmt.Println("You don't have a rope to help you get down safely, so you fall down and die. Better luck next time!")
			Player1.IsAlive = false
		}
	case Right:
		fmt.Println("You keep going until you find yourself in another cave chamber full of spiders. One of them climbs up your leg and bites you.")
		fmt.Println("You suddenly feel dizzy and cannot walk. Your vision blackens. The spider was deadly poisonous.")
		fmt.Println("Your death is quick. If only you have made a different choice...")
		Player1.IsAlive = false

	}

	if Player1.IsAlive {
		fmt.Println("Congratulations, you won!")
	} else {
		fmt.Println("GAME OVER")
	}
}

