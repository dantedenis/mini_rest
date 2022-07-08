package cache

import (
	"testing"
	"wb_test/pkg/model"
	_"fmt"
)

// run with -race

func TestCache(t *testing.T) {
	c := NewCache()
	
	for i:=0; i < 5; i++ {
		go c.Add(model.NewUser(100))
	}
	
	
}