package auth

import (
	"github.com/ilmu-merah/be-article/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.RepositoryInterface[repository.User]
}

func NewAuthService(repo repository.RepositoryInterface[repository.User]) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(newUser repository.User) error {
	// Enkripsi password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)

	// Simpan data user ke database via repository
	err = s.repo.Create(newUser)
	if err != nil {
		return err
	}

	return nil
}