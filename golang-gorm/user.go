package golanggorm

import (
	"time"

	"gorm.io/gorm"
)
type User struct {
	ID string `gorm:"primary_key;column:id;<-:create"`
	Password string `gorm:"column:password"`
	Name Name `gorm:"embedded"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Infromation string `gorm:"-"`
	Wallet Wallet `gorm:"foreignKey:user_id;references:id"`
	Addresses []Address `gorm:"foreignKey:user_id;references:id"`
	LikeProducts []Product `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;references:id;jonReferences:product_id"`
}

// change default table 
func (u *User) TableName()string  {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) (err error)  {
	if u.ID == ""{
		u.ID = "user-" + time.Now().Format("20060102150405")
	}
	return nil
}

type Name struct {
	FirstName string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName string `gorm:"column:last_name"`
}


