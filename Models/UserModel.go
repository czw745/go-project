package Models

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (b *User) TableName() string {
	return "user"
}
