package fdate

import (
	"testing"
	"time"
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
	expectedDay := 6

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

func TestDate_Date7(t *testing.T) {
	testDate := Date{82179}

	expectedYear := 226
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

func TestDate_Year(t *testing.T) {
	testDate := Date{673}

	expectedYear := 2

	actualYear := testDate.Year()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}
}

func TestDate_Year2(t *testing.T) {
	testDate := Date{82179}

	expectedYear := 226

	actualYear := testDate.Year()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}
}

func TestDate_Year3(t *testing.T) {
	testDate := Date{82178}

	expectedYear := 225

	actualYear := testDate.Year()

	if expectedYear != actualYear {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYear)
	}
}

func TestDate_RomanYear(t *testing.T) {
	testDate := Date{82179}

	expectedYear := "CCXXVI"
	expectedYearInt := 226

	actualYearString := testDate.RomanYear().String()
	actualYearInt := testDate.RomanYear().Number()

	if expectedYear != actualYearString {
		t.Error("Expected:\t", expectedYear, "\tActual:\t", actualYearString)
	}

	if expectedYearInt != actualYearInt {
		t.Error("Expected:\t", expectedYearInt, "\tActual:\t", actualYearInt)
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
	testDate := NewDate(226, Prairial, 1)

	expectedWeekday := primidi
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

func TestDate_Weekday3(t *testing.T) {
	testDate := Date{82075}

	expectedWeekday := duodi
	actualWeekday := testDate.Weekday()

	if expectedWeekday != actualWeekday {
		t.Error("Expected:\t", expectedWeekday, "\tActual:\t", actualWeekday)
	}
}

func TestWeekday_String(t *testing.T) {
	testDate := NewDate(226, Prairial, 3)

	expectedWeekday := "tridi"
	actualWeekday := testDate.Weekday().String()

	if expectedWeekday != actualWeekday {
		t.Error("Expected:\t", expectedWeekday, "\tActual:\t", actualWeekday)
	}
}

func TestWeekday_String2(t *testing.T) {
	testDate := NewDate(226, Prairial, 2)

	expectedWeekday := "duodi"
	actualWeekday := testDate.Weekday().String()

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

func TestDate_IsLeapYear(t *testing.T) {
	testDate := Date{4386}

	actual := testDate.IsLeapYear()
	expected := false

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual)
	}
}

func TestDate_IsLeapYear2(t *testing.T) {
	testDate := Date{4020}

	actual := testDate.IsLeapYear()
	expected := true

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual)
	}
}

func TestDate_IsLeapYear3(t *testing.T) {
	testDate := Date{82078}

	actual := testDate.IsLeapYear()
	expected := false

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual)
	}
}

func TestDate_IsLeapYear4(t *testing.T) {
	testDate := Date{81713}

	actual := testDate.IsLeapYear()
	expected := true

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual)
	}
}

func TestDate_Day(t *testing.T) {
	testDate := Date{82174}

	actual := testDate.Day()
	expected := 1

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual, "\tDay of year", testDate.DayOfYear())
	}
}

func TestDate_Day2(t *testing.T) {
	testDate := Date{82175}

	actual := testDate.Day()
	expected := 2

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual, "\tDay of year", testDate.DayOfYear())
	}
}

func TestDate_Day3(t *testing.T) {
	testDate := Date{82176}

	actual := testDate.Day()
	expected := 3

	if actual != expected {
		t.Error("Expected:\t", expected, "\tActual:\t", actual, "\tDay of year", testDate.DayOfYear())
	}
}

func TestDate_DateFromTime(t *testing.T) {
	date := time.Date(2018, time.May, 21, 4, 0, 0, 0, time.Local)

	testDate := DateFromTime(date)

	expectedStr := "2 Prairial CCXXVI"
	actualStr := testDate.String()

	if actualStr != expectedStr {
		t.Error("Expected:\t", expectedStr, "\tActual:\t", actualStr, "\tDate", date.String())
	}
}

func TestDate_DateFromTime2(t *testing.T) {
	date := time.Date(1792, time.September, 22, 0, 0, 0, 0, time.Local)

	testDate := DateFromTime(date)

	expectedStr := "1 Vendémiaire I"
	actualStr := testDate.String()

	if actualStr != expectedStr {
		t.Error("Expected:\t", expectedStr, "\tActual:\t", actualStr, "\tDate", date.String())
	}
}

func TestCompDay_String(t *testing.T) {
	testDate := vertu

	expectedStr := "La Fête de la Vertu"
	actualStr := testDate.String()

	if actualStr != expectedStr {
		t.Error("Expected:\t", expectedStr, "\tActual:\t", actualStr,)
	}
}

func TestMonth_String(t *testing.T) {
	testDate := Pluviôse

	expectedStr := "Pluviôse"
	actualStr := testDate.String()

	if actualStr != expectedStr {
		t.Error("Expected:\t", expectedStr, "\tActual:\t", actualStr,)
	}
}
