package model

import (
	"testing"
	"reflect"
)

func TestNewUser(t *testing.T) {
	us1 := NewUser(100)
	
	if !reflect.DeepEqual(*us1, User{1, 100}) {
		t.Error("Unequal")
	}
	
	us2 := NewUser(120)
		if !reflect.DeepEqual(*us2, User{2, 120}) {
		t.Error("Unequal")
	}
}