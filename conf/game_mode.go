package conf

type GameMode struct {
	ModeId     int32       `json:"mode_id" bson:"mode_id"`
	Desc       string      `json:"desc" bson:"desc"`
	Difficults []float64   `json:"difficults" bson:"difficults"`
	ModeInfos  []*ModeInfo `json:"mode_infos" bson:"mode_infos"`
	ModeExtra  *ModeExtra  `json:"mode_extra" bson:"mode_extra"`
}

type ModeInfo struct {
	ModeName string    `json:"mode_name" bson:"mode_name"`
	ModeKind int32     `json:"mode_kind" bson:"mode_kind"`
	Weights  []float64 `json:"weights" bson:"weights"`
}

type ModeExtra struct {
	Paths  []float64 `json:"paths" bson:"paths"`
	Factor float64   `json:"factor" bson:"factor"`
}
