package database

import "github.com/qiniu/qmgo"

func Users() *qmgo.Collection {
	client := Database()
	db := client.Database("class")
	coll := db.Collection("users")
	return coll
}