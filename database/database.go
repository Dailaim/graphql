package database

import (
	"context"

	"github.com/qiniu/qmgo"
)

func Database() *qmgo.Client{
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:3060"})
	if (err != nil ){
		panic(err)
	}
	return client
}