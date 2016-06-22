package notify

import (
	"encoding/json"
	"io/ioutil"
)

const (
	// Email is id for email notifier
	Email = "email"
	// Stdout is id for stdout notifier
	Stdout = "stdout"
)

// Notifier is an interface grouping various notifiers (like eamil or stdout)
type Notifier interface {
	Notify(options *User) error
	Name() string
}

// User to notify
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Users is collection of User structs
type Users []*User

// UsersFromFile loads users from json file into Users
func UsersFromFile(file string) (Users, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var users Users
	if err := json.Unmarshal(content, &users); err != nil {
		return nil, err
	}
	return users, nil
}
