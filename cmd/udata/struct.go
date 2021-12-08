package udata

import "json"

type User struct {
    Name string `json:name`
    Age uint `json:age` // come on.
    Inventory []Item  `json:inv`
}

type Item struct {
    Name string `json:name,default:"undefined"`
}

func MakeUser(name string, age uint) *User {
    return &User{ name, age, make([]Item, 10)}
}

func (user *User) AddItem() {

}

