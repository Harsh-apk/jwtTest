package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type IncomingUser struct {
	UserName string `json:"userName,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"userName,omitempty" bson:"userName,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	EncPw    string             `json:"encPw,omitempty" bson:"encPw,omitempty"`
}
