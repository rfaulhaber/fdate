package fdate

import (
	"testing"
)

func TestDate_Date(t *testing.T) {
	testDate := Date{82061}

	expectedYear := 225
	expectedMonth := Prairial
	expectedDay := 8

	actualYear, actualMonth, actualDay := testDate.Date()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}

	if expectedMonth != actualMonth {
		t.Error("Expected:\t", expectedMonth, "\tActual:\t", actualMonth)
	}

	if expectedDay != actualDay {
		t.Error("Expected:\t", expectedDay, "\tActual:\t", actualDay)
	}
}

func TestDate_Date2(t *testing.T) {
	testDate := Date{0}

	expectedYear := 1
	expectedMonth := Vendémiaire
	expectedDay := 1

	actualYear, actualMonth, actualDay := testDate.Date()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}

	if expectedMonth != actualMonth {
		t.Error("Expected:\t", expectedMonth, "\tActual:\t", actualMonth)
	}

	if expectedDay != actualDay {
		t.Error("Expected:\t", expectedDay, "\tActual:\t", actualDay)
	}
}

func TestDate_Date3(t *testing.T) {
	testDate := Date{2603}

	// The 18th Brumaire, when Napoleon staged his famous coup, ending the Revolution
	expectedYear := 8
	expectedMonth := Brumaire
	expectedDay := 18

	actualYear, actualMonth, actualDay := testDate.Date()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}

	if expectedMonth != actualMonth {
		t.Error("Expected:\t", expectedMonth, "\tActual:\t", actualMonth)
	}

	if expectedDay != actualDay {
		t.Error("Expected:\t", expectedDay, "\tActual:\t", actualDay)
	}
}

func TestDate_Date4(t *testing.T) {
	testDate := Date{673}

	// The 9th Thermidor, when the Convention turned against Robespierre
	expectedYear := 2
	expectedMonth := Thermidor
	expectedDay := 9

	actualYear, actualMonth, actualDay := testDate.Date()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}

	if expectedMonth != actualMonth {
		t.Error("Expected:\t", expectedMonth, "\tActual:\t", actualMonth)
	}

	if expectedDay != actualDay {
		t.Error("Expected:\t", expectedDay, "\tActual:\t", actualDay)
	}
}

func TestDate_Date5(t *testing.T) {
	testDate := Date{81814}

	expectedYear := 225
	expectedMonth := Vendémiaire
	expectedDay := 1

	actualYear, actualMonth, actualDay := testDate.Date()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}

	if expectedMonth != actualMonth {
		t.Error("Expected:\t", expectedMonth, "\tActual:\t", actualMonth)
	}

	if expectedDay != actualDay {
		t.Error("Expected:\t", expectedDay, "\tActual:\t", actualDay)
	}
}

func TestDate_Date6(t *testing.T) {
	testDate := Date{81813}

	expectedYear := 224
	expectedMonth := Complémentaires
	expectedDay := 5

	actualYear, actualMonth, actualDay := testDate.Date()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}

	if expectedMonth != actualMonth {
		t.Error("Expected:\t", expectedMonth, "\tActual:\t", actualMonth)
	}

	if expectedDay != actualDay {
		t.Error("Expected:\t", expectedDay, "\tActual:\t", actualDay)
	}
}

func TestDate_Year(t *testing.T) {
	testDate := Date{673}

	expectedYear := 2

	actualYear := testDate.Year()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}
}

func TestDate_Month(t *testing.T) {
	testDate := Date{2603}

	expectedMonth := Brumaire

	actualMonth := testDate.Month()

	if expectedMonth != actualMonth {
		t.Error("Expected:\t", expectedMonth, "\tActual:\t", actualMonth)
	}
}

func TestDate_NewDate(t *testing.T) {
	expectedDate := Date{82061}

	testYear := 225
	testMonth := Prairial
	testDay := 8

	actualDate := NewDate(testYear, testMonth, testDay)

	if expectedDate.days != actualDate.days {
		t.Error("Expected:\t", expectedDate.days, "\tActual:\t", actualDate.days)
	}
}

func TestDate_NewDate2(t *testing.T) {
	expectedDate := Date{0}

	testYear := 1
	testMonth := Vendémiaire
	testDay := 1

	actualDate := NewDate(testYear, testMonth, testDay)

	if expectedDate.days != actualDate.days {
		t.Error("Expected:\t", expectedDate.days, "\tActual:\t", actualDate.days)
	}
}

func TestDate_NewDate3(t *testing.T) {
	expectedDate := Date{2603}

	// The 18th Brumaire, when Napoleon staged his famous coup, ending the Revolution
	testYear := 8
	testMonth := Brumaire
	testDay := 18

	actualDate := NewDate(testYear, testMonth, testDay)

	if expectedDate.days != actualDate.days {
		t.Error("Expected:\t", expectedDate.days, "\tActual:\t", actualDate.days)
	}
}

func TestDate_NewDate4(t *testing.T) {
	expectedDate := Date{673}

	// The 9th Thermidor, when the Convention turned against Robespierre
	testYear := 2
	testMonth := Thermidor
	testDay := 9

	actualDate := NewDate(testYear, testMonth, testDay)

	if expectedDate.days != actualDate.days {
		t.Error("Expected:\t", expectedDate.days, "\tActual:\t", actualDate.days)
	}
}

func TestDate_NewDate5(t *testing.T) {
	expectedDate := Date{81814}

	testYear := 225
	testMonth := Vendémiaire
	testDay := 1

	actualDate := NewDate(testYear, testMonth, testDay)

	if expectedDate.days != actualDate.days {
		t.Error("Expected:\t", expectedDate.days, "\tActual:\t", actualDate.days)
	}
}

func TestDate_NewDate6(t *testing.T) {
	expectedDate := Date{81813}

	testYear := 224
	testMonth := Complémentaires
	testDay := 5

	actualDate := NewDate(testYear, testMonth, testDay)

	if expectedDate.days != actualDate.days {
		t.Error("Expected:\t", expectedDate.days, "\tActual:\t", actualDate.days)
	}
}

func TestDate_NewDate7(t *testing.T) {
	expectedDate := Date{81810}

	testYear := 224
	testMonth := Complémentaires
	testDay := 2

	actualDate := NewDate(testYear, testMonth, testDay)

	if expectedDate.days != actualDate.days {
		t.Error("Expected:\t", expectedDate.days, "\tActual:\t", actualDate.days)
	}
}

func TestDate_Weekday(t *testing.T) {
	testDate := Date{82073}

	expectedWeekday := décadi
	actualWeekday := testDate.Weekday()

	if expectedWeekday != actualWeekday {
		t.Error("Expected:\t", expectedWeekday, "\tActual:\t", actualWeekday)
	}
}

func TestDate_Weekday2(t *testing.T) {
	testDate := Date{82074}

	expectedWeekday := primidi
	actualWeekday := testDate.Weekday()

	if expectedWeekday != actualWeekday {
		t.Error("Expected:\t", expectedWeekday, "\tActual:\t", actualWeekday)
	}
}

func TestDate_String(t *testing.T) {
	testDate := Date{82074}

	expected := "21 Prairial CCXXV"
	actual := testDate.String()

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual)
	}
}

func TestToday(t *testing.T) {
	// this is for 10 June 2017 and would need to be updated to pass
	expected := Date{82078}
	actual := Today()

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual)
	}
}

func TestDate_IsLeapYear(t *testing.T) {
	testDate := Date{4386}

	actual := testDate.IsLeapYear()
	expected := true

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual)
	}
}

func TestDate_IsLeapYear2(t *testing.T) {
	testDate := Date{4020}

	actual := testDate.IsLeapYear()
	expected := false

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual)
	}
}
