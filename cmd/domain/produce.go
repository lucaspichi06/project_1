package domain

import (
	"errors"
	"math"

	"github.com/project_1/cmd/utils"
)

type Produce struct {
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	UnitPrice float64 `json:"unit_price"`
}

// IsValid validates the produce data
func (p *Produce) IsValid() (bool, error) {
	isValid, err := utils.ValidateCode(p.Code)
	if err != nil {
		return false, errors.New("error while validation the code data")
	}

	if !isValid {
		return isValid, nil
	}

	isValid, err = utils.ValidateName(p.Name)
	if err != nil {
		return false, errors.New("error while validation the name data")
	}

	if p.UnitPrice != math.Round(p.UnitPrice*100)/100 {
		isValid = false
	}

	return isValid, nil
}
