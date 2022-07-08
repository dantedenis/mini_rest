package cache

import (
	"testing"
	"wb_test/pkg/model"
	"sync"
	"math/rand"
	"time"
)


func TestCache(t *testing.T) {
	c := NewCache()
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())
	
	for i:=0; i < 10; i++ {
		wg.Add(1)
		go func(){
			wg.Done()
			c.Add(model.NewUser("test", 1000))
		}()
	}
	
	wg.Wait()
	
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			c.Update(c.GetUser(rand.Int()%10+1), c.GetUser(rand.Int()%10+1), 100)
		}()
	}
	wg.Wait()
	
	var sum int
	
	for i:=1; i < 11; i++ {
		sum += c.GetUser(i).GetBalance()
	}
	
	if sum != 10000 {
		t.Error("Excpected 10000, actual:", sum)
	}
}
