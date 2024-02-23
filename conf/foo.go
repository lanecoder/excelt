package conf

type Foo struct {
	Name   string `json:"name" bson:"name"`
	Age    int8   `json:"age" bson:"age"`
	Height int    `json:"height" bson:"height"`
	Money  int32  `json:"money" bson:"money"`
	Fav    int64  `json:"fav" bson:"fav"`
	// Holiday map[string]bool `excel:"假期,expand:regexp(^\\d{4}-\\d{2}-\\d{2}$)" json:"holiday"`
}
