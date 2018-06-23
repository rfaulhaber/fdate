package fdate

import (
	"strconv"
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

var ruralNames = [...]string{
	"Raisin", "Safran", "Châtaigne", "Colchique", "Cheval", "Balsamine", "Carotte", "Amaranthe", "Panais", "Cuve", "Pomme de terre", "Immortelle", "Potiron", "Réséda", "Âne", "Belle de nuit", "Citrouille", "Sarrasin", "Tournesol", "Pressoir", "Chanvre", "Pêche", "Navet", "Amaryllis", "Bœuf", "Aubergine", "Piment", "Tomate", "Orge", "Tonneau", "Pomme", "Céleri", "Poire", "Betterave", "Oie", "Héliotrope", "Figue", "Scorsonère", "Alisier", "Charrue", "Salsifis", "Mâcre", "Topinambour", "Endive", "Dindon", "Chervis", "Cresson", "Dentelaire", "Grenade", "Herse", "Bacchante", "Azerole", "Garance", "Orange", "Faisan", "Pistache", "Macjonc", "Coing", "Cormier", "Rouleau", "Raiponce", "Turneps", "Chicorée", "Nèfle", "Cochon", "Mâche", "Chou-fleur", "Miel", "Genièvre", "Pioche", "Cire", "Raifort", "Cèdre", "Sapin", "Chevreuil", "Ajonc", "Cyprès", "Lierre", "Sabine", "Hoyau", "Érable à sucre", "Bruyère", "Roseau", "Oseille", "Grillon", "Pignon", "Liège", "Truffe", "Olive", "Pelle", "Tourbe", "Houille", "Bitume", "Soufre", "Chien", "Lave", "Terre végétale", "Fumier", "Salpêtre", "Fléau", "Granit", "Argile", "Ardoise", "Grès", "Lapin", "Silex", "Marne", "Pierre à chaux", "Marbre", "Van", "Pierre à plâtre", "Sel", "Fer", "Cuivre", "Chat", "Étain", "Plomb", "Zinc", "Mercure", "Crible", "Lauréole", "Mousse", "Fragon", "Perce-neige", "Taureau", "Laurier-thym", "Amadouvier", "Mézéréon", "Peuplier", "Coignée", "Ellébore", "Brocoli", "Laurier", "Avelinier", "Vache", "Buis", "Lichen", "If", "Pulmonaire", "Serpette", "Thlaspi", "Thimelé", "Chiendent", "Trainasse", "Lièvre", "Guède", "Noisetier", "Cyclamen", "Chélidoine", "Traîneau", "Tussilage", "Cornouiller", "Violier", "Troène", "Bouc", "Asaret", "Alaterne", "Violette", "Marceau", "Bêche", "Narcisse", "Orme", "Fumeterre", "Vélar", "Chèvre", "Épinard", "Doronic", "Mouron", "Cerfeuil", "Cordeau", "Mandragore", "Persil", "Cochléaria", "Pâquerette", "Thon", "Pissenlit", "Sylvie", "Capillaire", "Frêne", "Plantoir", "Primevère", "Platane", "Asperge", "Tulipe", "Poule", "Bette", "Bouleau", "Jonquille", "Aulne", "Couvoir", "Pervenche", "Charme", "Morille", "Hêtre", "Abeille", "Laitue", "Mélèze", "Ciguë", "Radis", "Ruche", "Gainier", "Romaine", "Marronnier", "Roquette", "Pigeon", "Lilas", "Anémone", "Pensée", "Myrtille", "Greffoir", "Rose", "Chêne", "Fougère", "Aubépine", "Rossignol", "Ancolie", "Muguet", "Champignon", "Hyacinthe", "Râteau", "Rhubarbe", "Sainfoin", "Bâton d'or", "Chamerisier", "Ver à soie", "Consoude", "Pimprenelle", "Corbeille d'or", "Arroche", "Sarcloir", "Statice", "Fritillaire", "Bourrache", "Valériane", "Carpe", "Fusain", "Civette", "Buglosse", "Sénevé", "Houlette", "Luzerne", "Hémérocalle", "Trèfle", "Angélique", "Canard", "Mélisse", "Fromental", "Martagon", "Serpolet", "Faux", "Fraise", "Bétoine", "Pois", "Acacia", "Caille", "Œillet", "Sureau", "Pavot", "Tilleul", "Fourche", "Barbeau", "Camomille", "Chèvrefeuille", "Caille-lait", "Tanche", "Jasmin", "Verveine", "Thym", "Pivoine", "Chariot", "Seigle", "Avoine", "Oignon", "Véronique", "Mulet", "Romarin", "Concombre", "Échalote", "Absinthe", "Faucille", "Coriandre", "Artichaut", "Girofle", "Lavande", "Chamois", "Tabac", "Groseille", "Gesse", "Cerise", "Parc", "Menthe", "Cumin", "Haricot", "Orcanète", "Pintade", "Sauge", "Ail", "Vesce", "Blé", "Chalémie", "Épeautre", "Bouillon blanc", "Melon", "Ivraie", "Bélier", "Prêle", "Armoise", "Carthame", "Mûre", "Arrosoir", "Panic", "Salicorne", "Abricot", "Basilic", "Brebis", "Guimauve", "Lin", "Amande", "Gentiane", "Écluse", "Carline", "Câprier", "Lentille", "Aunée", "Loutre", "Myrte", "Colza", "Lupin", "Coton", "Moulin", "Prune", "Millet", "Lycoperdon", "Escourgeon", "Saumon", "Tubéreuse", "Sucrion", "Apocyn", "Réglisse", "Échelle", "Pastèque", "Fenouil", "Épine vinette", "Noix", "Truite", "Citron", "Cardère", "Nerprun", "Tagette", "Hotte", "Églantier", "Noisette", "Houblon", "Sorgho", "Écrevisse", "Bigarade", "Verge d'or", "Maïs", "Marron", "Panier",
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

// this makes the assumption that this calendar starts globally, regardless of timezone, on this date
var startDate = time.Date(1792, time.September, 22, 0, 0, 0, 0, time.Local)

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
	return Date{int(time.Since(startDate) / (24 * time.Hour))}
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
		d += uint64(daysBefore(month-1)) + uint64(day)
	} else {
		d += uint64(daysBefore(month-1)) + uint64(day-1)
	}

	date := Date{int(d)}
	return date
}

func DateFromTime(time time.Time) Date {
	return Date{daysSince(startDate, time)}
}

func (d Date) String() string {
	year, month, day, _ := d.date()

	numeralYear := NumeralFromNumber(year)
	monthStr := months[month-1]
	dayStr := strconv.Itoa(day)

	return dayStr + " " + monthStr + " " + numeralYear.String()
}

func (d Date) RuralDay() string {
	_, _, _, yday := d.date()

	if yday > 360 {
		return ""
	}

	return ruralNames[yday]
}

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

func (d Date) Weekday() Weekday {
	_, _, day, yday := d.date()

	if d.Month() < 13 {
		return Weekday(yday % 10)
	} else {
		return Weekday(day % 10)
	}
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

func (d Date) IsLeapYear() bool {
	year := startDate.Year() + d.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func daysSince(start time.Time, end time.Time) int {
	return int(end.Sub(start) / (24 * time.Hour))
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
		day = day - 360 + 1
	} else {
		month = Month(day / 30)
		end := int(daysBefore(month + 1))
		var begin int
		if day >= end {
			month++
			begin = end
		} else {
			begin = int(daysBefore(month))
		}

		month++
		day = day - begin + 1
	}

	return
}

func daysBefore(month Month) int {
	return 30 * int(month)
}
