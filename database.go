package main

import (
	"github.com/fzzy/radix/redis"
	"github.com/golang/glog"
	"github.com/golang/groupcache"
	"gopkg.in/mgo.v2"
)

type Database struct {
	Mgo    *mgo.Session  "mongodb"
	Client *redis.Client "redis"
	Cache  *groupcache.Group
}

func (d *Database) Init() error {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		glog.Error(err)
		return err
	}

	d.Mgo = sess
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		glog.Error(err)
		sess.Close()
		return err
	}
	d.Client = client
	return nil
}
