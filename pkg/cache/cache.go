package cache

import (
	"sync"
	"wb_test/pkg/model"
	"errors"
)

type Cache struct {
	sync.RWMutex
	data map[int]*model.User
}

func NewCache() *Cache {
	return &Cache{
		data: map[int]*model.User{},
	}
}

func (c *Cache) Add(user *model.User) {
	c.Lock()
	defer c.Unlock()
	
	c.data[user.GetID()] = user
}

func (c *Cache) Update(id, amount int) error {
	c.RLock()
	v, ok := c.data[id]
	if !ok {
		return errors.New("ID doesn't exist")
	}
	c.RUnlock()
	
	c.Lock()
	defer c.Unlock()
	v.AddAmount(amount)
	return nil
}

func (c *Cache) GetUser(id int) *model.User {
	c.RLock()
	defer c.RUnlock()
	
	return c.data[id]
}