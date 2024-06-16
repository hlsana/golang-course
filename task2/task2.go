package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

type User struct {
	ID string
}

func checkDuplicates(u []User, e User) bool {
	for _, a := range u {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	var users []User
	for i := 0; i < 20; i++ {
		users = append(users, User{
			ID: fmt.Sprintf("User #%d", rand.IntN(15)),
		})
	}

	fmt.Printf("ID before removing duplicates: %+v\n", users)
	for i := 0; i < len(users); i++ {
		if checkDuplicates(users[i+1:], users[i]) {
			users = remove(users, i)
			i--
		}
	}
	fmt.Printf("ID after removing duplicates: %+v\n", users)

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})
	fmt.Printf("ID after sorting: %+v\n", users)
}
func remove(slice []User, s int) []User {
	return append(slice[:s], slice[s+1:]...)
}
