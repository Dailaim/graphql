package resolver

import (
	"example/database"
	"example/generated"
	"example/model"

	"github.com/qiniu/qmgo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {

	UserMap map[string]*model.User
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (r *Resolver) Coll(Type string) *qmgo.Collection {

	if Type == "user" {
		return database.Users()
	} else if Type == "post" {
		return database.Posts()
	}
	return nil

}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
