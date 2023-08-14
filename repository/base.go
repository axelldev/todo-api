package repository

type Repository[T any] interface {
	FindAll() ([]T, error)
	FindById(id int) (*T, error)
	Save(any) error
	Update(any) error
	Delete(id int) error
}
