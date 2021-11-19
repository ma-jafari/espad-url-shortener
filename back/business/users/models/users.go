package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"url-shortener/db"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	LastLogin time.Time          `json:"last_login" json:"last_login"`
	Token     string             `json:"token,omitempty" bson:"token,omitempty"`
}

func GetUserWithUserID(userID primitive.ObjectID) (user *User, err error) {
	collection, dbCtx, cancel := db.GetUsersCollection()
	defer cancel()

	result := collection.FindOne(dbCtx, bson.M{"_id": userID})
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err = result.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserWithEmail(email string) (*User, error) {
	collection, dbCtx, cancel := db.GetUsersCollection()
	defer cancel()

	result := collection.FindOne(dbCtx, bson.M{"email": email})
	if result.Err() != nil {
		return nil, result.Err()
	}

	user := &User{}
	if err := result.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) Insert() (*User, error) {
	collection, dbCtx, cancel := db.GetUsersCollection()
	defer cancel()

	user.CreatedAt = time.Now()
	user.LastLogin = time.Now()
	user.ID = primitive.NewObjectID()
	if _, err := collection.InsertOne(dbCtx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) Update() (*User, error) {
	collection, dbCtx, cancel := db.GetUsersCollection()
	defer cancel()

	if _, err := collection.ReplaceOne(dbCtx, bson.M{"_id": user.ID}, user); err != nil {
		return nil, err
	}

	return user, nil
}
