package manager

import (
	"fmt"
	"password-manager/internal/storage"
)

type Manager struct {
	storage *storage.Storage
}

func NewManager(s *storage.Storage) *Manager {
	return &Manager{storage: s}
}

func (m *Manager) SavePassword(name, password string) error {
	return m.storage.SavePassword(name, password)
}

func (m *Manager) GetPassword(name string) {
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
