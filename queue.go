package main

import (
	"fmt"
	"sync"
)

type Node struct {
	User     *User
	NextUser *Node
}

func (n *Node) FreeNode() *User {
	user := n.User
	n.User = nil
	n.NextUser = nil

	return user
}

type UserManager struct {
	Head        *Node
	Tail        *Node
	MaxCapacity int
	UserCount   int
	Notifier    chan bool
	Mu          sync.Mutex
}

func NewUserRegister(queueSize int) *UserManager {
	return &UserManager{
		Head:        nil,
		Tail:        nil,
		MaxCapacity: queueSize,
		UserCount:   0,
		Notifier:    make(chan bool, queueSize),
		Mu:          sync.Mutex{},
	}
}

func (um *UserManager) IsFull() bool {
	return um.UserCount == um.MaxCapacity
}

func (um *UserManager) PrintStats() {
	fmt.Printf(
		"Queue stats:\n" +
		"Maximun Capacity -> %v | Current User Count -> %v\n",
		um.MaxCapacity, um.UserCount,
		)
}

func (um *UserManager) QueueUser(newUser *User) bool {
	um.Mu.Lock()
	defer func() {
		um.Notifier <- true
	}()
	defer um.Mu.Unlock()

	newNode := &Node{User: newUser}
	if um.IsFull() {
		println("Full queu")
		return false
	}

	um.UserCount++
	if um.Head == nil {
		um.Head = newNode
		um.Tail = newNode
		return true
	}

	um.Tail.NextUser = newNode
	um.Tail = newNode
	return true
}

func (um *UserManager) IsEmpty() bool {
	return um.UserCount == 0
}

func (um *UserManager) Dequeue() *User {
	um.Mu.Lock()
	defer um.Mu.Unlock()

	if um.IsEmpty() {
		return nil
	}

	um.UserCount--
	var firstUser *User
	if um.Head == um.Tail {
		firstUser = um.Head.FreeNode()
		um.Tail = nil
		um.Head = nil
	} else {
		firstNode := um.Head
		um.Head = firstNode.NextUser
		firstUser = firstNode.FreeNode()
	}

	return firstUser
}

func (um *UserManager) GetFirstUser() *User {
	if um.IsEmpty() {
		return nil
	}

	return um.Head.User
}

func (um *UserManager) GetLastUser() *User {
	if um.IsEmpty() {
		return nil
	}

	return um.Tail.User
}
