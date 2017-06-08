package fdate

import (
	"time"
)

type Month int

const (
	Vendémiaire Month = 1 + iota
	Brumaire
	Frimaire
	Nivôse
	Pluviôse
	Ventôse
	Germinal
	Floréal
	Prairial
	Messidor
	Thermidor
	Fructidor
	Complémentaires
)

var months = [...]string{
	"Vendémiaire",
	"Brumaire",
	"Frimaire",
	"Nivôse",
	"Pluviôse",
	"Ventôse",
	"Germinal",
	"Floréal",
	"Prairial",
	"Messidor",
	"Thermidor",
	"Fructidor",
	"Complémentaires",
}

func (month Month) String() string {
	return months[month-1]
}

type Weekday int

const (
	primidi Weekday = iota
	duodi
	tridi
	quartidi
	quintidi
	sextidi
	septidi
	octidi
	nonidi
	décadi
)

var days = [...]string{
	"primidi",
	"duodi",
	"tridi",
	"quartidi",
	"quintidi",
	"sextidi",
	"septidi",
	"octidi",
	"nonidi",
	"décadi",
}

func (day Weekday) String() string {
	return days[day]
}

type CompDay int

const (
	vertu CompDay = iota
	génie
	travail
	lOpinion
	récompenses
	révolution
)

var compDays = [...]string{
	"La Fête de la Vertu",
	"La Fête du Génie",
	"La Fête du Travail",
	"La Fête de l'Opinion",
	"La Fête des Récompenses",
	"La Fête de la Révolution",
}

func (day CompDay) String() string {
	return compDays[day]
}

type Duration struct {
	Days   int
	Months int
	Years  int
}

var startLocation, _ = time.LoadLocation("Europe/Paris")

var startDate = time.Date(1792, time.September, 22, 0, 0, 0, 0, startLocation)

const (
	daysPer400Years = 365*400 + 97
	daysPer100Years = 365*100 + 24
	daysPer4Years   = 365*4 + 1
)

type Date struct {
	// number of days since the start of the Republican calendar
	days int
}

func Today() Date {
	return Date{int((time.Since(startDate) * time.Hour).Hours() / 24)}
}

// returns a Date object where the ints correspond to FRC values and not Gregorian values
func NewDate(year int, month Month, day int) Date {
	// TODO: add validation

	// Add in days from 400-year cycles.
	y := uint64(year - 1)
	n := y / 400
	y -= 400 * n
	d := daysPer400Years * n

	// Add in 100-year cycles.
	n = y / 100
	y -= 100 * n
	d += daysPer100Years * n

	// Add in 4-year cycles.
	n = y / 4
	y -= 4 * n
	d += daysPer4Years * n

	// Add in non-leap years.
	n = y
	d += 365 * n

	// Add in days before today.
	if month == 13 {
		d += uint64(daysBefore[month-1]) + uint64(day)
	} else {
		d += uint64(daysBefore[month-1]) + uint64(day-1)
	}

	date := Date{int(d)}
	return date
}

func DateFromTime(time time.Time) Date {
	return Date{daysSince(startDate, time)}
}

//func Parse(value string) (Date, error) {
//
//}

//func (d Date) String() string {
//}

func (d Date) Date() (year int, month Month, day int) {
	year, month, day, _ = d.date()
	return
}

func (d Date) Year() int {
	year, _, _, _ := d.date()
	return year
}

func (d Date) RomanYear() RomanNumeral {
	return NumeralFromNumber(d.Year())
}

func (d Date) Month() Month {
	_, month, _, _ := d.date()
	return month
}

func (d Date) Day() int {
	_, _, day, _ := d.date()
	return day
}

func (d Date) DayOfYear() int {
	_, _, _, yday := d.date()
	return yday
}

//func (d Date) Weekday() Weekday {
//	// TODO: calculate weekday
//}

func (d Date) After(u Date) bool {
	return d.days > u.days
}

func (d Date) Before(u Date) bool {
	return d.days < u.days
}

func (d Date) Equals(u Date) bool {
	return d.days == u.days
}

func daysSince(start time.Time, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

func (date Date) date() (year int, month Month, day int, yday int) {
	// this was mostly taken from the code from the method in the time package of the same name
	// to see the original, see this https://golang.org/src/time/time.go

	d := date.days

	n := d / daysPer400Years
	y := 400 * n
	d -= daysPer400Years * n

	n = d / daysPer100Years
	n -= n >> 2
	y += 100 * n
	d -= daysPer100Years * n

	n = d / daysPer4Years
	y += 4 * n
	d -= daysPer4Years * n

	n = d / 365
	n -= n >> 2
	y += n
	d -= 365 * n

	year = int(int64(y) + 1)
	yday = int(d)

	day = yday

	if day > 360 {
		month = Complémentaires
		day = day - 360
	} else {
		month = Month(day / 30)
		end := int(daysBefore[month+1])
		var begin int
		if day >= end {
			month++
			begin = end
		} else {
			begin = int(daysBefore[month])
		}

		month++
		day = day - begin + 1
	}

	return
}

func isLeap(year int) bool {
	normalizedYear := year + startDate.Year()
	return normalizedYear%4 == 0 && (normalizedYear%100 != 0 || normalizedYear%400 == 0)
}

func norm(hi, lo, base int) (nhi, nlo int) {
	if lo < 0 {
		n := (-lo-1)/base + 1
		hi -= n
		lo += n * base
	}
	if lo >= base {
		n := lo / base
		hi += n
		lo -= n * base
	}
	return hi, lo
}

var daysBefore = [...]int32{
	0,
	30,
	2 * 30,
	3 * 30,
	4 * 30,
	5 * 30,
	6 * 30,
	7 * 30,
	8 * 30,
	9 * 30,
	10 * 30,
	11 * 30,
	12 * 30,
}
