package db

import (
	//"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
)

const (
	USER string = "user"
	MSG  string = "msg"
)

var (
	session      *mgo.Session
	databaseName = "test"
)

func Session() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial("mongodb://127.0.0.1:27017")
		if err != nil {
			panic(err)
		}
	}
	return session.Clone()
}

func M(collection string, f func(*mgo.Collection)) {
	session := Session()
	defer func() {
		session.Close()
		if err := recover(); err != nil {
			log.Println("M", err)
		}
	}()
	c := session.DB(databaseName).C(collection)
	f(c)
}
