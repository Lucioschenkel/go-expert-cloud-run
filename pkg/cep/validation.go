package cep

import "regexp"

func Validate(cep string) (bool, error) {
	matched, err := regexp.Match(`\d{8}`, []byte(cep))
	if err != nil {
		return false, err
	}

	return matched, nil
}
