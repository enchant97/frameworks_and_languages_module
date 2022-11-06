package core

/*
This contains the fake items's database
*/

// The fake database next index to use,
// really just a fake sql auto increment
var nextItemID int64 = 1

// The fake database,
// just a map with a int for the key and Item's for the value
var fakeItemDB map[int64]Item = make(map[int64]Item)
