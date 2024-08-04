package storage

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Storage struct {
	filepath string
}

func NewStorage(filepath string) *Storage {
	return &Storage{filepath: filepath}
}

func (s *Storage) SavePassword(name, password string) error {
	file, err := os.OpenFile(s.filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s:%s\n", name, password))
	return err
}

func (s *Storage) LoadPasswords() (map[string]string, error) {
	file, err := os.Open(s.filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	passwords := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			passwords[parts[0]] = parts[1]
		}
	}

	return passwords, scanner.Err()
}