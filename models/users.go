package models

// 结构体名称首字母必须大写 并和数据库表名称对应
type Users struct {
	Id       int32
	Username string
	Age      int16
	Email    string
	AddTime  int64
}

// 配置操作数据库表名称
func (Users) TableName() string {
	return "users"
}
