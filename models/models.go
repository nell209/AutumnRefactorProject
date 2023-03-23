package models

// TemporaryAccount TODO potentially add an expiry -> also normal gorm types
type TemporaryAccount struct {
	ID                string `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID            string
	User              User
	TemporaryPassword string
	Authenticated     bool `gorm:"default:false"`
}
