package models

type Navs struct {
	Id      int32  `gorm:"primary_key"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Status  int32  `json:"status"`
	AddTime int64  `json:"add_time"`
}

func (Navs) TableName() string {
	return "navs"
}
