package cep

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Fixture struct {
	Cep   string
	Valid bool
}

func TestValidate(t *testing.T) {
	fixtures := []Fixture{
		{
			"11111111",
			true,
		},
		{
			"91110150",
			true,
		},
		{
			"9A782750",
			false,
		},
		{
			"1111111",
			false,
		},
	}

	for _, fixture := range fixtures {
		valid, err := Validate(fixture.Cep)

		assert.Equal(t, fixture.Valid, valid)
		assert.Nil(t, err)
	}
}
