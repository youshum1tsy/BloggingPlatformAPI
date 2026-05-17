package main

import "time"

type Post struct {
	Id        uint64
	Title     string
	Content   string
	Tags      []string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
