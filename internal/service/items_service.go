package service

import (
	"github.com/sirupsen/logrus"
	"go-base-service/internal/model"
	mongodb2 "go-base-service/internal/repository/mongodb"
	"go-base-service/pkg/config"
	"sync"
)

var items = make(map[string]model.Item)
var (
	itemService *ItemService
	once        sync.Once
)

type ItemService struct {
	dao *mongodb2.ItemDAO
}

func NewItemService() *ItemService {
	once.Do(func() {
		// Initialize MongoDB connection inside ItemService initialization
		mongoURI := config.AppConfig.GetString("persistence.mongodb.uri")
		mongodb2.ConnectMongoDB(mongoURI)
		itemDAO := mongodb2.NewItemDAO()

		itemService = &ItemService{
			dao: itemDAO,
		}
	})
	return itemService
}

func init() {
	var item2 = model.Item{
		ID:   "123",
		Name: "Example Item",
	}
	items[item2.ID] = item2
}

// FindItem finds an item by ID using the ItemDAO.
func (s *ItemService) FindItem(id string) (model.Item, bool) {
	item, err := s.dao.FindByID(id)
	if err != nil {
		logrus.Errorf("Failed to find item: %v", err)
		logrus.Errorf("Failed to find item: %v")
		return model.Item{}, false
	}
	return *item, true
}

// AddItem adds a new item to the database using ItemDAO.
func (s *ItemService) AddItem(item model.Item) model.Item {
	_, err := s.dao.Insert(item)
	if err != nil {
		// Log the error
		logrus.Errorf("Failed to insert item: %v", err)
		// Return an empty Item if there's an error
		return model.Item{}
	}
	logrus.Infof("Item Inserted Successfully: %v", item)
	return item
}

// RemoveItem removes an item from the database by ID using ItemDAO.
func (s *ItemService) RemoveItem(id string) bool {
	//deleteResult, err := s.dao.DeleteByID(id)
	_, err := s.dao.DeleteByID(id)
	if err != nil {
		return false
	}
	// You can use deleteResult if needed, or just return true if no error
	return true
}

// ANOTHER WAY OF DOING THE SINGLETON
//func init() {
//	once.Do(func() {
//		// Initialize MongoDB connection inside init
//		mongoURI := config.AppConfig.GetString("persistence.mongodb.uri")
//		mongodb.ConnectMongoDB(mongoURI)
//
//		itemDAO := mongodb.NewItemDAO()
//
//		itemService = &ItemService{
//			dao: itemDAO,
//		}
//	})
//}
//
//// NewItemService would simply return the initialized service
//func NewItemService() *ItemService {
//	return itemService
//}
