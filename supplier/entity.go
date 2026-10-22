package supplier

type Supplier struct {
	ID            int    `gorm:"primaryKey"`
	Name          string `gorm:"not null"`
	ContactPerson string `gorm:"not null"`
	Phone         string `gorm:"not null"`
	Address       string `gorm:"not null"`
}
