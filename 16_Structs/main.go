package main

import (
	"fmt"
	"time"
)

type role struct {
	role   string
	roleId string
}

// user is a struct (similar to a class in OOP but simpler).
// It groups related fields together.
type user struct {
	name      string
	email     string
	role
	createdAt time.Time
}

// addUser is a constructor-like function.
func addUser(name string, email string) user {
	return user{
		name:      name,
		email:     email,
		createdAt: time.Now(),
	}
}

// addRole updates the user's role
func (u *user) addRole(roleName string) {
	u.role = role{
		role:   roleName,
		roleId: time.Now().Format("id"), // roleId from timestamp
	}
}

// changeName modifies the name
func (u *user) changeName(name string) {
	u.name = name
}

func main() {
	// Creating new users using addUser function
	user1 := addUser("Alok", "alok@gmail.com")
	user2 := addUser("Mia", "mia@gmail.com")

	fmt.Println("Before name change:", user2)

	// Modifying user2's name
	user2.changeName("Lol")
	fmt.Println("After name change:", user2)

	// Adding role
	user2.addRole("Admin")
	fmt.Println("After adding role:", user2)

	// user1 remains unchanged
	fmt.Println("User1:", user1)

	// Anonymous struct
	language := struct {
		language string
		value    int
	}{
		language: "Hindi",
		value:    1,
	}
	fmt.Println("Anonymous struct:", language)
}
