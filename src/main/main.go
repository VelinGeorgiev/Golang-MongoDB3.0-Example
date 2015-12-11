package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

//const MongoDb details
const (
	hosts      = "ds026491.mongolab.com:26491"
	database   = "messagingdb"
	username   = "admin"
	password   = "youPassword"
	collection = "messages"
)

func main() {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err1 := mgo.DialWithInfo(info)
	if err1 != nil {
		panic(err1)
	}

	col := session.DB(database).C(collection)

	count, err2 := col.Count()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(fmt.Sprintf("Messages count: %d", count))
}
