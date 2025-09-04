package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ISO4217 struct {
	CcyTbl struct {
		CcyNtry []struct {
			CcyNm  string `xml:"CcyNm"`
			Ccy    string `xml:"Ccy"`
			CcyNbr string `xml:"CcyNbr"`
			CcyMnr string `xml:"CcyMnrUnts"`
		} `xml:"CcyNtry"`
	} `xml:"CcyTbl"`
}

func main() {
	xmlFile := "src/iso4217-list.xml"
	data, err := os.ReadFile(xmlFile)
	if err != nil {
		panic(err)
	}

	var iso ISO4217
	if err := xml.Unmarshal(data, &iso); err != nil {
		panic(err)
	}

	seen := make(map[string]bool)
	var out strings.Builder

	out.WriteString("package iso4217\n\n")
	out.WriteString("type Currency struct {\n")
	out.WriteString("\tAlpha      string\n")
	out.WriteString("\tNumeric    int16\n")
	out.WriteString("\tMinorUnits int8\n")
	out.WriteString("\tName       string\n")
	out.WriteString("}\n\n")

	out.WriteString("var CurrenciesByAlpha = map[string]Currency{\n")
	var numericMap strings.Builder
	numericMap.WriteString("var CurrenciesByNumeric = map[int16]Currency{\n")

	for _, entry := range iso.CcyTbl.CcyNtry {
		if entry.Ccy == "" || seen[entry.Ccy] {
			continue
		}
		seen[entry.Ccy] = true

		var numeric int16 = -1
		if n, err := strconv.Atoi(entry.CcyNbr); err == nil {
			numeric = int16(n)
		}

		var minor int8 = -1
		if n, err := strconv.Atoi(entry.CcyMnr); err == nil {
			minor = int8(n)
		}

		line := fmt.Sprintf("{Alpha: \"%s\", Numeric: %d, MinorUnits: %d, Name: %q}", entry.Ccy, numeric, minor, entry.CcyNm)

		out.WriteString(fmt.Sprintf("\t\"%s\": %s,\n", entry.Ccy, line))
		numericMap.WriteString(fmt.Sprintf("\t%d: %s,\n", numeric, line))
	}

	out.WriteString("}\n\n")
	numericMap.WriteString("}\n")

	final := out.String() + numericMap.String()

	if err := os.WriteFile("./currencies.go", []byte(final), 0644); err != nil {
		panic(err)
	}
}
