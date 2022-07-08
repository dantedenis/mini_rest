package model

import (
	"errors"
)

var id_counter int

type User struct {
	id int		
	balance int
	name string
}

func NewUser(name string, bal int) *User {
	id_counter++
	return &User{balance: bal, id: id_counter, name: name}
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
