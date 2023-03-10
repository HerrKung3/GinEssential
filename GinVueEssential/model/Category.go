package model

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"` // 默认就是自增
	Name      string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"` //自动生成
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"` //自动生成

}
