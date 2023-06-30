package weather

type History struct {
	UserID   uint   `gorm:"foreignkey:Id"`
	Location string `gorm:"type:varchar(100);not null"`
}
