package core

import (
	"sort"
	"time"
)

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
		DateFrom:    time.Now(),
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

// Delete a item by it's id from the fake database
// returns (true) if item existed
func DeleteItemByID(itemID int64) bool {
	if _, exists := fakeItemDB[itemID]; exists {
		delete(fakeItemDB, itemID)
		return true
	}
	return false
}
