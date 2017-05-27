package fdate

import "testing"

func TestDate_Date(t *testing.T) {
	testDate := Date{82061}

	expectedYear := 225
	expectedMonth := Prairial
	expectedDay := 8

	actualYear, actualMonth, actualDay := testDate.Date()

	if expectedYear != actualYear {
		t.Error("Expected:", expectedYear, "Actual:", actualYear)
	}

	if  expectedMonth != actualMonth{
		t.Error("Expected:", expectedMonth, "Actual:", actualMonth)
	}

	if  expectedDay != actualDay{
		t.Error("Expected:", expectedDay, "Actual:", actualDay)
	}
}

func TestDate_Date2(t *testing.T) {
	testDate := Date{0}

	expectedYear := 1
	expectedMonth := Vendémiaire
	expectedDay := 1

	actualYear, actualMonth, actualDay := testDate.Date()

	if expectedYear != actualYear {
		t.Error("Expected:", expectedYear, "Actual:", actualYear)
	}

	if  expectedMonth != actualMonth{
		t.Error("Expected:", expectedMonth, "Actual:", actualMonth)
	}

	if  expectedDay != actualDay{
		t.Error("Expected:", expectedDay, "Actual:", actualDay)
	}
}

func TestDate_Date3(t *testing.T) {
	testDate := Date{2603}

	expectedYear := 8
	expectedMonth := Brumaire
	expectedDay := 18

	actualYear, actualMonth, actualDay := testDate.Date()

	if expectedYear != actualYear {
		t.Error("Expected:", expectedYear, "Actual:", actualYear)
	}

	if  expectedMonth != actualMonth{
		t.Error("Expected:", expectedMonth, "Actual:", actualMonth)
	}

	if  expectedDay != actualDay{
		t.Error("Expected:", expectedDay, "Actual:", actualDay)
	}
}
