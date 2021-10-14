package utils

import "regexp"

// ValidateCode validates the code to match with the especifications:
// - The produce codes are sixteen characters long, with dashes separating each four character group
// - The produce codes are alphanumeric and case insensitive
func ValidateCode(code string) (bool, error) {
	reg, err := regexp.Compile(`^([A-Za-z0-9]{4})-([A-Za-z0-9]{4})-([A-Za-z0-9]{4})-([A-Za-z0-9]{4})$`)
	if err != nil {
		return false, err
	}

	data := reg.FindAllStringSubmatch(code, -1)
	return len(data) == 1, nil
}

// ValidateName validates the name to match with the especifications:
// - The produce name is alphanumeric and case insensitive
func ValidateName(name string) (bool, error) {
	reg, err := regexp.Compile(`^[a-zA-Z0-9\s]+$`)
	if err != nil {
		return false, err
	}

	data := reg.FindAllStringSubmatch(name, -1)
	return len(data) == 1, nil
}
