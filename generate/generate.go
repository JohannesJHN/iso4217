package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type ISO4217 struct {
	CcyTbl struct {
		Ccies []struct {
			Ccy        string `xml:"Ccy"`
			CcyNbr     string `xml:"CcyNbr"`
			CcyMnrUnts string `xml:"CcyMnrUnts"`
			CcyNm      string `xml:"CcyNm"`
		} `xml:"CcyNtry"`
	} `xml:"CcyTbl"`
}

func main() {
	f, err := os.Open("generate/iso4217-list.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var data ISO4217
	if err := xml.NewDecoder(f).Decode(&data); err != nil {
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

	for _, c := range data.CcyTbl.Ccies {
		if c.Ccy == "" || c.CcyNbr == "" || seen[c.Ccy] {
			continue
		}
		seen[c.Ccy] = true

		var numeric int
		fmt.Sscanf(c.CcyNbr, "%d", &numeric)

		minor := -1
		if c.CcyMnrUnts != "" && c.CcyMnrUnts != "N.A." {
			fmt.Sscanf(c.CcyMnrUnts, "%d", &minor)
		}

		fmt.Fprintf(out, "\t%s = Currency{Alpha3: \"%s\", Numeric: %d, MinorUnits: %d, Name: \"%s\"}\n",
			c.Ccy, c.Ccy, numeric, minor, strings.ReplaceAll(c.CcyNm, "\"", ""))

		alpha3map.WriteString(fmt.Sprintf("\t\"%s\": %s,\n", c.Ccy, c.Ccy))
		numericMap.WriteString(fmt.Sprintf("\t%d: %s,\n", numeric, c.Ccy))
	}
	fmt.Fprintln(out, ")\n")

	// maps
	fmt.Fprintln(out, alpha3map.String())
	fmt.Fprintln(out, "}\n")
	fmt.Fprintln(out, numericMap.String())
	fmt.Fprintln(out, "}\n")
}
