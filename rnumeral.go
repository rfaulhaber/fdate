package fdate

import "bytes"

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

func NumeralFromNumber(number int) *RomanNumeral {
	return &RomanNumeral{number, convertToNumeralString(number)}
}

func NumeralFromString(numeral string) *RomanNumeral {
	return &RomanNumeral{convertToNumber(numeral), numeral}
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
	var buffer bytes.Buffer

	for key := range numberToNumeralMap {
		for key <= number {
			buffer.WriteString(numberToNumeralMap[key])
			number -= key
		}
	}

	return buffer.String()
}

func convertToNumber(numeral string) int {
	number := 0

	previousEntered := false

	for i, r := range numeral {
		if num, ok := numeralToNumberMap[string(r) + string(numeral[i + 1])]; ok {
			number += num
			previousEntered = true
		} else {
			if previousEntered {
				previousEntered = false
				break
			} else {
				number += numeralToNumberMap[string(r)]
				previousEntered = false
			}
		}
	}

	return number
}
