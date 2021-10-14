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
	db := make(map[string]domain.Produce)
	db["A12T-4GH7-QPL9-3N4M"] = domain.Produce{
		Code:      "A12T-4GH7-QPL9-3N4M",
		Name:      "Lettuce",
		UnitPrice: 3.46,
	}
	db["E5T6-9UI3-TH15-QR88"] = domain.Produce{
		Code:      "E5T6-9UI3-TH15-QR88",
		Name:      "Peach",
		UnitPrice: 2.99,
	}
	db["YRT6-72AS-K736-L4AR"] = domain.Produce{
		Code:      "YRT6-72AS-K736-L4AR",
		Name:      "Green Pepper",
		UnitPrice: 0.79,
	}
	db["TQ4C-VV6T-75ZX-1RMR"] = domain.Produce{
		Code:      "TQ4C-VV6T-75ZX-1RMR",
		Name:      "Gala Apple",
		UnitPrice: 3.59,
	}

	return database{
		db:  db,
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
