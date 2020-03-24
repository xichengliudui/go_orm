package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type Operation interface {
	GetList()
	Insert()
	Save()
	Delete()
}

type User struct {

}

func (u User) Insert() {

}

func (u User) Save() {

}

func (u User) GetList() {

}

func (u User) Delete() {

}

func init() {

}

func main() {

}
