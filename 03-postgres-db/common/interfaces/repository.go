package interfaces

type Repository[T any] interface {
	// Create a new record in the database and return the ID of the new record
	Create(data *T) (T, error)
	GetByID(id uint64) error
	Update(data *T) error
	Delete(id uint64) error
}
