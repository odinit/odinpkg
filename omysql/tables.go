package omysql

type Temp struct {
	Id   int    `gorm:"column:id;type:int(11);not null;primary_key" json:"-"`
	Name string `gorm:"column:name;type:varchar(255);default null;unique" json:"name"`
}
