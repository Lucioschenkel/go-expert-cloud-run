package usecase

import (
	"go-expert-cloud-run/pkg/weather"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ErrorFixture struct {
	Err error
	Cep string
}

func TestCurrentWeatherUseCase_Execute(t *testing.T) {
	allErrors := []ErrorFixture{
		{
			ErrCannotFindZipCode,
			"11111111",
		},
		{
			ErrInvalidZipCode,
			"9A178652",
		},
	}

	weatherApiClient := weather.NewWeatherClient(os.Getenv("API_KEY"))

	useCase := NewCurrentWeatherUseCase(*weatherApiClient)

	// Test error scenarios
	for _, fixture := range allErrors {
		res, err := useCase.Execute(Input{
			Cep: fixture.Cep,
		})

		assert.Nil(t, res)
		assert.Equal(t, err, fixture.Err)
	}

	// Test success case
	res, err := useCase.Execute(Input{
		Cep: "70165900",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.TempC)
	assert.NotEmpty(t, res.TempF)
	assert.NotEmpty(t, res.TempK)
}
