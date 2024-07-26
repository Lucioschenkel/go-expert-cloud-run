package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CelsiusToFarenheitFixture struct {
	Celsius   float64
	Farenheit float64
}

func TestCelsiusToFarenheit(t *testing.T) {
	fixtures := []CelsiusToFarenheitFixture{
		{
			1.0,
			33.8,
		},
		{
			39.0,
			102.2,
		},
		{
			15.0,
			59.0,
		},
	}

	for _, fixture := range fixtures {
		f := CelsiusToFarenheit(fixture.Celsius)

		assert.Equal(t, fixture.Farenheit, f)
	}
}

type CelsiusToKelvinFixture struct {
	Celsius float64
	Kelvin  float64
}

func TestCelsiusToKelvin(t *testing.T) {
	fixtures := []CelsiusToKelvinFixture{
		{
			1,
			274,
		},
		{
			10,
			283,
		},
		{
			0,
			273,
		},
		{
			50,
			323,
		},
	}

	for _, fixture := range fixtures {
		k := CelsiusToKelvin(fixture.Celsius)

		assert.Equal(t, fixture.Kelvin, k)
	}
}
