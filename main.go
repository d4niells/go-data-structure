package main

import "fmt"

type User struct {
	Name     string
	Lastname string
	Age      int
	Gender   string
}

func main() {
	users := []User{
		{"Artur", "Oliveira", 16, "M"},
		{"Daniel", "Oliveira", 20, "M"},
		{"Mariana", "Oliveira", 18, "F"},
		{"Mariana", "Silva", 19, "F"},
	}

	hashtable := HashTable{}

	keys := make([]int, len(users))

	for i, user := range users {
		keys[i] = hashtable.Put(user)
	}

	for _, key := range keys {
		item := hashtable.Get(key)

		for _, user := range item {
			fmt.Println(user.Name, user.Lastname)
		}
	}

	// Remove user "Daniel" from hash table
	hashtable.Remove("Daniel")
	fmt.Println(hashtable.items)

	// Search user by name
	item := hashtable.Search("Mariana")
	fmt.Println(item)
}
