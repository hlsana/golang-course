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

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  ./password-manager save <name> <password>")
		fmt.Println("  ./password-manager get <name>")
		fmt.Println("  ./password-manager get-all")
		return
	}

	command := os.Args[1]

	switch command {
	case "save":
		if len(os.Args) != 4 {
			fmt.Println("Usage: ./password-manager save <name> <password>")
			return
		}
		name := os.Args[2]
		password := os.Args[3]
		err := manager.SavePassword(name, password)
		if err != nil {
			fmt.Println("Error saving password:", err)
		} else {
			fmt.Println("Password saved successfully!")
		}
	case "get":
		if len(os.Args) != 3 {
			fmt.Println("Usage: ./password-manager get <name>")
			return
		}
		name := os.Args[2]
		manager.GetPassword(name)
	case "get-all":
		manager.ListPasswordNames()
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Usage:")
		fmt.Println("  ./password-manager save <name> <password>")
		fmt.Println("  ./password-manager get <name>")
		fmt.Println("  ./password-manager get-all")
	}
}
