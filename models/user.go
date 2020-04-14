package models

import (
	"errors"
	"fmt"
)

//User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers fetched the users
func GetUsers() []*User {
	return users
}

//AddUser adds a user to the list of users
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not have an ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//GetUserByID get user by id
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("No user found with the id %v", id)
}

//UpdateUser update user
func UpdateUser(u User) (User, error) {
	for i, user := range users {
		if user.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("No user found with the id %v", u.ID)
}

//RemoveUserByID delete user by id
func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No user found with the id %v", id)
}
