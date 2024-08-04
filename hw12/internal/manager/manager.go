package manager

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"password-manager/internal/storage"
)

type Manager struct {
	storage *storage.Storage
}

func NewManager(s *storage.Storage) *Manager {
	return &Manager{storage: s}
}

func (m *Manager) SavePassword() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	err := m.storage.SavePassword(name, password)
	if err != nil {
		fmt.Println("Error saving password:", err)
	} else {
		fmt.Println("Password saved successfully!")
	}
}

func (m *Manager) GetPassword() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	passwords, err := m.storage.LoadPasswords()
	if err != nil {
		fmt.Println("Error loading passwords:", err)
		return
	}

	password, exists := passwords[name]
	if exists {
		fmt.Println("Password:", password)
	} else {
		fmt.Println("Password not found.")
	}
}

func (m *Manager) ListPasswordNames() {
	passwords, err := m.storage.LoadPasswords()
	if err != nil {
		fmt.Println("Error loading passwords:", err)
		return
	}

	fmt.Println("Saved passwords:")
	for name := range passwords {
		fmt.Println("-", name)
	}
}