package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/andreylm/vm_test/pkg/models"
)

// Cache - cache
type Cache struct {
	mu    sync.Mutex
	store map[int]*models.User
}

// CreateCache - creates cache
func CreateCache() *Cache {
	return &Cache{store: make(map[int]*models.User)}
}

// Load - loads data from jsonFile
func (c *Cache) Load(filePath string) error {
	jsonFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return err
	}

	users := []*models.User{}

	if err := json.Unmarshal(jsonFile, &users); err != nil {
		log.Println(err)
	}

	for _, u := range users {
		c.Add(u)
	}

	return nil
}

// Add - adds user
func (c *Cache) Add(u *models.User) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[u.ID] = u
	return nil
}

// GetAll gets all user
func (c *Cache) GetAll() ([]*models.User, error) {
	users := make([]*models.User, 0, len(c.store))
	for _, u := range c.store {
		users = append(users, u)
	}
	return users, nil
}

// Get - gets user by id
func (c *Cache) Get(ID int) (*models.User, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.store[ID], nil
}

// Close - closes
func (c *Cache) Close() {

}
