package repositories

import "github.com/project_1/cmd/domain"

type Produce interface {
	Add(produce *domain.Produce) error
	Fetch() ([]domain.Produce, error)
	Delete(code string) error
}