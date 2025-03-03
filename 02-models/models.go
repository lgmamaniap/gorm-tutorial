package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Declaring model example
type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivateAt   sql.NullTime
	ignored      string
}

// Embedded struct example
type Author struct {
	Name  string
	Email string
}

type Blog struct {
	Author
	ID      int
	Upvotes int32
}

// Equivalent to:
type EBlog struct {
	ID      int
	Name    string
	Email   string
	Upvotes int32
}

func main() {
	user := User{
		Name:         "Jinzhu",
		Email:        nil,
		Age:          18,
		Birthday:     &[]time.Time{time.Now()}[0],
		MemberNumber: sql.NullString{String: "201010", Valid: true},
		ActivateAt:   sql.NullTime{Time: time.Now(), Valid: true},
		ignored:      "Ignored",
	}

	fmt.Println(user)
}
