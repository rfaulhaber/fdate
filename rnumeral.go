package fdate

type RomanNumeral struct {
	number  int
	numeral string
}

const pattern = "^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$"

var numberToNumeralMap = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var numeralToNumberMap = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func NewNumeral(number int) *RomanNumeral {
	return &RomanNumeral{number, convertToNumeralString(number)}
}

func (n *RomanNumeral) String() string {
	return n.numeral
}

func (n *RomanNumeral) Number() int {
	return n.number
}

func (n *RomanNumeral) Equal(r RomanNumeral) bool {
	return n.number == r.number && n.numeral == r.numeral
}

func convertToNumeralString(number int) string {
	thousands := number / 1000
	hundreds := (number - (thousands * 1000)) / 100
	tens := (number - (thousands * 1000) - (hundreds * 100)) / 10
	ones := (number - (thousands * 1000) - (hundreds * 100) - (tens * 10))
}
