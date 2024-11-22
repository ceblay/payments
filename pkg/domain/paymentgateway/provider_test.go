package paymentgateway_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pg "github.com/ceblay/payments/pkg/domain/paymentgateway"
)

func TestNewProvider(t *testing.T) {
	t.Parallel()

	platformName := "PESAPAL"

	_provider, err := pg.NewProvider(
		platformName,
	)

	require.NoError(t, err)

	assert.Equal(t, platformName, (_provider.Platform().String()))
	assert.NotNil(t, _provider.UUID())
}

func TestNewProvider_platformdiffcases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name         string
		PlatformName string
		Id           string
		Expected     string
	}{
		{
			Name:         "lower_case",
			PlatformName: "pesapal",
			Expected:     "PESAPAL",
		},
		{
			Name:         "mixed_case",
			PlatformName: "pesApaL",
			Expected:     "PESAPAL",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.Name, func(t *testing.T) {
			_provider, err := pg.NewProvider(testCase.PlatformName)
			require.NoError(t, err)

			assert.Equal(t, testCase.Expected, _provider.Platform().String())
		})
	}
}

func TestNewProvider_invalid(t *testing.T) {
	t.Parallel()

	platformName := "WRONGNAME"

	_, err := pg.NewProvider("")
	assert.Error(t, err)

	// _, err = pg.NewProvider(providerUUID, "")
	// assert.Error(t, err)

	// _, err = pg.NewProvider("", platformName)
	// assert.Error(t, err)

	_, err = pg.NewProvider(platformName)
	assert.Error(t, err)
}
