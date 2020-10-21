package Models

type Company struct {
	ID   uint   `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (b *Company) TableName() string {
	return "companys"
}
