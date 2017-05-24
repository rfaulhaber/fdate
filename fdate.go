package fdate

import "time"

type FDate struct {
	Month int
	Day   int
	Year  int
}

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
}

func (month Month) String() string {
	return months[month-1]
}

type Weekday int

const (
	primidi  Weekday = iota
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
	vertu       CompDay = iota
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

func Today() *FDate {

}

func NewFDate(day int, month int, year int) *FDate {
}

func NewDateFromTime(time time.Time) *FDate {

}

func (date *FDate) String() string {

}
