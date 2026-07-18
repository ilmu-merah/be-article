package repository

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id uint					`gorm:"primaryKey"`
	Fullname string					
	Username string
	Email string			`gorm:"not null"`
	Password string
	IsVerified bool			`gorm:"default:false"`
	EmailVerification EmailVerification	`gorm:"foreignKey:UserId"`
}

type StatusCode string

const (
	Pending StatusCode = "pending"
	Verified StatusCode = "verified"
	Revoked StatusCode = "revoked"
)

type EmailVerification struct {
	gorm.Model
	UserId uint			`gorm:"unique"`
	OtpCode uint		
	Status	StatusCode 	`gorm:"not null;default:'pending'"`
}

type PasswordReset struct {

}

var _ RepositoryInterface[User] = (*UserRepository[User])(nil)


type UserRepository[T any] struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository[User] {
	return &UserRepository[User]{db: db}
}

func (r UserRepository[T]) Create(data T) error {
	result := r.db.Create(&data)
	return result.Error
}

func (r UserRepository[T]) FindById(model T, id int) T {
	return model
}

func (r UserRepository[T]) UpdateById(data T, id int) bool {
	return true
}

func (r UserRepository[T]) DeleteById(model T, id int) bool {
	return  true
}