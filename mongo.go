package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// 数据库操作对象
var (
	ClientObj *mongo.Client   = nil
	DbObj     *mongo.Database = nil

	FooCol      *mongo.Collection = nil
	GameModeCol *mongo.Collection = nil
)

func init() {
	var err error = nil

	// 数据库连接池默认是100 connect=direct 直连主节点
	opt := options.Client().ApplyURI("mongodb://192.168.2.64:17017/?connect=direct&replicaSet=my-mongo-set")
	ClientObj, err = mongo.Connect(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	if err = ClientObj.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("连接mongodb失败")
		panic(err)
	}

	DbObj = ClientObj.Database("game_db")
	FooCol = DbObj.Collection("foo")
	GameModeCol = DbObj.Collection("gamemode")

	// DbObj = ClientObj.Database(conf.GameEnvConfig.Mongo.Db)
	fmt.Println("连接mongodb成功")
}

func GetClient() *mongo.Client {
	return ClientObj
}

func GetDb() *mongo.Database {
	return DbObj
}

func GetSession(structName string) (*mongo.Collection, mongo.Session, error) {
	sess, err := ClientObj.StartSession()
	if structName == HELLO_WORLD_STRUCT {
		return FooCol, sess, err
	} else if structName == GAME_MODE_STRUCT {
		return GameModeCol, sess, err
	}
	return nil, nil, errors.New("get session error")
}
