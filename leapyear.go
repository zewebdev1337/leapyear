package leapyear

// Is checks whether the given year is a leap year or not
func Is(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
