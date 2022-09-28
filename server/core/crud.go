package core

import (
	"sort"
	"strings"
	"time"
)

const defaultRadius = 5

// Create a new item in fake database
func CreateNewItem(item ItemCreate) Item {
	currentID := nextItemID
	createdItem := Item{
		ID:          currentID,
		UserID:      item.UserID,
		Keywords:    item.Keywords,
		Description: item.Description,
		Image:       item.Image,
		Lat:         item.Lat,
		Lon:         item.Lon,
		DateFrom:    PythonISOTime(time.Now()),
		DateTo:      nil,
	}
	fakeItemDB[currentID] = createdItem
	nextItemID++
	return createdItem
}

// Get a single item by it's id from the fake database,
// If item does not exist returns nil
func GetItemByID(itemID int64) *Item {
	if item, exists := fakeItemDB[itemID]; exists {
		return &item
	}
	return nil
}

// Get all items from database, sorted by id ascending
func GetItems() []Item {
	// Get the keys (item id's) from fake database
	keys := make([]int, len(fakeItemDB))
	i := 0
	for k := range fakeItemDB {
		keys[i] = int(k)
		i++
	}
	// Sort the key values (item id's)
	sort.Ints(keys)
	// Create a slice to contain the sorted items
	items := make([]Item, len(keys))
	for i, k := range keys {
		items[i] = fakeItemDB[int64(k)]
	}
	return items
}

// Get all items from database,
// filtering using given filters and sorted by id ascending
func GetItemsFiltered(filters ItemsFilter) []Item {
	items := GetItems()

	if filters == (ItemsFilter{}) {
		// no filter was provided, early return with all items
		return items
	}

	filteredItems := make([]Item, 0)

	if filters.Radius == nil {
		rad := float64(defaultRadius)
		filters.Radius = &rad
	}
	if filters.DateTo == nil {
		now := PythonISOTime(time.Now())
		filters.DateTo = &now
	}

	for _, item := range items {
		if filters.UserID != nil && item.UserID != *filters.UserID {
			// check if user does not match
			continue
		}
		if filters.CSVKeywords != nil && !containsAll(item.Keywords, strings.Split(*filters.CSVKeywords, ",")) {
			// check if keywords are missing
			continue
		}
		if filters.Lat != nil && filters.Lon != nil && !item.InRange(*filters.Radius, *filters.Lat, *filters.Lon) {
			// check if not in range
			continue
		}
		if filters.DateFrom != nil && time.Time(item.DateFrom).Before(time.Time(*filters.DateFrom)) {
			// check if date is less than item's date_from date
			continue
		}
		if item.DateTo != nil && time.Time(*item.DateTo).After(time.Time(*filters.DateTo)) {
			// check if date is after item's date_to date
			continue
		}
		filteredItems = append(filteredItems, item)
	}
	return filteredItems
}

// Delete a item by it's id from the fake database
// returns (true) if item existed
func DeleteItemByID(itemID int64) bool {
	if _, exists := fakeItemDB[itemID]; exists {
		delete(fakeItemDB, itemID)
		return true
	}
	return false
}
