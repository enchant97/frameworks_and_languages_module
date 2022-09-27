package core

import (
	"fmt"
	"time"
)

type PythonISOTime time.Time

func (t PythonISOTime) MarshalJSON() ([]byte, error) {
	formattedTime := time.Time(t).Format("2006-01-02T15:04:05")
	return []byte(fmt.Sprintf("\"%s\"", formattedTime)), nil
}

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
	ID          int64          `json:"id"`
	UserID      string         `json:"user_id"`
	Keywords    []string       `json:"keywords"`
	Description string         `json:"description"`
	Image       *string        `json:"image"`
	Lat         *float64       `json:"lat"`
	Lon         *float64       `json:"lon"`
	DateFrom    PythonISOTime  `json:"date_from"`
	DateTo      *PythonISOTime `json:"date_to"`
}
