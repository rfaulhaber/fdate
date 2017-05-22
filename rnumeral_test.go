package fdate

import "testing"

func TestNumeralFromString(t *testing.T) {
	actual := NumeralFromString("V")

	if actual.number != 5 {
		t.Error("Expected 5, got ", actual.number)
	}
}

func TestNumeralFromNumber(t *testing.T) {

}

func TestRomanNumeral_Equal(t *testing.T) {

}

func TestRomanNumeral_Number(t *testing.T) {

}

func TestRomanNumeral_String(t *testing.T) {

}
