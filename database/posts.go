package database

import "github.com/qiniu/qmgo"

func Posts() *qmgo.Collection {

	client := Database()
	db := client.Database("class")
	coll := db.Collection("posts")
	return coll

}
