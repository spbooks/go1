package main

import "time"

type Image struct {
	ID          string
	UserID      string
	Name        string
	Location    string
	Size        int64
	CreatedAt   time.Time
	Description string
}
