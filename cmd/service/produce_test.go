package service

import (
	"errors"
	"testing"

	"github.com/project_1/cmd/domain"
	"github.com/strech/testify/assert"
)

type dummyDB struct {
	add    func(produce *domain.Produce) error
	fetch  func() ([]domain.Produce, error)
	delete func(code string) error
}

func (d dummyDB) Add(produce *domain.Produce) error {
	return d.add(produce)
}

func (d dummyDB) Fetch() ([]domain.Produce, error) {
	return d.fetch()
}

func (d dummyDB) Delete(code string) error {
	return d.delete(code)
}

func TestAddProduce(t *testing.T) {
	t.Run("add produce success", func(t *testing.T) {
		produce := &domain.Produce{
			Name:      "Lettuce",
			Code:      "A12T-4GH7-QPL9-3N4M",
			UnitPrice: 3.16,
		}

		db := &dummyDB{
			add: func(produce *domain.Produce) error {
				return nil
			},
		}

		service := NewProduceService(db)
		err := service.Add(produce)

		assert.NoError(t, err)
	})
	t.Run("add produce invalid code", func(t *testing.T) {
		produce := &domain.Produce{
			Name:      "Lettuce",
			Code:      "A12T-4GH7-QPL9",
			UnitPrice: 3.16,
		}

		db := &dummyDB{
			add: func(produce *domain.Produce) error {
				return nil
			},
		}

		service := NewProduceService(db)
		err := service.Add(produce)

		assert.Error(t, err)
		assert.Equal(t, "the provided data is invalid", err.(domain.AppError).ErrorMessage)
	})
	t.Run("add produce invalid name", func(t *testing.T) {
		produce := &domain.Produce{
			Name:      "Lettuce*/-&#%",
			Code:      "A12T-4GH7-QPL9-3N4M",
			UnitPrice: 3.16,
		}

		db := &dummyDB{
			add: func(produce *domain.Produce) error {
				return nil
			},
		}

		service := NewProduceService(db)
		err := service.Add(produce)

		assert.Error(t, err)
		assert.Equal(t, "the provided data is invalid", err.(domain.AppError).ErrorMessage)
	})
	t.Run("add produce invalid unit price", func(t *testing.T) {
		produce := &domain.Produce{
			Name:      "Lettuce",
			Code:      "A12T-4GH7-QPL9-3N4M",
			UnitPrice: 3.1624,
		}

		db := &dummyDB{
			add: func(produce *domain.Produce) error {
				return nil
			},
		}

		service := NewProduceService(db)
		err := service.Add(produce)

		assert.Error(t, err)
		assert.Equal(t, "the provided data is invalid", err.(domain.AppError).ErrorMessage)
	})
}
func TestFetchProduce(t *testing.T) {
	t.Run("fetch produce success", func(t *testing.T) {
		db := &dummyDB{
			fetch: func() ([]domain.Produce, error) {
				return []domain.Produce{
					{
						Name:      "Lettuce",
						Code:      "A12T-4GH7-QPL9-3N4M",
						UnitPrice: 3.16,
					},
					{
						Name:      "Peach",
						Code:      "E5T6-9UI3-TH15-QR88",
						UnitPrice: 2.99,
					},
				}, nil
			},
		}

		service := NewProduceService(db)
		produces, err := service.Fetch()

		assert.NoError(t, err)
		assert.Equal(t, 2, len(produces))
	})
	t.Run("fetch produce error", func(t *testing.T) {
		db := &dummyDB{
			fetch: func() ([]domain.Produce, error) {
				return []domain.Produce{}, errors.New("test-error")
			},
		}

		service := NewProduceService(db)
		produces, err := service.Fetch()

		assert.Error(t, err)
		assert.Equal(t, "test-error", err.Error())
		assert.Equal(t, 0, len(produces))
	})
}
func TestDeleteProduce(t *testing.T) {
	t.Run("delete produce success", func(t *testing.T) {
		db := &dummyDB{
			delete: func(code string) error {
				return nil
			},
		}

		service := NewProduceService(db)
		err := service.Delete("A12T-4GH7-QPL9-3N4M")

		assert.NoError(t, err)
	})
	t.Run("delete produce error", func(t *testing.T) {
		db := &dummyDB{
			delete: func(code string) error {
				return errors.New("test-error")
			},
		}

		service := NewProduceService(db)
		err := service.Delete("A12T-4GH7-QPL9-3N4M")

		assert.Error(t, err)
		assert.Equal(t, "test-error", err.Error())
	})
}
