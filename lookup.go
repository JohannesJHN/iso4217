package iso4217

// LookupByAlpha3 looks up a currency by its three-letter alphabetic code.
// It returns the Currency and a boolean indicating whether it was found.
// It searches both active and inactive currencies.
func LookupByAlpha3(alpha3 string) (Currency, bool) {
	// First check active currencies
	if c, ok := activeCurrencies[alpha3]; ok {
		return c, true
	}
	// Then check inactive currencies
	if c, ok := inactiveCurrencies[alpha3]; ok {
		return c, true
	}
	return Currency{}, false
}

// LookupByNumeric looks up a currency by its three-digit numeric code.
// It returns the Currency and a boolean indicating whether it was found.
// Only active currencies are searched to avoid conflicts with inactive currencies.
func LookupByNumeric(numericCode int) (Currency, bool) {
	c, ok := currenciesByNumeric[numericCode]
	return c, ok
}

// All returns a cloned map of all currencies keyed by their three-letter alphabetic code.
// This includes both active and inactive currencies.
func All() map[string]Currency {
	cloned := make(map[string]Currency, len(activeCurrencies)+len(inactiveCurrencies))
	for k, v := range activeCurrencies {
		cloned[k] = v
	}
	for k, v := range inactiveCurrencies {
		cloned[k] = v
	}
	return cloned
}

// AllActive returns a cloned map of all active currencies keyed by their three-letter alphabetic code.
func AllActive() map[string]Currency {
	cloned := make(map[string]Currency, len(activeCurrencies))
	for k, v := range activeCurrencies {
		cloned[k] = v
	}
	return cloned
}

// AllInactive returns a cloned map of all inactive currencies keyed by their three-letter alphabetic code.
func AllInactive() map[string]Currency {
	cloned := make(map[string]Currency, len(inactiveCurrencies))
	for k, v := range inactiveCurrencies {
		cloned[k] = v
	}
	return cloned
}

// LookupByCountry looks up a currency by its two-letter ISO 3166-1 alpha-2 country code.
// It returns the Currency and a boolean indicating whether it was found.
// Only active currencies are returned to avoid conflicts with inactive currencies.
func LookupByCountry(countryCode string) (Currency, bool) {
	c, ok := currenciesByCountry[countryCode]
	return c, ok
}

// String returns the three-letter alphabetic code of the currency.
func (c Currency) String() string {
	return c.Alpha3
}

// Countries returns the list of ISO 3166-1 alpha-2 country codes that use this currency.
// Returns nil if no countries use this currency or if the currency is inactive.
func (c Currency) Countries() []string {
	if countries, ok := countryList[c]; ok {
		return countries
	}
	return nil
}

// IsActive returns true if the currency is currently active (not withdrawn).
// A currency is considered active if it has no withdrawal date.
func (c Currency) IsActive() bool {
	return c.WithdrawalDate == ""
}
