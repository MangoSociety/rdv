package models

type User struct {
	ID       int64
	Email    string
	Phone    string
	PassHash []byte
}
