package db

import (
	"context"
	"fmt"

	"github.com/Harsh-apk/jwtTest/types"
	"github.com/Harsh-apk/jwtTest/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	LoginUser(*types.IncomingUser, context.Context) (*types.User, error)
	CreateUser(*types.User, context.Context) error
	GetUserById(*primitive.ObjectID, context.Context) (*types.User, error)
}
type MongoUserStore struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func NewMongoUserStore(Client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		Client: Client,
		Coll:   Client.Database(DBNAME).Collection(USERCOLL),
	}
}

func (n *MongoUserStore) LoginUser(inUser *types.IncomingUser, ctx context.Context) (*types.User, error) {
	var user types.User
	err := n.Coll.FindOne(ctx, bson.D{{Key: "email", Value: inUser.Email}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	if utils.ComparePassword(&inUser.Password, &user.EncPw) {
		return &user, nil
	}
	return nil, fmt.Errorf("bad credentials")

}

func (n *MongoUserStore) CreateUser(User *types.User, ctx context.Context) error {
	res, err := n.Coll.InsertOne(ctx, User)
	if err != nil {
		return err
	}
	User.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (n *MongoUserStore) GetUserById(id *primitive.ObjectID, ctx context.Context) (*types.User, error) {
	var user types.User
	err := n.Coll.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
