package model

// Book struct represents books table in darabase
type Book struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
