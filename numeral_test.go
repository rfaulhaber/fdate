package fdate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumeralFromString(t *testing.T) {
	testCases := []struct{
		numeralStr string
		expectedInt int
	}{
		{
			numeralStr: "V",
			expectedInt: 5,
		},
		{
			numeralStr: "XXIX",
			expectedInt: 29,
		},
		{
			numeralStr: "IV",
			expectedInt: 4,
		},
	}

	for _, tc := range testCases {
		actual := NumeralFromString(tc.numeralStr)

		assert.Equal(t, tc.expectedInt, actual.number)
	}
}

func TestNumeralFromNumber(t *testing.T) {
	testCases := []struct{
		testInt int
		expectedStr string
	}{
		{
			testInt: 4,
			expectedStr: "IV",
		},
		{
			testInt:24,
			expectedStr:"XXIV",
		},
		{
			testInt:109,
			expectedStr:"CIX",
		},
		{
			testInt:225,
			expectedStr:"CCXXV",
		},
		{
			testInt:226,
			expectedStr:"CCXXVI",
		},
		{
			testInt:227,
			expectedStr:"CCXXVII",
		},
		{
			testInt:1,
			expectedStr:"I",
		},
		{
			testInt:11,
			expectedStr:"XI",
		},
	}

	for _, tc := range testCases {
		actual := NumeralFromNumber(tc.testInt)

		assert.Equal(t, tc.expectedStr, actual.String())
	}
}
