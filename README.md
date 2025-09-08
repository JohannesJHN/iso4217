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

# Examples
```go
package main

import (
    "fmt"
    "github.com/JohannesJHN/iso4217"
)

func main() {
    currency, ok := iso4217.LookupByAlpha3("USD")
    if ok {
      fmt.Println(currency.Name) // United States dollar
      fmt.Println(currency.Numeric) // 840
      fmt.Println(currency.MinorUnits) // 2
      fmt.Println(currency.Countries()) // [AS BQ IO EC SV GU HT MH FM MP PW PA PR TL TC UM US VG VI]
    }

    currency, ok = iso4217.LookupByNumeric(978)
    if ok {
      fmt.Println(currency.Name) // Euro
      fmt.Println(currency.Alpha3) // EUR
      fmt.Println(currency.MinorUnits) // 2
      fmt.Println(currency.Countries()) // [AX AD AT BE HR CY EE FI FR GF TF DE GR GP VA IE IT LV LT LU MT MQ YT MC ME NL PT RE BL MF PM SM SK SI ES]
    }

    currency, ok = iso4217.LookupByCountry("LI")
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
}
```

# Currency struct

A currency has the following attributes:

- `Alpha3 (string)`
  The alpha-3-code of the currency

- `Numeric (int)`
  The numeric code of the currency according to ISO 4217

- `MinorUnits (int)` 
  The number of decimal places used for the currency (e.g., 2 for USD, 0 for JPY)

- `Name (string)`  
  The official name of the currency

- `Countries() []string`  
  Returns a list of ISO 3166-1 alpha-2 country codes where the currency is used

# Get Currency

- `LookupByCode(code string) (Currency, bool)`  
  Retrieves a currency by its ISO 4217 alphabetic (alpha-3) code, such as "USD" or "EUR".  
  Returns the corresponding currency and a boolean indicating whether the lookup was successful.

- `LookupByNumeric(numeric int) (Currency, bool)`  
  Retrieves a currency using its ISO 4217 numeric code.  
  Returns the corresponding currency and a boolean indicating whether the lookup was successful.

- `LookupByCountry(alpha2 string) (Currency, bool)`
  Retrieves a currency by a ISO 3166-1 country code (alpha2), such as "DE" or "US".
  Returns the corresponding currency and a boolean indicating whether the lookup was successful.

- `All() map[string]Currency`  
  Returns a map containing all ISO 4217 currencies, where each key is the currency's alpha-3 code (e.g., "USD", "EUR").  
  This allows easy iteration and access to the complete set of supported currencies.

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
