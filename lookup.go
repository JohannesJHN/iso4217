package iso4217

// FindByAlpha3 looks up a currency by its three-letter alphabetic code.
// It returns the Currency and a boolean indicating whether it was found.
func FindByAlpha3(alpha3 string) (Currency, bool) {
	c, ok := currenciesByAlpha3[alpha3]
	return c, ok
}

// FindByNumeric looks up a currency by its three-digit numeric code.
// It returns the Currency and a boolean indicating whether it was found.
func FindByNumeric(numericCode int) (Currency, bool) {
	c, ok := currenciesByNumeric[numericCode]
	return c, ok
}

// GetAllCurrencies returns a cloned map of all currencies keyed by their three-letter alphabetic code.
func GetAllCurrencies() map[string]Currency {
	cloned := make(map[string]Currency, len(currenciesByAlpha3))
	for k, v := range currenciesByAlpha3 {
		cloned[k] = v
	}
	return cloned
}
