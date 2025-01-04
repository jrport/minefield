package main

import (
	"os"

	"golang.org/x/crypto/argon2"
)

var salt string
func init() {
	var ok bool
	salt, ok = os.LookupEnv("SALT")
	if !ok {
		println("Missing SALT, using 'DEFAULT'!")
	}
}

type User struct {
	Name string
	Id   string
}

func NewUser(name, ip string) *User {
	id := getAnomUserId(name, ip)

	return &User{
		Name: name,
		Id:   id,
	}
}

func getAnomUserId(name, ip string) string {
	userId := argon2.IDKey([]byte(ip + name), []byte(salt), 1, 64*1024, 4, 32)
	return string(userId)
}

