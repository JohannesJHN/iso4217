package iso4217

// FindByAlpha looks up a currency by its three-letter alphabetic code.
// It returns the Currency and a boolean indicating whether it was found.
func FindByAlpha(alpha string) (Currency, bool) {
	c, ok := CurrenciesByAlpha[alpha]
	return c, ok
}

// FindByNumeric looks up a currency by its three-digit numeric code.
// It returns the Currency and a boolean indicating whether it was found.
func FindByNumeric(numeric int16) (Currency, bool) {
	c, ok := CurrenciesByNumeric[numeric]
	return c, ok
}
