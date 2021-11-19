package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"url-shortener/db"
)

type URL struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	OriginalURL string             `json:"original_url" bson:"original_url"`
	ShortURL    string             `json:"short_url" bson:"short_url"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	ExpireAt    time.Time          `json:"expire_at" bson:"expire_at"`
}

const (
	DEFAULT_EXPIRE_TIME = 7
)

func GetOriginalURLFromHash(hashURL string) (*URL, error) {
	collection, dbCtx, cancel := db.GetUrlsCollection()
	defer cancel()

	result := collection.FindOne(dbCtx, bson.M{"short_url": hashURL})
	if result.Err() != nil {
		return nil, result.Err()
	}

	resultURL := &URL{}
	if err := result.Decode(resultURL); err != nil {
		return nil, err
	}

	return resultURL, nil
}

func (url *URL) Insert() (*URL, error) {
	collection, dbCtx, cancel := db.GetUrlsCollection()
	defer cancel()

	url.CreatedAt = time.Now()
	url.ExpireAt = time.Now().AddDate(0, 0, DEFAULT_EXPIRE_TIME)
	url.ID = primitive.NewObjectID()
	if _, err := collection.InsertOne(dbCtx, url); err != nil {
		return nil, err
	}

	return url, nil
}
