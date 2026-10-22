package user

type User struct {
	ID             int    `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	Email          string `gorm:"not null"`
	PasswordHash   string
	AvatarFileName string
	PhoneNumber    string `gorm:"not null"`
	Role           string `gorm:"not null"`
}
