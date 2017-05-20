package fdate

type FDate struct {
}

type Month int

const (
	Vendémiaire = 1 + iota
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

type Weekday int

const (
	primidi  = 1 + iota
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
