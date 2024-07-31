package service

import (
	"go-base-service/model"
)

var items = make(map[string]model.Item)

func init() {
	var item2 = model.Item{
		ID:   "123",
		Name: "Example Item",
	}
	items[item2.ID] = item2
}

// FindItem finds an item by ID.
func FindItem(id string) (model.Item, bool) {
	item, exists := items[id]
	return item, exists
}

// AddItem adds a new item to the store.
func AddItem(item model.Item) model.Item {
	items[item.ID] = item
	return item
}

// RemoveItem removes an item from the store by ID.
func RemoveItem(id string) bool {
	if _, exists := items[id]; exists {
		delete(items, id)
		return true
	}
	return false
}
