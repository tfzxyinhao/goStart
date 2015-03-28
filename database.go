package main

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/golang/glog"
	"github.com/golang/groupcache"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// all database connection
type Database struct {
	Mgo    *mgo.Session  "mongodb"
	Client *redis.Client "redis"
	Cache  *groupcache.Group
}

type MUser struct {
	id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Pwd  string        `bson:"pwd"`
}

type MProduct struct {
	id        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Price     float32       `bson:"price"`
	Timestamp time.Time
}

// init data connection
func (d *Database) Init() error {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		glog.Error(err)
		return err
	}

	u := MUser{}
	db := sess.DB("User")

	fmt.Println("------Find--------")
	iter := db.C("user").Find(nil).Limit(10).Sort("-name").Iter()
	for iter.Next(&u) {
		fmt.Println(u)
	}
	iter.Close()

	fmt.Println("------Find--------")
	db.C("user").Find(nil).Limit(10).One(&u)
	fmt.Println(u)

	fmt.Println("------Insert--------")
	fmt.Println(db.C("product").Insert(&MProduct{Name: "吹风机", Price: 5.25, Timestamp: time.Now()},
		&MProduct{Name: "洗衣机", Price: 2255.25, Timestamp: time.Now()}))

	fmt.Println("------Update--------")
	fmt.Println(db.C("product").Update(bson.M{"name": "吹风机"}, bson.M{"$set": bson.M{"price": 5.15}}))

	d.Mgo = sess
	//	client, err := redis.Dial("tcp", "localhost:6379")
	//	if err != nil {
	//		glog.Error(err)
	//		sess.Close()
	//		return err
	//	}
	//	d.Client = client
	return nil
}
