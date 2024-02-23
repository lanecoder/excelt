package main

import (
	"context"
	"encoding/json"
	"excelt/conf"
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// type Builder struct {
// 	field []reflect.StructField
// }

// func (b *Builder) NewBuilder() *Builder {
// 	return &Builder{}
// }

// func (b *Builder) AddField(field string, typ reflect.Type) {
// 	b.field = append(b.field, reflect.StructField{Name: field, Type: typ})
// }

// // 根据预先添加的字段构建出结构体
// func (b *Builder) Build() *Struct {
// 	stu := reflect.StructOf(b.field)
// 	index := make(map[string]int)
// 	for i := 0; i < stu.NumField(); i++ {
// 		index[stu.Field(i).Name] = i
// 	}
// 	return &Struct{stu, index}
// }

// func (b *Builder) AddString(name string) {
// 	b.AddField(name, reflect.TypeOf(""))
// }

// func (b *Builder) AddBool(name string) {
// 	b.AddField(name, reflect.TypeOf(true))
// }

// func (b *Builder) AddInt8(name string) {
// 	b.AddField(name, reflect.TypeOf(int8(0)))
// }

// func (b *Builder) AddInt32(name string) {
// 	b.AddField(name, reflect.TypeOf(int32(0)))
// }

// func (b *Builder) AddInt64(name string) {
// 	b.AddField(name, reflect.TypeOf(int64(0)))
// }

// func (b *Builder) AddInt(name string) {
// 	b.AddField(name, reflect.TypeOf(int(0)))
// }

// func (b *Builder) AddFloat64(name string) {
// 	b.AddField(name, reflect.TypeOf(float64(1.2)))
// }

// func (b *Builder) AddSliceStr(name string) {
// 	b.AddField(name, reflect.TypeOf([]string{}))
// }

// func (b *Builder) AddSliceInt8(name string) {
// 	b.AddField(name, reflect.TypeOf([]int8{}))
// }

// func (b *Builder) AddSliceInt32(name string) {
// 	b.AddField(name, reflect.TypeOf([]int32{}))
// }

// func (b *Builder) AddSliceInt64(name string) {
// 	b.AddField(name, reflect.TypeOf([]int64{}))
// }

// func (b *Builder) AddSliceInt(name string) {
// 	b.AddField(name, reflect.TypeOf([]int{}))
// }

// // 实际生成的结构体，基类
// // 结构体的类型
// type Struct struct {
// 	typ reflect.Type
// 	// <fieldName : 索引> // 用于通过字段名称，从Builder的[]reflect.StructField中获取reflect.StructField
// 	index map[string]int
// }

// func (s Struct) New() *Instance {
// 	return &Instance{reflect.New(s.typ).Elem(), s.index}
// }

// // 结构体的值
// type Instance struct {
// 	instance reflect.Value
// 	// <fieldName : 索引>
// 	index map[string]int
// }

// var (
// 	FieldNoExist error = errors.New("field no exist")
// )

// func (in Instance) Field(name string) (reflect.Value, error) {
// 	if i, ok := in.index[name]; ok {
// 		return in.instance.Field(i), nil
// 	} else {
// 		return reflect.Value{}, FieldNoExist
// 	}
// }
// func (in *Instance) SetString(name, value string) {
// 	if i, ok := in.index[name]; ok {
// 		in.instance.Field(i).SetString(value)
// 	}
// }

// func (in *Instance) SetBool(name string, value bool) {
// 	if i, ok := in.index[name]; ok {
// 		in.instance.Field(i).SetBool(value)
// 	}
// }

// func (in *Instance) SetInt64(name string, value int64) {
// 	if i, ok := in.index[name]; ok {
// 		in.instance.Field(i).SetInt(value)
// 	}
// }

// func (in *Instance) SetFloat64(name string, value float64) {
// 	if i, ok := in.index[name]; ok {
// 		in.instance.Field(i).SetFloat(value)
// 	}
// }
// func (i *Instance) Interface() interface{} {
// 	return i.instance.Interface()
// }

// func (i *Instance) Addr() interface{} {
// 	return i.instance.Addr().Interface()
// }

func main() {
	// ParseStr()
	for filename, sname := range File2StructMap {
		file, fileErr := excelize.OpenFile(filename)
		if fileErr != nil {
			fmt.Printf("open file %s err %s\n", filename, fileErr.Error())
			continue
		}
		sheetName := File2SheetMap[filename]
		// 获取行
		rows, rowErr := file.GetRows(sheetName)
		if rowErr != nil {
			fmt.Printf("get rows err %s\n", rowErr.Error())
			continue
		}
		fmt.Printf("get row %v\n", rows)
		if len(rows) <= 0 {
			fmt.Printf("excel row count error\n")
			continue
		}
		if ftyp, ok := ConfMap[sname]; ok {
			DynamicRefreshDb(sname, ftyp, rows)
		} else {
			fmt.Printf("DynamicRefreshDb error %s\n", sname)
		}
		// ptr := ConfMap[sname]
		// if sname == "Foo" {
		// 	fooPtr := (*conf.Foo)(ptr)
		// 	ftyp := reflect.TypeOf(*fooPtr)
		// 	DynamicRefreshDb(sname, ftyp, rows)
		// } else if sname == "GameMode" {
		// 	gameModePtr := (*conf.GameMode)(ptr)
		// 	gtyp := reflect.TypeOf(*gameModePtr)
		// 	DynamicRefreshDb(sname, gtyp, rows)
		// }
		// f := &conf.Foo{}
		// if len(rows) > 1 { // 小于等于1的只有行头
		// 	ftyp := reflect.TypeOf(*f)
		// 	DynamicRefreshDb(ftyp, rows)
		// }
	}
	// file, fileErr := excelize.OpenFile("helloworld.xlsx")
	// if fileErr != nil {
	// 	fmt.Printf("open file err %s\n", fileErr.Error())
	// }
	// // opts := &excelize.Options{}
	// // 获取行
	// rows, rowErr := file.GetRows("helloworld")
	// if rowErr != nil {
	// 	fmt.Printf("get rows err %s\n", rowErr.Error())
	// }
	// fmt.Printf("get row %v\n", rows)
	// if len(rows) <= 0 {
	// 	fmt.Printf("excel row count error\n")
	// }
	// f := &conf.Foo{}
	// if len(rows) > 1 { // 小于等于1的只有行头
	// 	ftyp := reflect.TypeOf(*f)
	// 	DynamicRefreshDb(ftyp, rows)
	// }

}

func DynamicRefreshDb(structName string, ftyp reflect.Type, rows [][]string) {
	var interfaceVal []interface{}
	numField := ftyp.NumField()
	headerRow := rows[0]
	if numField <= len(headerRow) {
		var isStructValid bool = true
		for i := 0; i < numField; i++ {
			if headerRow[i] == ftyp.Field(i).Tag.Get("json") {
				continue
			} else {
				fmt.Printf("missing json tag label %s %s\n", headerRow[i], ftyp.Field(i).Name)
				isStructValid = false
			}
		}
		// fmt.Printf("is struct field valid %d %d %v\n", len(rows), numField, isStructValid)
		if isStructValid { // 只有所有字段都吻合了才进行下一步操作
			for i := 1; i < len(rows); i++ {
				r := rows[i]
				fv := reflect.New(ftyp).Elem()
				fmt.Printf("DynamicParseStructField row str %s\n", r)
				DynamicParseStructField(numField, ftyp, fv, r)
				var ret interface{} = DynamicConvertStructPointer(ftyp, fv)
				interfaceVal = append(interfaceVal, ret)
			}
		}
	}
	Refurbish(structName, interfaceVal)
}

// 遍历反射动态解析struct的field
// @arg numField: 结构体的字段数量
// @arg ftyp: 结构体的类型
// @arg fv: 结构体的值
// @arg rawStrVal: 行原生数据
func DynamicParseStructField(numField int, ftyp reflect.Type, fv reflect.Value, rawStrVal []string) {
	for j := 0; j < numField; j++ {
		fmt.Printf("fv value %d %v\n", numField, fv)
		fieldOfStruct := fv.Field(j)
		stringTyp := fieldOfStruct.Type().Kind().String()
		fmt.Printf("struct field kind: %d %s\n", j, stringTyp)
		// 判断struct中每一个字段的类型并解析
		if stringTyp == "string" {
			fieldOfStruct.SetString(rawStrVal[j])
			// reflect.Complex64
		} else if stringTyp == "int" || stringTyp == "int8" || stringTyp == "int32" || stringTyp == "int64" {
			var parseVal int64
			var parseErr error
			if stringTyp == "int" || stringTyp == "int64" {
				parseVal, parseErr = strconv.ParseInt(rawStrVal[j], 0, 64)
				if parseErr != nil {
					fmt.Printf("parse int or int64 err %s\n", parseErr.Error())
				} else {
					fieldOfStruct.SetInt(parseVal)
				}
			} else if stringTyp == "int32" {
				parseVal, parseErr = strconv.ParseInt(rawStrVal[j], 0, 32)
				if parseErr != nil {
					fmt.Printf("parse int32 err %s\n", parseErr.Error())
				} else {
					fieldOfStruct.SetInt(parseVal)
				}
			} else if stringTyp == "int8" {
				parseVal, parseErr = strconv.ParseInt(rawStrVal[j], 0, 8)
				if parseErr != nil {
					fmt.Printf("parse int8 err %s\n", parseErr.Error())
				} else {
					fieldOfStruct.SetInt(parseVal)
				}
			}
		} else if stringTyp == "float32" {
			parseVal, parseErr := strconv.ParseFloat(rawStrVal[j], 32)
			if parseErr != nil {
				fmt.Printf("parse float32 err %s\n", parseErr.Error())
			} else {
				fieldOfStruct.SetFloat(parseVal)
			}
		} else if stringTyp == "float64" {
			parseVal, parseErr := strconv.ParseFloat(rawStrVal[j], 64)
			if parseErr != nil {
				fmt.Printf("parse float64 err %s\n", parseErr.Error())
			} else {
				fieldOfStruct.SetFloat(parseVal)
			}
		} else if stringTyp == "slice" {
			// 判断slice的元素类型
			elemVal := fv.Field(j).Type().Elem()
			// fmt.Printf("DynamicParseStructField slice %s\n", elemVal.Name())
			// fmt.Printf("DynamicParseStructField slice %s elem type %s\n", ftyp.Field(j).Name, elemVal.Name())
			// if ftyp.Field(j).Name == "Difficults" && elemVal.Name() == "float64" {
			if elemVal.Name() == "float64" {
				float64Slice := []float64{}
				unmarshalErr := json.Unmarshal([]byte(rawStrVal[j]), &float64Slice)
				if unmarshalErr != nil {
					fmt.Printf("json unmarshal err %s %s\n", rawStrVal[j], unmarshalErr.Error())
					continue
				} else {
					fmt.Printf("json unmarshal success %v\n", float64Slice)
					fieldOfStruct.Set(reflect.ValueOf(float64Slice))
				}
			} else if elemVal.Name() == "int8" {
				int8Slice := []int8{}
				unmarshalErr := json.Unmarshal([]byte(rawStrVal[j]), &int8Slice)
				if unmarshalErr != nil {
					fmt.Printf("json unmarshal err %s\n", unmarshalErr.Error())
					continue
				} else {
					fmt.Printf("json unmarshal success %v\n", int8Slice)
					fieldOfStruct.Set(reflect.ValueOf(int8Slice))

				}
			} else if elemVal.Name() == "int32" {
				intSlice := []int32{}
				unmarshalErr := json.Unmarshal([]byte(rawStrVal[j]), &intSlice)
				if unmarshalErr != nil {
					fmt.Printf("json unmarshal err %s\n", unmarshalErr.Error())
					continue
				} else {
					fmt.Printf("json unmarshal success %v\n", intSlice)
					fieldOfStruct.Set(reflect.ValueOf(intSlice))

				}
			} else if elemVal.Name() == "int64" {
				int64Slice := []int64{}
				unmarshalErr := json.Unmarshal([]byte(rawStrVal[j]), &int64Slice)
				if unmarshalErr != nil {
					fmt.Printf("json unmarshal err %s\n", unmarshalErr.Error())
					continue
				} else {
					fmt.Printf("json unmarshal success %v\n", int64Slice)
					fieldOfStruct.Set(reflect.ValueOf(int64Slice))
				}
			} else if elemVal.Name() == "int" {
				intSlice := []int{}
				unmarshalErr := json.Unmarshal([]byte(rawStrVal[j]), &intSlice)
				if unmarshalErr != nil {
					fmt.Printf("json unmarshal err %s\n", unmarshalErr.Error())
					continue
				} else {
					fmt.Printf("json unmarshal success %v\n", intSlice)
					fieldOfStruct.Set(reflect.ValueOf(intSlice))
				}
			} else if elemVal.Name() == "string" {
				stringSlice := []string{}
				unmarshalErr := json.Unmarshal([]byte(rawStrVal[j]), &stringSlice)
				if unmarshalErr != nil {
					fmt.Printf("json unmarshal err %s\n", unmarshalErr.Error())
					continue
				} else {
					fmt.Printf("json unmarshal success %v\n", stringSlice)
					fieldOfStruct.Set(reflect.ValueOf(stringSlice))
				}
			} else {
				var unmarshalErr error
				for n, v := range EmbeddingStructSliceMap {
					if ftyp.Field(j).Name == n {
						// FIXME 这里的判断和下面的断言无可避免，拓展性不佳，假如后续新增结构切片，此处还需要增加判断逻辑
						// 有没有更通用的方式去实现呢
						if n == "ModeInfos" {
							s := v.([]*conf.ModeInfo)
							unmarshalErr = json.Unmarshal([]byte(rawStrVal[j]), &s)
							if unmarshalErr != nil {
								fmt.Printf("json unmarshal err %s\n", unmarshalErr.Error())
								continue
							} else {
								fmt.Printf("json unmarshal success %d %v\n", len(s), s)
								fieldOfStruct.Set(reflect.ValueOf(s))
								break
							}
						}
					}
				}
			} /*else if ftyp.Field(j).Name == "ModeInfos" {
				// slice elem val ModeInfo ptr
				fmt.Printf("slice elem val %s %s\n", elemVal.Elem().Name(), elemVal.Kind().String())
				modeInfoSlice := []*conf.ModeInfo{}
				unmarshalErr := json.Unmarshal([]byte(rawStrVal[j]), &modeInfoSlice)
				if unmarshalErr != nil {
					fmt.Printf("json unmarshal err %s\n", unmarshalErr.Error())
					continue
				} else {
					fmt.Printf("json unmarshal success %d %v\n", len(modeInfoSlice), modeInfoSlice)
					fieldOfStruct.Set(reflect.ValueOf(modeInfoSlice))
				}
			}*/
		} else if stringTyp == "ptr" { // 一个指针，一般是嵌套结构体的指针
			for embeddingStructName, embeddingReflectTyp := range EmbeddingStructMap {
				if embeddingReflectTyp == fieldOfStruct.Type() {
					fmt.Printf("embeddingStructName is %s %v\n", embeddingStructName, embeddingReflectTyp)
					// embeddingReflectVal := reflect.Zero(embeddingReflectTyp.Elem())
					embeddingReflectVal := reflect.New(embeddingReflectTyp.Elem())
					// 指针指向的值
					embeddingReflectVal1 := embeddingReflectVal.Elem()
					// 调用Elem函数时，若调用方是指针，则返回的是指针指向的值 | 若调用方是非指针，则返回的是reflect.Value本身
					// Value => Type => Elem
					fieldOfStructElem := fieldOfStruct.Type().Elem()
					// 解析string => []string 例如要将 string {"paths":[3.0,2.0,5.0], "factor":3.8} 解析成 []string
					rawValStrSlice, convErr := convertToStrSlice(rawStrVal[j])
					if convErr != nil {
						fmt.Printf("convertToStrSlice err %s\n", convErr.Error())
						continue
					} else {
						fmt.Printf("convertToStrSlice success %s %v %v\n", rawValStrSlice, fieldOfStructElem, fieldOfStruct)
						DynamicParseStructField(fieldOfStructElem.NumField(), embeddingReflectTyp, embeddingReflectVal1, rawValStrSlice)
						fieldOfStruct.Set(embeddingReflectVal)
						break
					}
				}
			}

			// 硬编码的方式
			// rawVal := &conf.ModeExtra{}
			// // Type *conf.ModeExtra
			// modeExtraPtrMetaType := reflect.TypeOf(rawVal)
			// // Value pointer of conf.ModeExtra, is a pointer
			// modeExtraPtrMetaVal := reflect.ValueOf(rawVal)
			// // conf.ModeExtra
			// modeExtraMetaVal := modeExtraPtrMetaVal.Elem()
			// // 调用Elem函数时，若调用方是指针，则返回的是指针指向的值 | 若调用方是非指针，则返回的是reflect.Value本身
			// // Value => Type => Elem
			// fieldOfStructElem := fieldOfStruct.Type().Elem()
			// fmt.Printf("reflect.Type compare %v %v\n", modeExtraPtrMetaType, fieldOfStruct.Type())
			// if modeExtraPtrMetaType == fieldOfStruct.Type() {
			// 	// 解析string => []string 例如要将 string {"paths":[3.0,2.0,5.0], "factor":3.8} 解析成 []string
			// 	rawValStrSlice, convErr := convertToStrSlice(rawStrVal[j])
			// 	if convErr != nil {
			// 		fmt.Printf("convertToStrSlice err %s\n", convErr.Error())
			// 	} else {
			// 		fmt.Printf("convertToStrSlice success %s %v %v\n", rawValStrSlice, fieldOfStructElem, fieldOfStruct)
			// 		DynamicParseStructField(fieldOfStructElem.NumField(), modeExtraPtrMetaType, modeExtraMetaVal, rawValStrSlice)
			// 		fieldOfStruct.Set(modeExtraPtrMetaVal)
			// 	}
			// }
		}
	}
}

// raw string convert to string slice
// @doc 这里已编码的切片的数据经map中间转换，顺序可能会产生变化,因此需要先依次提取出json key
func convertToStrSlice(jsonEncodedData string) ([]string, error) {
	re := regexp.MustCompile(`"([^"]+)":`)
	matches := re.FindAllStringSubmatch(jsonEncodedData, -1)
	// 提取匹配到的结果
	var keys []string = []string{}
	for _, match := range matches {
		// fmt.Printf("print matches %v\n", match)
		keys = append(keys, match[1])
	}
	if len(keys) <= 0 {
		return keys, nil
	}
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonEncodedData), &data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("aft regex match keys %v\n", keys)
	var ret []string
	for i := 0; i < len(keys); i++ {
		if val, ok := data[keys[i]]; ok {
			strValue, err := json.Marshal(val)
			if err != nil {
				return nil, err
			}
			ret = append(ret, string(strValue))
		}
	}
	return ret, nil
}

// FIXME 此处也需要硬编码转换为相应结构指针，拓展性不佳
func DynamicConvertStructPointer(ftyp reflect.Type, fv reflect.Value) interface{} {
	fmt.Printf("DynamicConvertStructPointer %v %v\n", ftyp.Name(), fv.Addr().Type().Name())
	if ftyp.Name() == "Foo" {
		fooPtr := (*conf.Foo)(fv.Addr().UnsafePointer())
		return fooPtr
	} else if ftyp.Name() == "GameMode" {
		modePtr := (*conf.GameMode)(fv.Addr().UnsafePointer())
		return modePtr
	} else {
		return nil
	}
}

func Refurbish(structName string, docs []interface{}) error {
	col, sess, err := GetSession(structName)

	if err != nil {
		fmt.Printf("db refurbish GetSession %s\n", err.Error())
		return err
	}
	defer sess.EndSession(context.TODO())
	sessCtx := mongo.NewSessionContext(context.TODO(), sess)
	if err = sess.StartTransaction(); err != nil {
		fmt.Printf("db refurbish StartTransaction %s\n", err.Error())
		return err
	}
	_, err = col.DeleteMany(sessCtx, bson.D{})
	if err != nil {
		fmt.Printf("db refurbish DeleteMany %s\n", err.Error())
		sess.AbortTransaction(context.Background())
		return err
	}
	docss := []interface{}{}
	docss = append(docss, docs...)
	_, err = col.InsertMany(sessCtx, docss)
	if err != nil {
		fmt.Printf("db refurbish InsertMany %s\n", err.Error())
		sess.AbortTransaction(context.Background())
		return err
	}
	return sess.CommitTransaction(context.Background())
}

func ParseStr() {
	p := `"{"paths":[3.0,2.0,5.0], "factor":3.8}"`
	// ps := strings.Split(p, ":")
	// for i := 0; i < len(ps); i++ {
	// 	fmt.Printf("parse str %d %s\n", i, ps[i])
	// }
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(p), &m)
	if err != nil {
		fmt.Printf("parse str error %s\n", err.Error())
	}
	for k, v := range m {
		fmt.Printf("parse str ret %v %v \n", k, v)
	}
}
