// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package city

import (
	"github.com/jittakal/go-apps/ams/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	host       = common.Host
	database   = common.Database
	collection = "city"
)

type City struct {
	Id      bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	StateId bson.ObjectId `json:"state_id"`
	Code    string        `json:"code"`
	Name    string        `json:"name"`
}

func (c City) Create() error {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	city := session.DB(database).C(collection)

	err = city.Insert(&c)
	return err
}

func (c City) Update() error {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	city := session.DB(database).C(collection)

	err = city.Update(bson.M{"_id": c.Id}, c)
	return err
}

func FindById(id string) (City, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	city := session.DB(database).C(collection)
	result := City{}

	err = city.FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

func FindByName(name string) ([]City, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	city := session.DB(database).C(collection)

	result := []City{}
	err = city.Find(bson.M{"name": name}).All(&result)
	return result, err
}

func All() ([]City, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	city := session.DB(database).C(collection)

	result := []City{}
	err = city.Find(bson.M{}).All(&result)
	return result, err
}
