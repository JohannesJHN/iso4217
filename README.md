# iso4217
currency data for Go
## Overview

`iso4217` provides ISO 4217 currency data for Go. It includes currency codes, names, numeric codes, and minor units, making it easy to work with standardized currency information in your Go applications.

## Features

- Complete ISO 4217 currency list
- Lookup by code or numeric value
- Access to currency name and minor units
- Lightweight and dependency-free

## Installation

```sh
go get github.com/JohannesJHN/iso4217
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/JohannesJHN/iso4217"
)

func main() {
    currency, ok := iso4217.FindByAlpha3("USD")
    if ok {
      fmt.Println(currency.Name) // United States dollar
      fmt.Println(currency.Numeric) // 840
      fmt.Println(currency.MinorUnits) // 2
      fmt.Println(currency.Countries()) // [AS BQ IO EC SV GU HT MH FM MP PW PA PR TL TC UM US VG VI]
    }

    currency, ok = iso4217.FindByNumeric(978)
    if ok {
      fmt.Println(currency.Name) // Euro
      fmt.Println(currency.Alpha3) // EUR
      fmt.Println(currency.MinorUnits) // 2
      fmt.Println(currency.Countries()) // [AX AD AT BE HR CY EE FI FR GF TF DE GR GP VA IE IT LV LT LU MT MQ YT MC ME NL PT RE BL MF PM SM SK SI ES]
    }

    currency, ok = iso4217.FindByCountry("LI")
    if ok {
      fmt.Println(currency.Name) // Swiss franc
      fmt.Println(currency.Alpha3) // CHF
      fmt.Println(currency.Numeric) // 756
      fmt.Println(currency.MinorUnits) // 2
      fmt.Println(currency.Countries()) // [LI CH]
    }

    currency = iso4217.GBP
    if ok {
      fmt.Println(currency.Name) // Pound sterling
      fmt.Println(currency.Numeric) // 826
      fmt.Println(currency.MinorUnits) // 2
      fmt.Println(currency.Countries()) // [GG IM JE GB]
    }

    currency = iso4217.FindByCountry("")
}
```

## Get Currency

- `ByCode(code string) (Currency, bool)`  
  Lookup a currency by its alphabetic code.

- `ByNumeric(numeric int) (Currency, bool)`  
  Lookup a currency by its numeric code.

## Update Currency List
The currency list is based on the official ISO 4217 standard. To update the list:

1. Download the latest ISO 4217 data from the [official ISO website](https://www.iso.org/iso-4217-currency-codes.html).
2. Replace the `generate/iso4217-list.xml` file in the repository with the new data.
3. If necessary, replace the `generate/iso_3166_1.csv` with the latest version.
4. Run the generator script to update the Go source files:
    ```sh
    go run generate/generate.go
    ```

Contributions to keep the currency list up-to-date are welcome!

## License

MIT License. See [LICENSE](LICENSE) for details.
