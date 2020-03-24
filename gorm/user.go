package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var mysql *gorm.DB

type Operation interface {
	Get() *User
	GetList() *[]User
	Insert()
	Save()
	Delete()
}

type User struct {
	ID       int64 `gorm: "primary_key"`
	Name     string
	Password string
}

func (u User) Insert() {
	mysql.Table("users").Create(&u)
}

func (u User) Save() {
	mysql.Table("users").Save(&u)
}

func (u User) Get() *User {
	if u.ID == 0 {
		// get the first user by id
		mysql.Table("users").First(&u)
	}
	// get the user by id
	mysql.Table("users").First(&u, u.ID)
	return &u
}

func (u User) GetList() *[]User {
	var users []User
	mysql.Table("users").Find(&users)
	return &users
}

func (u User) Delete() {
	mysql.Table("users").Unscoped().Delete(&u)
}

func init() {
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "127.0.0.1", 3306, "testdb")
	log.Printf("Database url is: %v\n", url)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Fatal()
	}
	// disable the plural form of the default table name
	db.SingularTable(true)
	mysql = db
}

func main() {
	/*
		// create user
		user := User{
			Name:     "kubernetes",
			Password: "kubernetesonceas",
		}
		user.Insert()
	*/

	/*
		// update user
		user := User{
			ID:		  2,
			Name:     "docker",
			Password: "dockeronceas",
		}
		user.Save()
	*/

	/*
		// delete user
		user := &User{
			ID:		  1,
		}
		if user.ID != 0 {
			user.Delete()
		}
	*/

	/*
		// get user
		var u  = &User{
			ID:       4,
		}
		u = u.Get()
		fmt.Printf("id is %d, name is %s, password is %v", u.ID, u.Name, u.Password)
	*/

	/*
		// get all user
		u := &User{}
		users := u.GetList()
		json, err := json.Marshal(users)
		if err != nil {
			log.Fatalf("Get users fatal %v", err)
		}
		fmt.Printf("Find result is %v", string(json))
	*/

}
