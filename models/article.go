package models

type Article struct {
	Id     int
	Title  string
	CateId int //外键
	State  int
	// ArticleCa ArticleCate `gorm:"foreignKey:CateId"`
}

// 当你在关联一个表的时候 比方现在关联了一个ArticleCate表 默认会把ArticleCate的id作为主键 如果你在当前表中设置了ArticleCateId 就会把这个作为外键关联起来
func (Article) TableName() string {
	return "article"
}
