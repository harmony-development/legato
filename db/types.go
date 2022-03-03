package db

import "time"

type User struct {
	Id       int64
	Username string
	Created  time.Time
}

type UserAccount struct {
	User
	Email         string
	Password_Hash []byte
}
