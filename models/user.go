package models

type User struct {
	ID      int     `gorm:"primaryKey"`
	Name    string  `gorm:"not null"`
	Age     int     `gorm:"check:age >= 0"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Profile struct {
	ID                int `gorm:"primaryKey"`
	UserID            int `gorm:"uniqueIndex"`
	Bio               string
	ProfilePictureURL string
}
