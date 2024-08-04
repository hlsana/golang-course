package main

import (
	"fmt"
	"os"

	"password-manager/internal/manager"
	"password-manager/internal/storage"
)

func main() {
	storage := storage.NewStorage("passwords.txt")
	manager := manager.NewManager(storage)

	for {
		fmt.Println("\nPassword Manager")
		fmt.Println("1. List passwords")
		fmt.Println("2. Save password")
		fmt.Println("3. Get password")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			manager.ListPasswordNames()
		case 2:
			manager.SavePassword()
		case 3:
			manager.GetPassword()
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}