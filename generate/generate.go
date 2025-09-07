package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ISO4217 struct {
	CcyTbl struct {
		Ccies []struct {
			CtryNm     string `xml:"CtryNm"`
			Ccy        string `xml:"Ccy"`
			CcyNbr     string `xml:"CcyNbr"`
			CcyMnrUnts string `xml:"CcyMnrUnts"`
			CcyNm      string `xml:"CcyNm"`
		} `xml:"CcyNtry"`
	} `xml:"CcyTbl"`
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
	f, err := os.Open("generate/iso4217-list.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var data ISO4217
	if err := xml.NewDecoder(f).Decode(&data); err != nil {
		panic(err)
	}

	// parse ISO3166 CSV
	iso3166, err := loadISO3166("generate/iso_3166_1.csv")
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
	fmt.Fprintln(out, "\tAlpha3     string")
	fmt.Fprintln(out, "\tNumeric    int")
	fmt.Fprintln(out, "\tMinorUnits int")
	fmt.Fprintln(out, "\tName       string")
	fmt.Fprintln(out, "}\n")

	seen := map[string]bool{}

	// per-currency vars
	fmt.Fprintln(out, "var (")
	var alpha3map strings.Builder
	alpha3map.WriteString("var currenciesByAlpha3 = map[string]Currency{\n")
	var numericMap strings.Builder
	numericMap.WriteString("var currenciesByNumeric = map[int]Currency{\n")
	var countryMap = make(map[string]string)          // alpha2 -> currency
	var currencyCountries = make(map[string][]string) // currency -> []alpha2

	for _, c := range data.CcyTbl.Ccies {
		if c.Ccy == "" || c.CcyNbr == "" {
			continue
		}

		// define currency once
		if !seen[c.Ccy] {
			seen[c.Ccy] = true

			numeric, _ := strconv.Atoi(c.CcyNbr)
			minor := -1
			if c.CcyMnrUnts != "" && c.CcyMnrUnts != "N.A." {
				if v, err := strconv.Atoi(c.CcyMnrUnts); err == nil {
					minor = v
				}
			}

			fmt.Fprintf(out, "\t%s = Currency{Alpha3: \"%s\", Numeric: %d, MinorUnits: %d, Name: %q}\n",
				c.Ccy, c.Ccy, numeric, minor, c.CcyNm)

			alpha3map.WriteString(fmt.Sprintf("\t\"%s\": %s,\n", c.Ccy, c.Ccy))
			numericMap.WriteString(fmt.Sprintf("\t%d: %s,\n", numeric, c.Ccy))
		}

		// link to ISO3166 country
		if c.CtryNm != "" {
			if info, ok := iso3166[normalize(c.CtryNm)]; ok {
				countryMap[info.Alpha2] = c.Ccy
				currencyCountries[c.Ccy] = append(currencyCountries[c.Ccy], info.Alpha2)
			} else {
				fmt.Println("unknown country:", c.CtryNm)
			}

		}
	}
	fmt.Fprintln(out, ")\n")

	// maps
	alpha3map.WriteString("}\n")
	fmt.Fprintln(out, alpha3map.String())

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
}
