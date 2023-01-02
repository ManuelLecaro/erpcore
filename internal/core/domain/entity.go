package domain

import (
	"time"
)

type User struct {
	Name      string
	Email     string
	URI       string
	CreatedAt time.Time `bson:"created_at"`
}

type Article struct {
	Name        string
	EAN         string
	Description string
	Images      []string
	Category    Category
	CreatedAt   time.Time `bson:"created_at"`
}

type Category struct {
	Name        string
	Description string
	Category    *Category
	CreatedAt   time.Time `bson:"created_at"`
}

type Transaction struct {
	ID        string
	PDF       string
	User      User
	CreatedAt time.Time `bson:"created_at"`
}
