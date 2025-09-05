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
	xmlFile := "generate/iso4217-list.xml"
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
	out.WriteString("\tNumeric    int\n")
	out.WriteString("\tMinorUnits int\n")
	out.WriteString("\tName       string\n")
	out.WriteString("}\n\n")

	out.WriteString("var currenciesByAlpha3 = map[string]Currency{\n")
	var numericMap strings.Builder
	numericMap.WriteString("var currenciesByNumeric = map[int]Currency{\n")

	for _, entry := range iso.CcyTbl.CcyNtry {
		if entry.Ccy == "" || seen[entry.Ccy] {
			continue
		}
		seen[entry.Ccy] = true

		var numeric int = -1
		if n, err := strconv.Atoi(entry.CcyNbr); err == nil {
			numeric = int(n)
		}

		var minor int = -1
		if n, err := strconv.Atoi(entry.CcyMnr); err == nil {
			minor = int(n)
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
