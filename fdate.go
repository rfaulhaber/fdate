package fdate

import "time"

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

type Date struct {
	// number of days since the start of the Republican calendar
	days int
}

func Today() Date {
	return Date{int((time.Since(startDate) * time.Hour).Hours() / 24)}
}

func NewDate(year int, month Month, day int) Date {
	// TODO: treat arguments as those on FRC, convert to gregorian and call datefromtime?
}

func DateFromTime(time time.Time) Date {
	return Date{daysSince(startDate, time)}
}

func Parse(value string) (Date, error) {

}

func (date Date) String() string {
}

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
