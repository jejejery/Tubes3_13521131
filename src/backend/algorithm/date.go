package algorithm

import (
	"math"
) 

func calculateDate(day int, month int, year int) string {
	// array with number of days each month
	daysMonth := [12]int{31, 28 + leap(year), 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	
	totalDays := (year - 1) * 365 + int(math.Floor(float64((year-1)/4))) - int(math.Floor(float64((year-1)/100))) + int(math.Floor(float64((year-1)/400))) + day

	// calculate days from the start of the year
	for i := 0; i < month-1; i++ {
		totalDays += daysMonth[i]
	}

	remainder := totalDays % 7

	resultday := mapRemainderOfTheWeek(remainder)

	return resultday
}

func mapRemainderOfTheWeek(remainder int) string {
	switch  remainder {
	case 0:
		return "Minggu"
	case 1:
		return "Senin"
	case 2:
		return "Selasa"
	case 3:
		return "Rabu"
	case 4:
		return "Kamis"
	case 5:
		return "Jumat"
	default:
		return "Sabtu"
	}
}

func leap(year int) int {
    if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return 1
	} else {
		return 0
	}
}

func leap_bool(year int) bool {
    if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	} else {
		return false
	}
}