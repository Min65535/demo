package model


//---------------------------base data--------------------------//
type NameAndValue struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Block  string `json:"block"`
	Remark string `json:"remark"`
}

func(n NameAndValue) TableName() string {
	return "case"
}
