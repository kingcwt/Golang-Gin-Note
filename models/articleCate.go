package models

// foreignKey 外键 如果是表名加上Id默认也可以不配置
// references 主键 默认就是id 如果是id的话可以不配置
type ArticleCate struct {
	Id    int
	Title string
	State int
	// Article []Article `gorm:"foreignKey:CateId;references:Id"`
	Article []Article `gorm:"foreignKey:CateId"`
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
