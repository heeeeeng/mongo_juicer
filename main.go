package main

import (
"fmt"
"github.com/globalsign/mgo"
//"github.com/globalsign/mgo/bson"
"time"
)

const (
	MongoUrl string = "root:123456@172.28.152.101:27017"
	Db string = "uni"
	Collection string = "proxy"
)

func main() {
	session, err := mgo.Dial(MongoUrl)
	if err != nil {
		fmt.Println(fmt.Sprintf("cannot dial mongodb: %s, err: %v", MongoUrl, err))
		return
	}

	c := session.DB(Db).C(Collection)

	type mongoData struct {
		InsertTime string `bson:"insert_time"`
	}
	doc := mongoData{
		time.Now().Format("2006-01-02T15:04:05.999999-07:00"),
	}
	//data, _ := bson.Marshal(doc)

	err = c.Insert(doc)
	if err != nil {
		fmt.Println("insert error: ", err)
		return
	}

	fmt.Println("finish insert")
}


