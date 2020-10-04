package telephony

import (
	"github.com/alloyzeus/go-azcore/azcore"
)

// A CountryCode holds information about country code.
type CountryCode int32

// A NationalNumber holds information about national number.
type NationalNumber int64

// PhoneNumber holds a parsed phone number.
type PhoneNumber struct {
	countryCode    CountryCode
	nationalNumber NationalNumber
}

var _ azcore.ValueObject = PhoneNumber{}

// NewPhoneNumber returns a new instance of PhoneNumber with
// the provided details.
func NewPhoneNumber(
	countryCode CountryCode,
	nationalNumber NationalNumber,
) PhoneNumber {
	return PhoneNumber{
		countryCode:    countryCode,
		nationalNumber: nationalNumber,
	}
}

// CountryCode returns country code.
func (pn PhoneNumber) CountryCode() CountryCode { return pn.countryCode }

// WithCountryCode returns a new copy with provided country code.
func (pn PhoneNumber) WithCountryCode(countryCode CountryCode) PhoneNumber {
	return PhoneNumber{
		countryCode:    countryCode,
		nationalNumber: pn.nationalNumber,
	}
}

// NationalNumber returns national number.
func (pn PhoneNumber) NationalNumber() NationalNumber { return pn.nationalNumber }

// WithNationalNumber returns a new copy with provided national number.
func (pn PhoneNumber) WithNationalNumber(nationalNumber NationalNumber) PhoneNumber {
	return PhoneNumber{
		countryCode:    pn.countryCode,
		nationalNumber: nationalNumber,
	}
}

// Equals returns true if the other instance has the same values as this.
func (pn PhoneNumber) Equals(other interface{}) bool {
	if x, ok := other.(PhoneNumber); ok {
		return x.countryCode == pn.countryCode &&
			x.nationalNumber == pn.nationalNumber
	}
	if x, _ := other.(*PhoneNumber); x != nil {
		return x.countryCode == pn.countryCode &&
			x.nationalNumber == pn.nationalNumber
	}
	return false
}

// Equal is required for conformance with github.com/google/go-cmp package.
func (pn PhoneNumber) Equal(other interface{}) bool { return pn.Equals(other) }

// EqualsPhoneNumber returns true if the other instance has the same values as this.
func (pn PhoneNumber) EqualsPhoneNumber(other PhoneNumber) bool {
	return pn.countryCode == other.countryCode &&
		pn.nationalNumber == other.nationalNumber
}
