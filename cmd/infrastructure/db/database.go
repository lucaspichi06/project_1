package db

import (
	"sync"

	errors2 "errors"

	"github.com/project_1/cmd/domain"
	"github.com/project_1/cmd/errors"
)

type database struct {
	db  map[string]domain.Produce
	mtx sync.Mutex
}

func NewDataBase() database {
	return database{
		db:  make(map[string]domain.Produce),
		mtx: sync.Mutex{},
	}
}

func (d database) Add(produce *domain.Produce) error {
	if _, ok := d.db[produce.Code]; ok {
		return errors.NewInternalServerAppError("the produce already exists", errors2.New("the produce already exists"))
	}

	d.mtx.Lock()
	defer d.mtx.Unlock()
	d.db[produce.Code] = *produce

	return nil
}

func (d database) Fetch() ([]domain.Produce, error) {
	result := []domain.Produce{}

	for _, v := range d.db {
		result = append(result, v)
	}

	return result, nil
}

func (d database) Delete(code string) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	delete(d.db, code)

	return nil
}
