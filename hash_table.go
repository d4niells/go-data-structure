package main

import (
	"sync"
)

type HashTable struct {
	items map[int][]User
	lock  sync.RWMutex
}

func hash(name string) (key int) {
	for _, char := range name {
		key = 31*key + int(char)
	}
	return
}

func (ht *HashTable) Put(user User) int {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	key := hash(user.Name)

	if ht.items == nil {
		ht.items = make(map[int][]User)
	}

	ht.items[key] = append(ht.items[key], user)

	return key
}

func (ht *HashTable) Remove(name string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	key := hash(name)
	delete(ht.items, key)
}

func (ht *HashTable) Get(key int) []User {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return ht.items[key]
}

func (ht *HashTable) Search(name string) []User {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	key := hash(name)
	return ht.items[key]
}
