package cache

import (
	"sync"
	"wb_test/pkg/model"
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

func (c *Cache) Update(send *model.User, rec *model.User, amount int) error {	
	c.Lock()
	defer c.Unlock()
	
	err := send.TransferTo(rec, amount)
	return err
}

func (c *Cache) GetUser(id int) *model.User {
	c.RLock()
	defer c.RUnlock()
	
	return c.data[id]
}
