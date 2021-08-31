package store

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	client *mongo.Client
)

const (
	url      = ""
	database = "fastChat"
)

func InitMongoClient() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	client = conn

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(fmt.Errorf("Fatal error mongolib connect: %s \n", err))
	}
}

func GetDatabase() *mongo.Database {
	return client.Database(database)
}

func GetColl(coll string) *mongo.Collection {
	return client.Database(database).Collection(coll)
}

func InsertOne(coll string, m interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := GetColl(coll).InsertOne(ctx, m)
	if err != nil {
		fmt.Println("mongodb 添加数据异常", err)
	}
	return res, err
}

func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func FindOne(coll string, filter interface{}, info interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := GetColl(coll).FindOne(ctx, filter).Decode(info)
	return err
}

func FindAll(coll string, filter interface{}) (*mongo.Cursor, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	find, err := GetColl(coll).Find(ctx, filter)
	return find, err
}
