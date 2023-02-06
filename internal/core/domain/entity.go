package domain

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	ID        uint64 `bson:"id"`
	Name      string
	Email     string
	URI       string
	CreatedAt time.Time `bson:"created_at"`
}

type Article struct {
	ID          uint64 `bson:"id"`
	Name        string
	EAN         string
	Description string
	Images      []Image
	Category    []*Category
	CreatedAt   time.Time `bson:"created_at"`
}

type Image struct {
	ID          uint64
	Name        string
	Description string
	URL         string
	File        io.ReadSeeker
}

type Category struct {
	ID          string `bson:"id,omitempty"`
	Description string
	Path        string
	Name        string
	Children    []*Category
	CreatedAt   time.Time `bson:"created_at"`
}

type Transaction struct {
	ID        uint64 `bson:"id"`
	PDF       string
	User      User
	CreatedAt time.Time `bson:"created_at"`
}

type ErrorResponse struct {
	Code   uint64          `json:"code"`
	ErrMsg json.RawMessage `json:"content"`
}
