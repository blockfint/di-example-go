package controller

type dbRepository[M interface{}] interface {
	List() (*[]M, error)
	FindByID(id uint) (*M, error)
	Create(m *M) error
}
