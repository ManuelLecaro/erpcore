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
	Category    []*Category
	Images      []Image
	CreatedAt   time.Time `bson:"created_at"`
	Name        string
	EAN         string
	Description string
	ID          uint64 `bson:"id"`
}

type Image struct {
	Name        string
	Description string
	URL         string
	File        io.ReadSeeker
	ID          uint64
}

type Category struct {
	Children    []*Category
	CreatedAt   time.Time `bson:"created_at"`
	ID          string    `bson:"id,omitempty"`
	Description string
	Path        string
	Name        string
}

type Transaction struct {
	ID        uint64 `bson:"id"`
	PDF       string
	Receiver  string
	Sender    string
	CreatedAt time.Time `bson:"created_at"`
}

type ErrorResponse struct {
	Code   uint64          `json:"code"`
	ErrMsg json.RawMessage `json:"content"`
}
