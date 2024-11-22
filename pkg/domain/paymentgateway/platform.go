package paymentgateway

import (
	"errors"
	"slices"
)

var ErrUnsupportedPaymentPlatform = errors.New("Unsupported payment platform")

const (
	Pesapal  = "PESAPAL"
	Paydunya = "PAYDUNYA"
)

var supportedCountries map[string][]string = map[string][]string{
	Pesapal:  {"KE", "TZ", "MW", "RW", "UG", "ZA", "ZM"},
	Paydunya: {"BF", "BJ", "CI", "ML", "SN", "TG"},
}

type Platform struct {
	ref string
}

func (p Platform) String() string {
	return p.ref
}

func (p Platform) IsZero() bool {
	return p == Platform{}
}

func (p Platform) IsSupportedInCountry(countryCodes []string) bool {
	status := true
	for _, countryCode := range countryCodes {
		if !slices.Contains(supportedCountries[p.ref], countryCode) {
			status = false
			break
		}
	}
	return status
}

func (p Platform) GetSupportedCountries(platformName string) []string {
	return supportedCountries[platformName]
}

func NewPlatformFromString(s string) (Platform, error) {
	switch s {
	case Pesapal:
		return Platform{Pesapal}, nil
	case Paydunya:
		return Platform{Paydunya}, nil
	}

	return Platform{}, ErrUnsupportedPaymentPlatform
}
