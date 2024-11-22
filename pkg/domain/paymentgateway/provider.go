package paymentgateway

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type Provider struct {
	uuid     string
	platform Platform
	country  []string
}

func NewProvider(platformName string) (*Provider, error) {
	// if uuid == "" {
	// 	return nil, errors.New("Empty provider uuid")
	// }
	if platformName == "" {
		return nil, errors.New("Empty provider platform")
	}

	_platform, err := NewPlatformFromString(strings.ToUpper(platformName))
	if err != nil {
		return nil, err
	}

	// if country == "" {
	// 	return nil, errors.New("Empty provider country code")
	// }

	// if !_platform.IsSupportedInCountry(country) {
	// 	return nil, errors.New("Payment provider not supported in specified country")
	// }

	return &Provider{
		uuid:     uuid.New().String(),
		platform: _platform,
		country:  _platform.GetSupportedCountries(platformName),
	}, nil
}

func UnmarshalProviderFromDatabase(id string, platformName string, countries []string) (*Provider, error) {
	_platform, err := NewPlatformFromString(platformName)
	if err != nil {
		return nil, err
	}
	return &Provider{
		uuid:     id,
		platform: _platform,
		country:  countries,
	}, nil
}

func (p Provider) UUID() string {
	return p.uuid
}

func (p Provider) Platform() Platform {
	return p.platform
}

func (p Provider) Country() []string {
	return p.country
}

func (p Provider) IsZero() bool {
	return len(p.country) == 0
}
