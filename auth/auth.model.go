package auth

import "time"

type User struct {
	Name         string    `bson: "name"`
	Email        string    `bson: "email"`
	PassWd       string    `bson: "password"`
	LastActivity time.Time `bson: "lastActivity"`
}
