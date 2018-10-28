package fdate

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDate_Date(t *testing.T) {
	testCases := []struct {
		testDate      Date
		expectedYear  int
		expectedMonth Month
		expectedDay   int
	}{
		{
			testDate:      Date{0},
			expectedYear:  1,
			expectedMonth: Vendémiaire,
			expectedDay:   1,
		},
		{

			testDate: Date{2603},

			// The 18th Brumaire, when Napoleon staged his famous coup, ending the Revolution
			expectedYear:  8,
			expectedMonth: Brumaire,
			expectedDay:   18,
		},
		{
			testDate: Date{673},

			// The 9th Thermidor, when the Convention turned against Robespierre
			expectedYear:  2,
			expectedMonth: Thermidor,
			expectedDay:   9,
		},
		{
			testDate: Date{81814},

			expectedYear:  225,
			expectedMonth: Vendémiaire,
			expectedDay:   1,
		},
		{
			testDate: Date{81850},

			expectedYear:  225,
			expectedMonth: Brumaire,
			expectedDay:   7,
		},
		{
			testDate: Date{82215},

			expectedYear:  226,
			expectedMonth: Brumaire,
			expectedDay:   7,
		},
		{
			testDate: Date{82580},

			expectedYear:  227,
			expectedMonth: Brumaire,
			expectedDay:   7,
		},
	}

	for _, tc := range testCases {
		actualYear, actualMonth, actualDay := tc.testDate.Date()

		assert.Equal(t, tc.expectedYear, actualYear)
		assert.Equal(t, tc.expectedMonth, actualMonth)
		assert.Equal(t, tc.expectedDay, actualDay)
	}
}

func TestDate_Year(t *testing.T) {
	testCases := []struct {
		testDate     Date
		expectedYear int
	}{
		{
			testDate:     Date{673},
			expectedYear: 2,
		},
		{
			testDate:     Date{82580},
			expectedYear: 227,
		},
		{
			testDate:     Date{82215},
			expectedYear: 226,
		},
		{
			testDate:     Date{81850},
			expectedYear: 225,
		},
	}

	for _, tc := range testCases {
		actualYear := tc.testDate.Year()

		if tc.expectedYear != actualYear {
			t.Error("Expected:\t", tc.expectedYear, "\tActual:\t", actualYear)
		}
	}
}

func TestDate_RomanYear(t *testing.T) {
	testCases := []struct {
		testDate           Date
		expectedYearString string
		expectedYearInt    int
	}{
		{
			testDate:           Date{82179},
			expectedYearString: "CCXXVI",
			expectedYearInt:    226,
		},
	}

	for _, tc := range testCases {
		actualYearString := tc.testDate.RomanYear().String()
		actualYearInt := tc.testDate.RomanYear().Number()

		if tc.expectedYearString != actualYearString {
			t.Error("Expected:\t", tc.expectedYearString, "\tActual:\t", actualYearString)
		}

		if tc.expectedYearInt != actualYearInt {
			t.Error("Expected:\t", tc.expectedYearInt, "\tActual:\t", actualYearInt)
		}
	}
}

func TestDate_Month(t *testing.T) {
	testCases := []struct {
		testDate      Date
		expectedMonth Month
	}{
		{
			testDate:      Date{2603},
			expectedMonth: Brumaire,
		},
	}

	for _, tc := range testCases {
		actualMonth := tc.testDate.Month()

		if tc.expectedMonth != actualMonth {
			t.Error("Expected:\t", tc.expectedMonth, "\tActual:\t", actualMonth)
		}
	}
}

func TestDate_NewDate(t *testing.T) {
	testCases := []struct {
		testYear     int
		testMonth    Month
		testDay      int
		expectedDate Date
	}{
		{
			expectedDate: Date{82061},

			testYear:  225,
			testMonth: Prairial,
			testDay:   8,
		},
		{
			expectedDate: Date{0},

			testYear:  1,
			testMonth: Vendémiaire,
			testDay:   1,
		},
		{
			expectedDate: Date{2603},

			// The 18th Brumaire, when Napoleon staged his famous coup, ending the Revolution
			testYear:  8,
			testMonth: Brumaire,
			testDay:   18,
		},
		{
			expectedDate: Date{673},

			// The 9th Thermidor, when the Convention turned against Robespierre
			testYear:  2,
			testMonth: Thermidor,
			testDay:   9,
		},
		{
			expectedDate: Date{81814},

			testYear:  225,
			testMonth: Vendémiaire,
			testDay:   1,
		},
		{
			expectedDate: Date{81813},

			testYear:  224,
			testMonth: Complémentaires,
			testDay:   5,
		},
		{
			expectedDate: Date{81810},

			testYear:  224,
			testMonth: Complémentaires,
			testDay:   2,
		},
	}

	for _, tc := range testCases {
		actualDate := NewDate(tc.testYear, tc.testMonth, tc.testDay)

		if tc.expectedDate.days != actualDate.days {
			t.Error("Expected:\t", tc.expectedDate.days, "\tActual:\t", actualDate.days)
		}
	}
}

func TestDate_Weekday(t *testing.T) {
	testCases := []struct {
		testDate        Date
		expectedWeekday Weekday
	}{
		{
			testDate:        NewDate(226, Prairial, 1),
			expectedWeekday: primidi,
		},
		{
			testDate:        Date{82074},
			expectedWeekday: primidi,
		},
		{
			testDate:        Date{82075},
			expectedWeekday: duodi,
		},
	}

	for _, tc := range testCases {
		actualWeekday := tc.testDate.Weekday()

		if tc.expectedWeekday != actualWeekday {
			t.Error("Expected:\t", tc.expectedWeekday, "\tActual:\t", actualWeekday)
		}
	}
}

func TestWeekday_String(t *testing.T) {
	testCases := []struct {
		testDate        Date
		expectedWeekday string
	}{
		{
			testDate:        NewDate(226, Prairial, 3),
			expectedWeekday: "tridi",
		},
		{
			testDate:        NewDate(226, Prairial, 2),
			expectedWeekday: "duodi",
		},
	}

	for _, tc := range testCases {
		actualWeekday := tc.testDate.Weekday().String()

		if tc.expectedWeekday != actualWeekday {
			t.Error("Expected:\t", tc.expectedWeekday, "\tActual:\t", actualWeekday)
		}
	}
}

func TestDate_String(t *testing.T) {
	testCases := []struct {
		testDate       Date
		expectedString string
	}{
		{
			testDate:       Date{82074},
			expectedString: "21 Prairial CCXXV",
		},
	}

	for _, tc := range testCases {
		actual := tc.testDate.String()

		if actual != tc.expectedString {
			t.Error("Expected:\t", tc.expectedString, "\tActual:\t", actual)
		}
	}
}

func TestDate_IsLeapYear(t *testing.T) {
	testCases := []struct{
		testDate Date
		expected bool
	}{
		{
			testDate: Date{4386},
			expected: false,
		},
		{
			testDate : Date{4020},
			expected : true,
		},
		{
			testDate : Date{82078},
			expected : false,
		},
		{
			testDate : Date{81713},
			expected : true,
		},
	}

	for _, tc := range testCases {
		actual := tc.testDate.IsLeapYear()

		if actual != tc.expected {
			t.Error("Expected:\t", tc.expected, "\tActual:\t", actual)
		}
	}
}

func TestDate_Day(t *testing.T) {
	testCases := []struct{
		testDate Date
		expectedDate int
	}{
		{
			testDate: Date{82174},
			expectedDate: 1,
		},
		{
			testDate: Date{82175},
			expectedDate: 2,
		},
		{
			testDate: Date{82176},
			expectedDate: 3,
		},
	}

	for _, tc := range testCases {
		actual := tc.testDate.Day()

		if actual != tc.expectedDate {
			t.Error("Expected:\t", tc.expectedDate, "\tActual:\t", actual, "\tDay of year", tc.testDate.DayOfYear())
		}
	}

}

func TestDate_DateFromTime(t *testing.T) {
	testCases := []struct{
		gregorianDate time.Time
		expectedStr string
	}{
		{
			gregorianDate: time.Date(2018, time.May, 21, 4, 0, 0, 0, time.Local),
			expectedStr : "2 Prairial CCXXVI",
		},
		{
			gregorianDate: time.Date(1792, time.September, 22, 0, 0, 0, 0, time.Local),
			expectedStr : "1 Vendémiaire I",
		},
	}

	for _, tc := range testCases {
		actualStr := DateFromTime(tc.gregorianDate).String()

		if actualStr != tc.expectedStr {
			t.Error("Expected:\t", tc.expectedStr, "\tActual:\t", actualStr, "\tDate", tc.gregorianDate.String())
		}
	}
}

func TestCompDay_String(t *testing.T) {
	testDate := vertu

	expectedStr := "La Fête de la Vertu"
	actualStr := testDate.String()

	if actualStr != expectedStr {
		t.Error("Expected:\t", expectedStr, "\tActual:\t", actualStr)
	}
}

func TestMonth_String(t *testing.T) {
	testDate := Pluviôse

	expectedStr := "Pluviôse"
	actualStr := testDate.String()

	if actualStr != expectedStr {
		t.Error("Expected:\t", expectedStr, "\tActual:\t", actualStr)
	}
}

func TestRuralDay(t *testing.T) {
	testCases := []struct{
		testDate Date
		expectedStr string
	}{
		{
			testDate: NewDate(226, Messidor, 5),
			expectedStr : "Mulet",
		},
		{
			testDate : NewDate(1, Vendémiaire, 1),
			expectedStr : "Raisin",
		},
		{
			testDate : NewDate(1, Complémentaires, 1),
			expectedStr : "",
		},
		{
			testDate : NewDate(226, Fructidor, 30),
			expectedStr : "Panier",
		},

	}

	for _, tc := range testCases {
		actualStr := tc.testDate.RuralDay()

		if actualStr != tc.expectedStr {
			t.Error("Expected:\t", tc.expectedStr, "\tActual:\t", actualStr)
		}
	}
}
