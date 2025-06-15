package models

type ProblemCategory struct {
	ID uint `gorm:"primarykey;" json:"id"`
	BaceModel
	ProblemId  uint      `gorm:"column:problem_id;type:int(11);" json:"problem_id"`     // 问题的ID
	CategoryId uint      `gorm:"column:category_id;type:int(11);" json:"category_id"`   // 分类的ID
	Category   *Category `gorm:"foreignKey:id;references:category_id;" json:"category"` // 关联分类的基础信息表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
