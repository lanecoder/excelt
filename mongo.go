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

	FooCol      *mongo.Collection            = nil
	GameModeCol *mongo.Collection            = nil
	DBColMap    map[string]*mongo.Collection = make(map[string]*mongo.Collection)
)

func init() {
	var err error = nil

	// @doc 数据库连接池默认是100
	// @doc connect=direct 直连主节点
	// @doc replicaSet 可替换成实际应用的数据集名称
	// @doc 192.168.2.64:17017 是主节点地址 可根据实际情况替换
	opt := options.Client().ApplyURI("mongodb://192.168.2.64:17017/?connect=direct&replicaSet=my-mongo-set")
	ClientObj, err = mongo.Connect(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	if err = ClientObj.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("连接mongodb失败")
		panic(err)
	}

	// 此示例中，数据库名是game_db。可根据相应情况进行替换
	DbObj = ClientObj.Database("game_db")

	// 此示例中两个集合名是foo gamemode。可根据相应情况进行替换
	for _, colName := range Struct2ColMap {
		col := DbObj.Collection(colName)
		DBColMap[colName] = col
	}
	// FooCol = DbObj.Collection("foo")
	// GameModeCol = DbObj.Collection("gamemode")

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
	// 根据structName获取对应的集合名
	for s, cn := range Struct2ColMap {
		if s == structName {
			// 根据集合名获取对应的集合对象
			col := DBColMap[cn]
			return col, sess, err
		}
	}
	// if structName == HELLO_WORLD_STRUCT {
	// 	return FooCol, sess, err
	// } else if structName == GAME_MODE_STRUCT {
	// 	return GameModeCol, sess, err
	// }
	return nil, nil, errors.New("get session error")
}
