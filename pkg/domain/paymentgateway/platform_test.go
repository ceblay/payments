package paymentgateway_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pg "github.com/ceblay/payments/pkg/domain/paymentgateway"
)

func TestNewFromString(t *testing.T) {
	t.Parallel()

	platformName := "PESAPAL"
	platform, err := pg.NewPlatformFromString(platformName)

	require.NoError(t, err)

	assert.Equal(t, platformName, platform.String())
	assert.True(t, len(platform.GetSupportedCountries(platformName)) > 0)
}

func TestNewFromString_invalid(t *testing.T) {
	platformName := "BADCHOICE"
	_, err := pg.NewPlatformFromString(platformName)

	assert.Error(t, err)
}
