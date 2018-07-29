package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

//type Product struct {
//	ID    uint `gorm:"primary_key:id"`
//	Num   int  `gorm:"AUTO_INCREMENT:number"`
//	Code  string
//	Price uint
//	Tag   []Tag     `gorm:"many2many:tag;ForeignKey:ID;AssociationForeignKey:ID"`
//	Date  time.Time `gorm:"-"`
//}
//
//type Email struct {
//	ID         int    `gorm:"primary_key:id"`
//	UserID     int    `gorm:"not null;index"`
//	Email      string `gorm:"type:varchar(100);unique_index"`
//	Subscribed bool
//}

type CustomizePerson struct {
	IdPerson string             `gorm:"primary_key:true"`
	Accounts []CustomizeAccount `gorm:"column:accounts;ForeignKey:IdAccount"`
}

type CustomizeAccount struct {
	IdAccount string `gorm:"primary_key:true"`
	Name      string
	Date      time.Time `gorm:"-"`
}

type Tag struct {
	Name string
}

func main() {
	db, err := gorm.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	a := CustomizeAccount{IdAccount: "1", Name: "shangan", Date: time.Now()}
	b := CustomizeAccount{IdAccount: "2", Name: "sergey", Date: time.Now()}
	account := []CustomizeAccount{a, b}

	person := &CustomizePerson{"test", account}
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "demo_" + defaultTableName
	//}

	//db.AutoMigrate(&Product{}, &Email{})
	db.AutoMigrate(&CustomizePerson{}, &CustomizeAccount{})
	//db.NewRecord(person)
	if err := db.Create(person).Error; err != nil {
		panic(err)
	}

	result := CustomizePerson{}

	db.First(&result).Related(&result.Accounts, "IdAccount")
	fmt.Printf("%#v,%v", result, result.Accounts)
}
