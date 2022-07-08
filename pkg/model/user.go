package model

import (
	"errors"
	"sync/atomic"
)

var id_counter uint32

type User struct {
	id int		
	balance int
	name string
}

func NewUser(name string, bal int) *User {
	atomic.AddUint32(&id_counter, 1)
	return &User{
		balance: bal, 
		id: int(atomic.LoadUint32(&id_counter)), 
		name: name,
	}
}


func (u *User) GetID() int {
	return u.id
}

func (u *User) AddAmount(amount int) {
	u.balance += amount
}

func (u *User) GetBalance() int {
	return u.balance
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) TransferTo(to *User, amount int) error {
	if u.balance - amount < 0 {
		return errors.New(u.name + ": not enough currency")
	}
	u.AddAmount(-amount)
	to.AddAmount(+amount)
	return nil
}
