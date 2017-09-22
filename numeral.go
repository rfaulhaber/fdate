package fdate

import (
	"bytes"
	"sort"
)

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

func NumeralFromNumber(number int) RomanNumeral {
	return RomanNumeral{number, convertToNumeralString(number)}
}

func NumeralFromString(numeral string) RomanNumeral {
	return RomanNumeral{convertToNumber(numeral), numeral}
}

func (n RomanNumeral) String() string {
	return n.numeral
}

func (n RomanNumeral) Number() int {
	return n.number
}

func (n RomanNumeral) Equal(r RomanNumeral) bool {
	return n.number == r.number && n.numeral == r.numeral
}

func convertToNumeralString(number int) string {
	var buffer bytes.Buffer

	numberArray := make([]int, 0)

	for key := range numberToNumeralMap {
		numberArray = append(numberArray, key)
	}

	sort.Ints(numberArray)

	if numeral, ok := numberToNumeralMap[number]; ok {
		buffer.WriteString(numeral)
	} else {
		for i := len(numberArray) - 1; i > 0; i-- {
			for numberArray[i] <= number {
				buffer.WriteString(numberToNumeralMap[numberArray[i]])
				number -= numberArray[i]
			}
		}

		if number == 1 {
			buffer.WriteString(numberToNumeralMap[1])
		}
	}

	return buffer.String()
}

func convertToNumber(numeral string) int {
	number := 0

	if len(numeral) == 1 {
		number = numeralToNumberMap[numeral]
	} else {
		for i := 0; i < len(numeral); i += 2 {
			fr := string(numeral[i])
			sr := string(numeral[i+1])
			r := string(fr + sr)

			if num, ok := numeralToNumberMap[r]; ok {
				number += num
			} else {
				number += numeralToNumberMap[fr] + numeralToNumberMap[sr]
			}
		}
	}

	return number
}
