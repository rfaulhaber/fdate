package fdate

import "testing"

func TestNumeralFromString(t *testing.T) {
	actual := NumeralFromString("V")

	if actual.number != 5 {
		t.Error("Expected 5, got ", actual.number)
	}
}

func TestNumeralFromString2(t *testing.T) {
	actual := NumeralFromString("XXIX")

	expected := 29

	if actual.number != expected {
		t.Error("Expected ", expected, " got ", actual.number)
	}
}

func TestNumeralFromString3(t *testing.T) {
	actual := NumeralFromString("IV")

	expected := 4

	if actual.number != 4 {
		t.Error("Expected ", expected, " got ", actual.number)
	}
}

func TestNumeralFromNumber(t *testing.T) {
	actual := NumeralFromNumber(4)

	expected := "IV"

	if actual.numeral != expected {
		t.Error("Expected ", expected, " got ", actual.numeral)
	}
}

func TestNumeralFromNumber2(t *testing.T) {
	actual := NumeralFromNumber(24)

	expected := "XXIV"

	if actual.numeral != expected {
		t.Error("Expected ", expected, " got ", actual.numeral)
	}
}

func TestNumeralFromNumber3(t *testing.T) {
	actual := NumeralFromNumber(109)

	expected := "CIX"

	if actual.numeral != expected {
		t.Error("Expected ", expected, " got ", actual.numeral)
	}
}

func TestNumeralFromNumber4(t *testing.T) {
	actual := NumeralFromNumber(225)

	expected := "CCXXV"

	if actual.numeral != expected {
		t.Error("Expected ", expected, " got ", actual.numeral)
	}
}

func TestNumeralFromNumber5(t *testing.T) {
	actual := NumeralFromNumber(226)

	expected := "CCXXVI"

	if actual.numeral != expected {
		t.Error("Expected ", expected, " got ", actual.numeral)
	}
}
