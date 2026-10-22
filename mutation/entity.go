package mutation

import (
	"management-warehouse/product"
	"management-warehouse/user"
	"time"
)

type MutationType string

const (
	MutationTypeInbound  MutationType = "inbound"
	MutationTypeOutbound MutationType = "outbound"
)

type Mutation struct {
	ID        int
	ProductID int `gorm:"column:product_id" json:"product_id"`
	Product   product.Product
	Type      MutationType `gorm:"not null"`
	Quantity  int          `gorm:"not null"`
	Note      string       `gorm:"not null"`
	UserID    int          `gorm:"column:user_id" json:"user_id"`
	User      user.User    `gorm:"not null"`
	CreatedAt time.Time    `gorm:"not null"`
}
