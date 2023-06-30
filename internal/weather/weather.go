package weather

type History struct {
	HistoryId uint   `json:"id"`
	UserID    uint   `gorm:"foreignkey:Id"`
	Location  string `gorm:"type:varchar(100);not null"`
	City      string `gorm:"type:varchar(100);not null"`
}
