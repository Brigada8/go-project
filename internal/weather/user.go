package weather

type User struct {
	Id       uint   `json:"id"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password []byte `gorm:"type:varchar(100);not null"`
}
