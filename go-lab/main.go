package main

import (
	"encoding/json"
	"log"
)

type Perm struct {
	Id   string `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type Menu struct {
	Perm
	 string `json:"url,omitempty" bson:"url,omitempty"`
}

type Role struct {
	Perms []Menu `json:"perms,omitempty" bson:"perms,omitempty"`
}

func main() {
	m := new(Menu)

	m.Perm.Id = "123"
	m.Url = "456"

	p := new(Perm)

	p.Id = "p123"

	r := &Role{Perms: []Menu{*m}}

	s, _ := json.Marshal(r)
	log.Println(string(s))

}
