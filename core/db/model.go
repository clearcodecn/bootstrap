package db

type ExampleTable struct {
	Id int64 `json:"id" xorm:"id pk autoincr"`
}

func (ExampleTable) TableName() string {
	return "example_table"
}
