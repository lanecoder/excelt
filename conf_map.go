package main

import (
	"excelt/conf"
	"reflect"
)

const (
	// top level data struct name
	HELLO_WORLD_STRUCT = "Foo"

	GAME_MODE_STRUCT = "GameMode"

	// config file name
	HELLO_WORLD_FILE = "helloworld.xlsx"

	GAME_MODE_FILE = "game_mode.xlsx"

	// excel sheet name
	HELLO_WORLD_SHEET_NAME = "helloworld"

	GAME_MODE_SHEET_NAME = "game_mode"

	// name of collection that excel data insert into in mongodb
	HELLO_WORLD_COL_NAME = "foo"

	GAME_MODE_COL_NAME = "gamemode"

	// embedding struct （嵌套结构体）
	EMBEDDING_STRUCT_MODE_EXTRA = "ModeExtra"

	// 结构体切片自定义元素结构
	EMBEDDING_STRUCT_SLICE_MODE_INFOS = "ModeInfos"
)

var (
	// config file name => game logic struct name
	File2StructMap map[string]string = make(map[string]string)

	// config file name => excel sheet name
	File2SheetMap map[string]string = make(map[string]string)

	// game logic struct name => pointer to struct
	// ConfMap map[string]unsafe.Pointer = make(map[string]unsafe.Pointer)
	ConfMap map[string]reflect.Type = make(map[string]reflect.Type)

	// struct name => mongo db collection name
	DBColMap map[string]string = make(map[string]string)

	// embedding struct name => reflect type of a pointer to struct
	EmbeddingStructMap map[string]reflect.Type = make(map[string]reflect.Type)

	// name of field of embedding struct in slice => type 'any' assert from slice type
	EmbeddingStructSliceMap map[string]any = make(map[string]any)
)

func init() {
	File2StructMap[HELLO_WORLD_FILE] = HELLO_WORLD_STRUCT

	File2SheetMap[HELLO_WORLD_FILE] = HELLO_WORLD_SHEET_NAME

	// ConfMap[HELLO_WORLD_STRUCT] = unsafe.Pointer(new(conf.Foo))
	ConfMap[HELLO_WORLD_STRUCT] = reflect.TypeOf(*(new(conf.Foo)))

	DBColMap[HELLO_WORLD_STRUCT] = HELLO_WORLD_COL_NAME

	File2StructMap[GAME_MODE_FILE] = GAME_MODE_STRUCT

	File2SheetMap[GAME_MODE_FILE] = GAME_MODE_SHEET_NAME

	// ConfMap[GAME_MODE_STRUCT] = unsafe.Pointer(new(conf.GameMode))
	ConfMap[GAME_MODE_STRUCT] = reflect.TypeOf(*(new(conf.GameMode)))

	DBColMap[GAME_MODE_STRUCT] = GAME_MODE_COL_NAME

	EmbeddingStructMap[EMBEDDING_STRUCT_MODE_EXTRA] = reflect.TypeOf((new(conf.ModeExtra)))

	// 嵌套结构体切片
	structTyp := reflect.TypeOf(conf.ModeInfo{})
	structPtrSliceTyp := reflect.SliceOf(reflect.PtrTo(structTyp))
	// FIXME 反射创建的切片会自动扩容吗
	refStructPtrSlice := reflect.MakeSlice(structPtrSliceTyp, 0, 10).Interface()
	EmbeddingStructSliceMap[EMBEDDING_STRUCT_SLICE_MODE_INFOS] = refStructPtrSlice
}
