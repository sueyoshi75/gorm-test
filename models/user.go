package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID uint
	Name string
	Email string
	EmailVerified_at *time.Time
	Password string
	RoleId uint
	RememberToken *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	ApiToken *string
}

func (u *User) Verify(tx *gorm.DB) error {
	password := u.Password
	if err := tx.First(&u).Error; err != nil {
		return err
	}
	hash := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

