package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ISO4217Active struct {
	CcyTbl struct {
		Ccies []struct {
			CtryNm     string `xml:"CtryNm"`
			Ccy        string `xml:"Ccy"`
			CcyNbr     string `xml:"CcyNbr"`
			CcyMnrUnts string `xml:"CcyMnrUnts"`
			CcyNm      struct {
				Name   string `xml:",chardata"`
				IsFund string `xml:"IsFund,attr"`
			} `xml:"CcyNm"`
		} `xml:"CcyNtry"`
	} `xml:"CcyTbl"`
}

type ISO4217Inactive struct {
	HstrcCcyTbl struct {
		Ccies []struct {
			CtryNm     string `xml:"CtryNm"`
			Ccy        string `xml:"Ccy"`
			CcyNbr     string `xml:"CcyNbr"`
			CcyNm      string `xml:"CcyNm"`
			WthdrwlDt  string `xml:"WthdrwlDt"`
		} `xml:"HstrcCcyNtry"`
	} `xml:"HstrcCcyTbl"`
}

type ISO3166 struct {
	Name   string
	Alpha2 string
	Alpha3 string
	Num    string
}

func loadISO3166(path string) (map[string]ISO3166, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	out := make(map[string]ISO3166)
	for i, rec := range records {
		if i == 0 {
			continue // header
		}
		if len(rec) < 4 {
			continue
		}
		out[normalize(rec[0])] = ISO3166{
			Name:   rec[0],
			Alpha2: rec[1],
			Alpha3: rec[2],
			Num:    rec[3],
		}
	}
	return out, nil
}

func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func main() {
	// parse ISO4217 XML
	f, err := os.Open("tools/generate/iso4217_list_1.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var activeData ISO4217Active
	if err := xml.NewDecoder(f).Decode(&activeData); err != nil {
		panic(err)
	}

	// parse ISO3166 CSV
	iso3166, err := loadISO3166("tools/generate/iso_3166_1.csv")
	if err != nil {
		panic(err)
	}

	out, err := os.Create("./currencies.go")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// type definition
	fmt.Fprintln(out, "package iso4217\n")
	fmt.Fprintln(out, "type Currency struct {")
	fmt.Fprintln(out, "\tAlpha3        string")
	fmt.Fprintln(out, "\tNumeric       int")
	fmt.Fprintln(out, "\tMinorUnits    int")
	fmt.Fprintln(out, "\tName          string")
	fmt.Fprintln(out, "\tWithdrawalDate string // Empty for active currencies")
	fmt.Fprintln(out, "}\n")

	seen := map[string]bool{}

	fmt.Fprintln(out, "var (")
	var numericMap strings.Builder
	numericMap.WriteString("var currenciesByNumeric = map[int]Currency{\n")
	numericSeen := make(map[int]bool)              // track numeric codes added to numeric map
	countryMap := make(map[string]string)          // alpha2 -> currency
	currencyCountries := make(map[string][]string) // currency -> []alpha2

	// Parse inactive currencies
	f2, err := os.Open("tools/generate/iso4217_list_3.xml")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	var inactiveData ISO4217Inactive
	if err := xml.NewDecoder(f2).Decode(&inactiveData); err != nil {
		panic(err)
	}

	for _, c := range activeData.CcyTbl.Ccies {
		if c.Ccy == "" || c.CcyNbr == "" {
			continue
		}

		// define all currencies
		if !seen[c.Ccy] {
			seen[c.Ccy] = true
			numeric, _ := strconv.Atoi(c.CcyNbr)
			minor := -1
			if c.CcyMnrUnts != "" && c.CcyMnrUnts != "N.A." {
				if v, err := strconv.Atoi(c.CcyMnrUnts); err == nil {
					minor = v
				}
			}

			fmt.Fprintf(out, "\t%s = Currency{Alpha3: \"%s\", Numeric: %d, MinorUnits: %d, Name: %q, WithdrawalDate: \"\"}\n",
				c.Ccy, c.Ccy, numeric, minor, c.CcyNm.Name)

			if !numericSeen[numeric] {
				numericSeen[numeric] = true
				numericMap.WriteString(fmt.Sprintf("\t%d: %s,\n", numeric, c.Ccy))
			}
		}

		// map countries only if not a fund
		if c.CcyNm.IsFund != "true" && c.CtryNm != "" {
			if info, ok := iso3166[normalize(c.CtryNm)]; ok {
				countryMap[info.Alpha2] = c.Ccy
				currencyCountries[c.Ccy] = append(currencyCountries[c.Ccy], info.Alpha2)
			} else {
				fmt.Println("unknown country:", c.CtryNm)
			}
		}
	}

	// Process inactive currencies
	for _, c := range inactiveData.HstrcCcyTbl.Ccies {
		if c.Ccy == "" {
			continue
		}

		// define inactive currencies (only if not already defined as active)
		if !seen[c.Ccy] {
			seen[c.Ccy] = true
			numeric := 0
			if c.CcyNbr != "" {
				numeric, _ = strconv.Atoi(c.CcyNbr)
			}
			// Inactive currencies don't have minor units in the XML
			minor := -1

			fmt.Fprintf(out, "\t%s = Currency{Alpha3: \"%s\", Numeric: %d, MinorUnits: %d, Name: %q, WithdrawalDate: %q}\n",
				c.Ccy, c.Ccy, numeric, minor, c.CcyNm, c.WthdrwlDt)

			// Don't add inactive currencies to numeric map to avoid conflicts
		}
	}
	fmt.Fprintln(out, ")\n")

	numericMap.WriteString("}\n")
	fmt.Fprintln(out, numericMap.String())

	fmt.Fprintln(out, "var currenciesByCountry = map[string]Currency{")
	for alpha2, cur := range countryMap {
		fmt.Fprintf(out, "\t\"%s\": %s,\n", alpha2, cur)
	}
	fmt.Fprintln(out, "}\n")

	fmt.Fprintln(out, "var countryList = map[Currency][]string{")
	for cur, countries := range currencyCountries {
		fmt.Fprintf(out, "\t%s: %#v,\n", cur, countries)
	}
	fmt.Fprintln(out, "}\n")

	// Create separate maps for active and inactive currencies
	activeCurrenciesSet := make(map[string]bool)
	for _, c := range activeData.CcyTbl.Ccies {
		if c.Ccy != "" {
			activeCurrenciesSet[c.Ccy] = true
		}
	}

	fmt.Fprintln(out, "var activeCurrencies = map[string]Currency{")
	for currency := range activeCurrenciesSet {
		fmt.Fprintf(out, "\t\"%s\": %s,\n", currency, currency)
	}
	fmt.Fprintln(out, "}\n")

	fmt.Fprintln(out, "var inactiveCurrencies = map[string]Currency{")
	inactiveCurrenciesSet := make(map[string]bool)
	for _, c := range inactiveData.HstrcCcyTbl.Ccies {
		if c.Ccy != "" && !activeCurrenciesSet[c.Ccy] {
			inactiveCurrenciesSet[c.Ccy] = true
		}
	}
	for currency := range inactiveCurrenciesSet {
		fmt.Fprintf(out, "\t\"%s\": %s,\n", currency, currency)
	}
	fmt.Fprintln(out, "}\n")
}
