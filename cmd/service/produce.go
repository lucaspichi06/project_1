package service

import (
	"github.com/project_1/cmd/domain"
	"github.com/project_1/cmd/domain/repositories"
	"github.com/project_1/cmd/errors"
)

type produce struct {
	db repositories.Produce
}

func NewProduceService(db repositories.Produce) produce {
	return produce{
		db: db,
	}
}

func (p *produce) Add(produce *domain.Produce) error {
	isValid, err := produce.IsValid()
	if err != nil {
		return errors.NewInternalServerAppError("error while validating the data", err)
	}

	if !isValid {
		return errors.NewBadRequestAppError("the provided data is invalid")
	}

	return p.db.Add(produce)
}

func (p *produce) Fetch() ([]domain.Produce, error) {
	return p.db.Fetch()
}

func (p *produce) Delete(code string) error {
	return p.db.Delete(code)
}
