package core

import (
	"fmt"
	"time"
)

// The datetime format from Python's isoformat()
const PythonISOTimeFormat = "2006-01-02T15:04:05"

func containsAll(values []string, toFind []string) bool {
	valuesMap := make(map[interface{}]*interface{})
	for _, v := range values {
		valuesMap[v] = nil
	}

	for _, v := range toFind {
		if _, exists := valuesMap[v]; !exists {
			return false
		}
	}
	return true
}

type PythonISOTime time.Time

func (t PythonISOTime) MarshalJSON() ([]byte, error) {
	formattedTime := time.Time(t).Format(PythonISOTimeFormat)
	return []byte(fmt.Sprintf("\"%s\"", formattedTime)), nil
}

// Parse a string made using Python's isoformat() method
func ParsePythonISOTime(s string) (PythonISOTime, error) {
	parsedTime, err := time.Parse(PythonISOTimeFormat, s)
	if err != nil {
		return PythonISOTime{}, err
	}
	return PythonISOTime(parsedTime), nil
}
