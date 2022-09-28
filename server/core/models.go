package core

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

// Check if in range of item's lat and lon
func (i *Item) InRange(radius float64, lat float64, lon float64) bool {
	if i.Lat != nil && i.Lon != nil {
		if (*i.Lat > lat-radius) &&
			(*i.Lat < lat+radius) &&
			(*i.Lon > lon-radius) &&
			(*i.Lon < lon+radius) {
			return true
		}
	}
	return false
}

// Struct used to pass filters to a crud method
type ItemsFilter struct {
	UserID      *string
	CSVKeywords *string
	Lat         *float64
	Lon         *float64
	Radius      *float64
	DateFrom    *PythonISOTime
	DateTo      *PythonISOTime
}

type ItemsFilterFromRequest struct {
	UserID      *string  `form:"user_id"`
	CSVKeywords *string  `form:"keywords"`
	Lat         *float64 `form:"lat"`
	Lon         *float64 `form:"lon"`
	Radius      *float64 `form:"radius"`
	DateFrom    *string  `form:"date_from"`
	DateTo      *string  `form:"date_to"`
}
