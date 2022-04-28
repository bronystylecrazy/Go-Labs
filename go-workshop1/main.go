package main

import (
	"encoding/json"
	"fmt"
	c "strconv"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Downline []User `json:"downline"`
	Children []User `json:"children"`
}

func (u User) String(data ...any) string {
	return fmt.Sprintf("%s is %d years old", u.Name, u.Age)
}

func main() {
	var x int = 10
	var message string = "Hello, World!"
	message += " fuck you!"
	message += c.Itoa(x)

	user1 := User{
		Name: "John",
		Age:  30,
		Downline: []User{
			{
				Name: "Jack",
				Age:  20,
			},
		},
	}

	user2 := &user1

	user2.Name = "Tom"

	user1.Children = append(user1.Children, *(&user1))

	result, err := json.MarshalIndent(&user1, "", "  ")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(result))
}
