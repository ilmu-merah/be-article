package repository


type RepositoryInterface[T any] interface {
	Create(data T) error
	FindById(model T, id int) T
	UpdateById(data T, id int) bool
	DeleteById(model T, id int) bool
}