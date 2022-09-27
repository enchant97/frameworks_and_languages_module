package core

import "time"

// The struct used when a user wants to add a new item
type ItemCreate struct {
	UserID      string   `json:"user_id" binding:"required"`
	Keywords    []string `json:"keywords" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Image       *string  `json:"image"`
	Lat         *float64 `json:"lat"`
	Lon         *float64 `json:"lon"`
}

// The struct used to represent a existing item
type Item struct {
	ID          int64      `json:"id"`
	UserID      string     `json:"user_id"`
	Keywords    []string   `json:"keywords"`
	Description string     `json:"description"`
	Image       *string    `json:"image"`
	Lat         *float64   `json:"lat"`
	Lon         *float64   `json:"lon"`
	DateFrom    time.Time  `json:"date_from"`
	DateTo      *time.Time `json:"date_to"`
}
