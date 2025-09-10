package iso4217

type Currency struct {
	Alpha3        string
	Numeric       int
	MinorUnits    int
	Name          string
	WithdrawalDate string // Empty for active currencies
}

var (
	AFN = Currency{Alpha3: "AFN", Numeric: 971, MinorUnits: 2, Name: "Afghani", WithdrawalDate: ""}
	EUR = Currency{Alpha3: "EUR", Numeric: 978, MinorUnits: 2, Name: "Euro", WithdrawalDate: ""}
	ALL = Currency{Alpha3: "ALL", Numeric: 8, MinorUnits: 2, Name: "Lek", WithdrawalDate: ""}
	DZD = Currency{Alpha3: "DZD", Numeric: 12, MinorUnits: 2, Name: "Algerian Dinar", WithdrawalDate: ""}
	USD = Currency{Alpha3: "USD", Numeric: 840, MinorUnits: 2, Name: "US Dollar", WithdrawalDate: ""}
	AOA = Currency{Alpha3: "AOA", Numeric: 973, MinorUnits: 2, Name: "Kwanza", WithdrawalDate: ""}
	XCD = Currency{Alpha3: "XCD", Numeric: 951, MinorUnits: 2, Name: "East Caribbean Dollar", WithdrawalDate: ""}
	XAD = Currency{Alpha3: "XAD", Numeric: 396, MinorUnits: 2, Name: "Arab Accounting Dinar", WithdrawalDate: ""}
	ARS = Currency{Alpha3: "ARS", Numeric: 32, MinorUnits: 2, Name: "Argentine Peso", WithdrawalDate: ""}
	AMD = Currency{Alpha3: "AMD", Numeric: 51, MinorUnits: 2, Name: "Armenian Dram", WithdrawalDate: ""}
	AWG = Currency{Alpha3: "AWG", Numeric: 533, MinorUnits: 2, Name: "Aruban Florin", WithdrawalDate: ""}
	AUD = Currency{Alpha3: "AUD", Numeric: 36, MinorUnits: 2, Name: "Australian Dollar", WithdrawalDate: ""}
	AZN = Currency{Alpha3: "AZN", Numeric: 944, MinorUnits: 2, Name: "Azerbaijan Manat", WithdrawalDate: ""}
	BSD = Currency{Alpha3: "BSD", Numeric: 44, MinorUnits: 2, Name: "Bahamian Dollar", WithdrawalDate: ""}
	BHD = Currency{Alpha3: "BHD", Numeric: 48, MinorUnits: 3, Name: "Bahraini Dinar", WithdrawalDate: ""}
	BDT = Currency{Alpha3: "BDT", Numeric: 50, MinorUnits: 2, Name: "Taka", WithdrawalDate: ""}
	BBD = Currency{Alpha3: "BBD", Numeric: 52, MinorUnits: 2, Name: "Barbados Dollar", WithdrawalDate: ""}
	BYN = Currency{Alpha3: "BYN", Numeric: 933, MinorUnits: 2, Name: "Belarusian Ruble", WithdrawalDate: ""}
	BZD = Currency{Alpha3: "BZD", Numeric: 84, MinorUnits: 2, Name: "Belize Dollar", WithdrawalDate: ""}
	XOF = Currency{Alpha3: "XOF", Numeric: 952, MinorUnits: 0, Name: "CFA Franc BCEAO", WithdrawalDate: ""}
	BMD = Currency{Alpha3: "BMD", Numeric: 60, MinorUnits: 2, Name: "Bermudian Dollar", WithdrawalDate: ""}
	INR = Currency{Alpha3: "INR", Numeric: 356, MinorUnits: 2, Name: "Indian Rupee", WithdrawalDate: ""}
	BTN = Currency{Alpha3: "BTN", Numeric: 64, MinorUnits: 2, Name: "Ngultrum", WithdrawalDate: ""}
	BOB = Currency{Alpha3: "BOB", Numeric: 68, MinorUnits: 2, Name: "Boliviano", WithdrawalDate: ""}
	BOV = Currency{Alpha3: "BOV", Numeric: 984, MinorUnits: 2, Name: "Mvdol", WithdrawalDate: ""}
	BAM = Currency{Alpha3: "BAM", Numeric: 977, MinorUnits: 2, Name: "Convertible Mark", WithdrawalDate: ""}
	BWP = Currency{Alpha3: "BWP", Numeric: 72, MinorUnits: 2, Name: "Pula", WithdrawalDate: ""}
	NOK = Currency{Alpha3: "NOK", Numeric: 578, MinorUnits: 2, Name: "Norwegian Krone", WithdrawalDate: ""}
	BRL = Currency{Alpha3: "BRL", Numeric: 986, MinorUnits: 2, Name: "Brazilian Real", WithdrawalDate: ""}
	BND = Currency{Alpha3: "BND", Numeric: 96, MinorUnits: 2, Name: "Brunei Dollar", WithdrawalDate: ""}
	BGN = Currency{Alpha3: "BGN", Numeric: 975, MinorUnits: 2, Name: "Bulgarian Lev", WithdrawalDate: ""}
	BIF = Currency{Alpha3: "BIF", Numeric: 108, MinorUnits: 0, Name: "Burundi Franc", WithdrawalDate: ""}
	CVE = Currency{Alpha3: "CVE", Numeric: 132, MinorUnits: 2, Name: "Cabo Verde Escudo", WithdrawalDate: ""}
	KHR = Currency{Alpha3: "KHR", Numeric: 116, MinorUnits: 2, Name: "Riel", WithdrawalDate: ""}
	XAF = Currency{Alpha3: "XAF", Numeric: 950, MinorUnits: 0, Name: "CFA Franc BEAC", WithdrawalDate: ""}
	CAD = Currency{Alpha3: "CAD", Numeric: 124, MinorUnits: 2, Name: "Canadian Dollar", WithdrawalDate: ""}
	KYD = Currency{Alpha3: "KYD", Numeric: 136, MinorUnits: 2, Name: "Cayman Islands Dollar", WithdrawalDate: ""}
	CLP = Currency{Alpha3: "CLP", Numeric: 152, MinorUnits: 0, Name: "Chilean Peso", WithdrawalDate: ""}
	CLF = Currency{Alpha3: "CLF", Numeric: 990, MinorUnits: 4, Name: "Unidad de Fomento", WithdrawalDate: ""}
	CNY = Currency{Alpha3: "CNY", Numeric: 156, MinorUnits: 2, Name: "Yuan Renminbi", WithdrawalDate: ""}
	COP = Currency{Alpha3: "COP", Numeric: 170, MinorUnits: 2, Name: "Colombian Peso", WithdrawalDate: ""}
	COU = Currency{Alpha3: "COU", Numeric: 970, MinorUnits: 2, Name: "Unidad de Valor Real", WithdrawalDate: ""}
	KMF = Currency{Alpha3: "KMF", Numeric: 174, MinorUnits: 0, Name: "Comorian Franc ", WithdrawalDate: ""}
	CDF = Currency{Alpha3: "CDF", Numeric: 976, MinorUnits: 2, Name: "Congolese Franc", WithdrawalDate: ""}
	NZD = Currency{Alpha3: "NZD", Numeric: 554, MinorUnits: 2, Name: "New Zealand Dollar", WithdrawalDate: ""}
	CRC = Currency{Alpha3: "CRC", Numeric: 188, MinorUnits: 2, Name: "Costa Rican Colon", WithdrawalDate: ""}
	CUP = Currency{Alpha3: "CUP", Numeric: 192, MinorUnits: 2, Name: "Cuban Peso", WithdrawalDate: ""}
	XCG = Currency{Alpha3: "XCG", Numeric: 532, MinorUnits: 2, Name: "Caribbean Guilder", WithdrawalDate: ""}
	CZK = Currency{Alpha3: "CZK", Numeric: 203, MinorUnits: 2, Name: "Czech Koruna", WithdrawalDate: ""}
	DKK = Currency{Alpha3: "DKK", Numeric: 208, MinorUnits: 2, Name: "Danish Krone", WithdrawalDate: ""}
	DJF = Currency{Alpha3: "DJF", Numeric: 262, MinorUnits: 0, Name: "Djibouti Franc", WithdrawalDate: ""}
	DOP = Currency{Alpha3: "DOP", Numeric: 214, MinorUnits: 2, Name: "Dominican Peso", WithdrawalDate: ""}
	EGP = Currency{Alpha3: "EGP", Numeric: 818, MinorUnits: 2, Name: "Egyptian Pound", WithdrawalDate: ""}
	SVC = Currency{Alpha3: "SVC", Numeric: 222, MinorUnits: 2, Name: "El Salvador Colon", WithdrawalDate: ""}
	ERN = Currency{Alpha3: "ERN", Numeric: 232, MinorUnits: 2, Name: "Nakfa", WithdrawalDate: ""}
	SZL = Currency{Alpha3: "SZL", Numeric: 748, MinorUnits: 2, Name: "Lilangeni", WithdrawalDate: ""}
	ETB = Currency{Alpha3: "ETB", Numeric: 230, MinorUnits: 2, Name: "Ethiopian Birr", WithdrawalDate: ""}
	FKP = Currency{Alpha3: "FKP", Numeric: 238, MinorUnits: 2, Name: "Falkland Islands Pound", WithdrawalDate: ""}
	FJD = Currency{Alpha3: "FJD", Numeric: 242, MinorUnits: 2, Name: "Fiji Dollar", WithdrawalDate: ""}
	XPF = Currency{Alpha3: "XPF", Numeric: 953, MinorUnits: 0, Name: "CFP Franc", WithdrawalDate: ""}
	GMD = Currency{Alpha3: "GMD", Numeric: 270, MinorUnits: 2, Name: "Dalasi", WithdrawalDate: ""}
	GEL = Currency{Alpha3: "GEL", Numeric: 981, MinorUnits: 2, Name: "Lari", WithdrawalDate: ""}
	GHS = Currency{Alpha3: "GHS", Numeric: 936, MinorUnits: 2, Name: "Ghana Cedi", WithdrawalDate: ""}
	GIP = Currency{Alpha3: "GIP", Numeric: 292, MinorUnits: 2, Name: "Gibraltar Pound", WithdrawalDate: ""}
	GTQ = Currency{Alpha3: "GTQ", Numeric: 320, MinorUnits: 2, Name: "Quetzal", WithdrawalDate: ""}
	GBP = Currency{Alpha3: "GBP", Numeric: 826, MinorUnits: 2, Name: "Pound Sterling", WithdrawalDate: ""}
	GNF = Currency{Alpha3: "GNF", Numeric: 324, MinorUnits: 0, Name: "Guinean Franc", WithdrawalDate: ""}
	GYD = Currency{Alpha3: "GYD", Numeric: 328, MinorUnits: 2, Name: "Guyana Dollar", WithdrawalDate: ""}
	HTG = Currency{Alpha3: "HTG", Numeric: 332, MinorUnits: 2, Name: "Gourde", WithdrawalDate: ""}
	HNL = Currency{Alpha3: "HNL", Numeric: 340, MinorUnits: 2, Name: "Lempira", WithdrawalDate: ""}
	HKD = Currency{Alpha3: "HKD", Numeric: 344, MinorUnits: 2, Name: "Hong Kong Dollar", WithdrawalDate: ""}
	HUF = Currency{Alpha3: "HUF", Numeric: 348, MinorUnits: 2, Name: "Forint", WithdrawalDate: ""}
	ISK = Currency{Alpha3: "ISK", Numeric: 352, MinorUnits: 0, Name: "Iceland Krona", WithdrawalDate: ""}
	IDR = Currency{Alpha3: "IDR", Numeric: 360, MinorUnits: 2, Name: "Rupiah", WithdrawalDate: ""}
	XDR = Currency{Alpha3: "XDR", Numeric: 960, MinorUnits: -1, Name: "SDR (Special Drawing Right)", WithdrawalDate: ""}
	IRR = Currency{Alpha3: "IRR", Numeric: 364, MinorUnits: 2, Name: "Iranian Rial", WithdrawalDate: ""}
	IQD = Currency{Alpha3: "IQD", Numeric: 368, MinorUnits: 3, Name: "Iraqi Dinar", WithdrawalDate: ""}
	ILS = Currency{Alpha3: "ILS", Numeric: 376, MinorUnits: 2, Name: "New Israeli Sheqel", WithdrawalDate: ""}
	JMD = Currency{Alpha3: "JMD", Numeric: 388, MinorUnits: 2, Name: "Jamaican Dollar", WithdrawalDate: ""}
	JPY = Currency{Alpha3: "JPY", Numeric: 392, MinorUnits: 0, Name: "Yen", WithdrawalDate: ""}
	JOD = Currency{Alpha3: "JOD", Numeric: 400, MinorUnits: 3, Name: "Jordanian Dinar", WithdrawalDate: ""}
	KZT = Currency{Alpha3: "KZT", Numeric: 398, MinorUnits: 2, Name: "Tenge", WithdrawalDate: ""}
	KES = Currency{Alpha3: "KES", Numeric: 404, MinorUnits: 2, Name: "Kenyan Shilling", WithdrawalDate: ""}
	KPW = Currency{Alpha3: "KPW", Numeric: 408, MinorUnits: 2, Name: "North Korean Won", WithdrawalDate: ""}
	KRW = Currency{Alpha3: "KRW", Numeric: 410, MinorUnits: 0, Name: "Won", WithdrawalDate: ""}
	KWD = Currency{Alpha3: "KWD", Numeric: 414, MinorUnits: 3, Name: "Kuwaiti Dinar", WithdrawalDate: ""}
	KGS = Currency{Alpha3: "KGS", Numeric: 417, MinorUnits: 2, Name: "Som", WithdrawalDate: ""}
	LAK = Currency{Alpha3: "LAK", Numeric: 418, MinorUnits: 2, Name: "Lao Kip", WithdrawalDate: ""}
	LBP = Currency{Alpha3: "LBP", Numeric: 422, MinorUnits: 2, Name: "Lebanese Pound", WithdrawalDate: ""}
	LSL = Currency{Alpha3: "LSL", Numeric: 426, MinorUnits: 2, Name: "Loti", WithdrawalDate: ""}
	ZAR = Currency{Alpha3: "ZAR", Numeric: 710, MinorUnits: 2, Name: "Rand", WithdrawalDate: ""}
	LRD = Currency{Alpha3: "LRD", Numeric: 430, MinorUnits: 2, Name: "Liberian Dollar", WithdrawalDate: ""}
	LYD = Currency{Alpha3: "LYD", Numeric: 434, MinorUnits: 3, Name: "Libyan Dinar", WithdrawalDate: ""}
	CHF = Currency{Alpha3: "CHF", Numeric: 756, MinorUnits: 2, Name: "Swiss Franc", WithdrawalDate: ""}
	MOP = Currency{Alpha3: "MOP", Numeric: 446, MinorUnits: 2, Name: "Pataca", WithdrawalDate: ""}
	MKD = Currency{Alpha3: "MKD", Numeric: 807, MinorUnits: 2, Name: "Denar", WithdrawalDate: ""}
	MGA = Currency{Alpha3: "MGA", Numeric: 969, MinorUnits: 2, Name: "Malagasy Ariary", WithdrawalDate: ""}
	MWK = Currency{Alpha3: "MWK", Numeric: 454, MinorUnits: 2, Name: "Malawi Kwacha", WithdrawalDate: ""}
	MYR = Currency{Alpha3: "MYR", Numeric: 458, MinorUnits: 2, Name: "Malaysian Ringgit", WithdrawalDate: ""}
	MVR = Currency{Alpha3: "MVR", Numeric: 462, MinorUnits: 2, Name: "Rufiyaa", WithdrawalDate: ""}
	MRU = Currency{Alpha3: "MRU", Numeric: 929, MinorUnits: 2, Name: "Ouguiya", WithdrawalDate: ""}
	MUR = Currency{Alpha3: "MUR", Numeric: 480, MinorUnits: 2, Name: "Mauritius Rupee", WithdrawalDate: ""}
	XUA = Currency{Alpha3: "XUA", Numeric: 965, MinorUnits: -1, Name: "ADB Unit of Account", WithdrawalDate: ""}
	MXN = Currency{Alpha3: "MXN", Numeric: 484, MinorUnits: 2, Name: "Mexican Peso", WithdrawalDate: ""}
	MXV = Currency{Alpha3: "MXV", Numeric: 979, MinorUnits: 2, Name: "Mexican Unidad de Inversion (UDI)", WithdrawalDate: ""}
	MDL = Currency{Alpha3: "MDL", Numeric: 498, MinorUnits: 2, Name: "Moldovan Leu", WithdrawalDate: ""}
	MNT = Currency{Alpha3: "MNT", Numeric: 496, MinorUnits: 2, Name: "Tugrik", WithdrawalDate: ""}
	MAD = Currency{Alpha3: "MAD", Numeric: 504, MinorUnits: 2, Name: "Moroccan Dirham", WithdrawalDate: ""}
	MZN = Currency{Alpha3: "MZN", Numeric: 943, MinorUnits: 2, Name: "Mozambique Metical", WithdrawalDate: ""}
	MMK = Currency{Alpha3: "MMK", Numeric: 104, MinorUnits: 2, Name: "Kyat", WithdrawalDate: ""}
	NAD = Currency{Alpha3: "NAD", Numeric: 516, MinorUnits: 2, Name: "Namibia Dollar", WithdrawalDate: ""}
	NPR = Currency{Alpha3: "NPR", Numeric: 524, MinorUnits: 2, Name: "Nepalese Rupee", WithdrawalDate: ""}
	NIO = Currency{Alpha3: "NIO", Numeric: 558, MinorUnits: 2, Name: "Cordoba Oro", WithdrawalDate: ""}
	NGN = Currency{Alpha3: "NGN", Numeric: 566, MinorUnits: 2, Name: "Naira", WithdrawalDate: ""}
	OMR = Currency{Alpha3: "OMR", Numeric: 512, MinorUnits: 3, Name: "Rial Omani", WithdrawalDate: ""}
	PKR = Currency{Alpha3: "PKR", Numeric: 586, MinorUnits: 2, Name: "Pakistan Rupee", WithdrawalDate: ""}
	PAB = Currency{Alpha3: "PAB", Numeric: 590, MinorUnits: 2, Name: "Balboa", WithdrawalDate: ""}
	PGK = Currency{Alpha3: "PGK", Numeric: 598, MinorUnits: 2, Name: "Kina", WithdrawalDate: ""}
	PYG = Currency{Alpha3: "PYG", Numeric: 600, MinorUnits: 0, Name: "Guarani", WithdrawalDate: ""}
	PEN = Currency{Alpha3: "PEN", Numeric: 604, MinorUnits: 2, Name: "Sol", WithdrawalDate: ""}
	PHP = Currency{Alpha3: "PHP", Numeric: 608, MinorUnits: 2, Name: "Philippine Peso", WithdrawalDate: ""}
	PLN = Currency{Alpha3: "PLN", Numeric: 985, MinorUnits: 2, Name: "Zloty", WithdrawalDate: ""}
	QAR = Currency{Alpha3: "QAR", Numeric: 634, MinorUnits: 2, Name: "Qatari Rial", WithdrawalDate: ""}
	RON = Currency{Alpha3: "RON", Numeric: 946, MinorUnits: 2, Name: "Romanian Leu", WithdrawalDate: ""}
	RUB = Currency{Alpha3: "RUB", Numeric: 643, MinorUnits: 2, Name: "Russian Ruble", WithdrawalDate: ""}
	RWF = Currency{Alpha3: "RWF", Numeric: 646, MinorUnits: 0, Name: "Rwanda Franc", WithdrawalDate: ""}
	SHP = Currency{Alpha3: "SHP", Numeric: 654, MinorUnits: 2, Name: "Saint Helena Pound", WithdrawalDate: ""}
	WST = Currency{Alpha3: "WST", Numeric: 882, MinorUnits: 2, Name: "Tala", WithdrawalDate: ""}
	STN = Currency{Alpha3: "STN", Numeric: 930, MinorUnits: 2, Name: "Dobra", WithdrawalDate: ""}
	SAR = Currency{Alpha3: "SAR", Numeric: 682, MinorUnits: 2, Name: "Saudi Riyal", WithdrawalDate: ""}
	RSD = Currency{Alpha3: "RSD", Numeric: 941, MinorUnits: 2, Name: "Serbian Dinar", WithdrawalDate: ""}
	SCR = Currency{Alpha3: "SCR", Numeric: 690, MinorUnits: 2, Name: "Seychelles Rupee", WithdrawalDate: ""}
	SLE = Currency{Alpha3: "SLE", Numeric: 925, MinorUnits: 2, Name: "Leone", WithdrawalDate: ""}
	SGD = Currency{Alpha3: "SGD", Numeric: 702, MinorUnits: 2, Name: "Singapore Dollar", WithdrawalDate: ""}
	XSU = Currency{Alpha3: "XSU", Numeric: 994, MinorUnits: -1, Name: "Sucre", WithdrawalDate: ""}
	SBD = Currency{Alpha3: "SBD", Numeric: 90, MinorUnits: 2, Name: "Solomon Islands Dollar", WithdrawalDate: ""}
	SOS = Currency{Alpha3: "SOS", Numeric: 706, MinorUnits: 2, Name: "Somali Shilling", WithdrawalDate: ""}
	SSP = Currency{Alpha3: "SSP", Numeric: 728, MinorUnits: 2, Name: "South Sudanese Pound", WithdrawalDate: ""}
	LKR = Currency{Alpha3: "LKR", Numeric: 144, MinorUnits: 2, Name: "Sri Lanka Rupee", WithdrawalDate: ""}
	SDG = Currency{Alpha3: "SDG", Numeric: 938, MinorUnits: 2, Name: "Sudanese Pound", WithdrawalDate: ""}
	SRD = Currency{Alpha3: "SRD", Numeric: 968, MinorUnits: 2, Name: "Surinam Dollar", WithdrawalDate: ""}
	SEK = Currency{Alpha3: "SEK", Numeric: 752, MinorUnits: 2, Name: "Swedish Krona", WithdrawalDate: ""}
	CHE = Currency{Alpha3: "CHE", Numeric: 947, MinorUnits: 2, Name: "WIR Euro", WithdrawalDate: ""}
	CHW = Currency{Alpha3: "CHW", Numeric: 948, MinorUnits: 2, Name: "WIR Franc", WithdrawalDate: ""}
	SYP = Currency{Alpha3: "SYP", Numeric: 760, MinorUnits: 2, Name: "Syrian Pound", WithdrawalDate: ""}
	TWD = Currency{Alpha3: "TWD", Numeric: 901, MinorUnits: 2, Name: "New Taiwan Dollar", WithdrawalDate: ""}
	TJS = Currency{Alpha3: "TJS", Numeric: 972, MinorUnits: 2, Name: "Somoni", WithdrawalDate: ""}
	TZS = Currency{Alpha3: "TZS", Numeric: 834, MinorUnits: 2, Name: "Tanzanian Shilling", WithdrawalDate: ""}
	THB = Currency{Alpha3: "THB", Numeric: 764, MinorUnits: 2, Name: "Baht", WithdrawalDate: ""}
	TOP = Currency{Alpha3: "TOP", Numeric: 776, MinorUnits: 2, Name: "Pa’anga", WithdrawalDate: ""}
	TTD = Currency{Alpha3: "TTD", Numeric: 780, MinorUnits: 2, Name: "Trinidad and Tobago Dollar", WithdrawalDate: ""}
	TND = Currency{Alpha3: "TND", Numeric: 788, MinorUnits: 3, Name: "Tunisian Dinar", WithdrawalDate: ""}
	TRY = Currency{Alpha3: "TRY", Numeric: 949, MinorUnits: 2, Name: "Turkish Lira", WithdrawalDate: ""}
	TMT = Currency{Alpha3: "TMT", Numeric: 934, MinorUnits: 2, Name: "Turkmenistan New Manat", WithdrawalDate: ""}
	UGX = Currency{Alpha3: "UGX", Numeric: 800, MinorUnits: 0, Name: "Uganda Shilling", WithdrawalDate: ""}
	UAH = Currency{Alpha3: "UAH", Numeric: 980, MinorUnits: 2, Name: "Hryvnia", WithdrawalDate: ""}
	AED = Currency{Alpha3: "AED", Numeric: 784, MinorUnits: 2, Name: "UAE Dirham", WithdrawalDate: ""}
	USN = Currency{Alpha3: "USN", Numeric: 997, MinorUnits: 2, Name: "US Dollar (Next day)", WithdrawalDate: ""}
	UYU = Currency{Alpha3: "UYU", Numeric: 858, MinorUnits: 2, Name: "Peso Uruguayo", WithdrawalDate: ""}
	UYI = Currency{Alpha3: "UYI", Numeric: 940, MinorUnits: 0, Name: "Uruguay Peso en Unidades Indexadas (UI)", WithdrawalDate: ""}
	UYW = Currency{Alpha3: "UYW", Numeric: 927, MinorUnits: 4, Name: "Unidad Previsional", WithdrawalDate: ""}
	UZS = Currency{Alpha3: "UZS", Numeric: 860, MinorUnits: 2, Name: "Uzbekistan Sum", WithdrawalDate: ""}
	VUV = Currency{Alpha3: "VUV", Numeric: 548, MinorUnits: 0, Name: "Vatu", WithdrawalDate: ""}
	VES = Currency{Alpha3: "VES", Numeric: 928, MinorUnits: 2, Name: "Bolívar Soberano", WithdrawalDate: ""}
	VED = Currency{Alpha3: "VED", Numeric: 926, MinorUnits: 2, Name: "Bolívar Soberano", WithdrawalDate: ""}
	VND = Currency{Alpha3: "VND", Numeric: 704, MinorUnits: 0, Name: "Dong", WithdrawalDate: ""}
	YER = Currency{Alpha3: "YER", Numeric: 886, MinorUnits: 2, Name: "Yemeni Rial", WithdrawalDate: ""}
	ZMW = Currency{Alpha3: "ZMW", Numeric: 967, MinorUnits: 2, Name: "Zambian Kwacha", WithdrawalDate: ""}
	ZWG = Currency{Alpha3: "ZWG", Numeric: 924, MinorUnits: 2, Name: "Zimbabwe Gold", WithdrawalDate: ""}
	XBA = Currency{Alpha3: "XBA", Numeric: 955, MinorUnits: -1, Name: "Bond Markets Unit European Composite Unit (EURCO)", WithdrawalDate: ""}
	XBB = Currency{Alpha3: "XBB", Numeric: 956, MinorUnits: -1, Name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)", WithdrawalDate: ""}
	XBC = Currency{Alpha3: "XBC", Numeric: 957, MinorUnits: -1, Name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", WithdrawalDate: ""}
	XBD = Currency{Alpha3: "XBD", Numeric: 958, MinorUnits: -1, Name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", WithdrawalDate: ""}
	XTS = Currency{Alpha3: "XTS", Numeric: 963, MinorUnits: -1, Name: "Codes specifically reserved for testing purposes", WithdrawalDate: ""}
	XXX = Currency{Alpha3: "XXX", Numeric: 999, MinorUnits: -1, Name: "The codes assigned for transactions where no currency is involved", WithdrawalDate: ""}
	XAU = Currency{Alpha3: "XAU", Numeric: 959, MinorUnits: -1, Name: "Gold", WithdrawalDate: ""}
	XPD = Currency{Alpha3: "XPD", Numeric: 964, MinorUnits: -1, Name: "Palladium", WithdrawalDate: ""}
	XPT = Currency{Alpha3: "XPT", Numeric: 962, MinorUnits: -1, Name: "Platinum", WithdrawalDate: ""}
	XAG = Currency{Alpha3: "XAG", Numeric: 961, MinorUnits: -1, Name: "Silver", WithdrawalDate: ""}
	AFA = Currency{Alpha3: "AFA", Numeric: 4, MinorUnits: -1, Name: "Afghani", WithdrawalDate: "2003-01"}
	FIM = Currency{Alpha3: "FIM", Numeric: 246, MinorUnits: -1, Name: "Markka", WithdrawalDate: "2002-03"}
	ALK = Currency{Alpha3: "ALK", Numeric: 8, MinorUnits: -1, Name: "Old Lek", WithdrawalDate: "1989-12"}
	ADP = Currency{Alpha3: "ADP", Numeric: 20, MinorUnits: -1, Name: "Andorran Peseta", WithdrawalDate: "2003-07"}
	ESP = Currency{Alpha3: "ESP", Numeric: 724, MinorUnits: -1, Name: "Spanish Peseta", WithdrawalDate: "2002-03"}
	FRF = Currency{Alpha3: "FRF", Numeric: 250, MinorUnits: -1, Name: "French Franc", WithdrawalDate: "2002-03"}
	AOK = Currency{Alpha3: "AOK", Numeric: 24, MinorUnits: -1, Name: "Kwanza", WithdrawalDate: "1991-03"}
	AON = Currency{Alpha3: "AON", Numeric: 24, MinorUnits: -1, Name: "New Kwanza", WithdrawalDate: "2000-02"}
	AOR = Currency{Alpha3: "AOR", Numeric: 982, MinorUnits: -1, Name: "Kwanza Reajustado", WithdrawalDate: "2000-02"}
	ARA = Currency{Alpha3: "ARA", Numeric: 32, MinorUnits: -1, Name: "Austral", WithdrawalDate: "1992-01"}
	ARP = Currency{Alpha3: "ARP", Numeric: 32, MinorUnits: -1, Name: "Peso Argentino", WithdrawalDate: "1985-07"}
	ARY = Currency{Alpha3: "ARY", Numeric: 32, MinorUnits: -1, Name: "Peso", WithdrawalDate: "1989 to 1990"}
	RUR = Currency{Alpha3: "RUR", Numeric: 810, MinorUnits: -1, Name: "Russian Ruble", WithdrawalDate: "1994-08"}
	ATS = Currency{Alpha3: "ATS", Numeric: 40, MinorUnits: -1, Name: "Schilling", WithdrawalDate: "2002-03"}
	AYM = Currency{Alpha3: "AYM", Numeric: 945, MinorUnits: -1, Name: "Azerbaijan Manat", WithdrawalDate: "2005-10"}
	AZM = Currency{Alpha3: "AZM", Numeric: 31, MinorUnits: -1, Name: "Azerbaijanian Manat", WithdrawalDate: "2005-12"}
	BYB = Currency{Alpha3: "BYB", Numeric: 112, MinorUnits: -1, Name: "Belarusian Ruble", WithdrawalDate: "2001-01"}
	BYR = Currency{Alpha3: "BYR", Numeric: 974, MinorUnits: -1, Name: "Belarusian Ruble", WithdrawalDate: "2017-01"}
	BEC = Currency{Alpha3: "BEC", Numeric: 993, MinorUnits: -1, Name: "Convertible Franc", WithdrawalDate: "1990-03"}
	BEF = Currency{Alpha3: "BEF", Numeric: 56, MinorUnits: -1, Name: "Belgian Franc", WithdrawalDate: "2002-03"}
	BEL = Currency{Alpha3: "BEL", Numeric: 992, MinorUnits: -1, Name: "Financial Franc", WithdrawalDate: "1990-03"}
	BOP = Currency{Alpha3: "BOP", Numeric: 68, MinorUnits: -1, Name: "Peso boliviano", WithdrawalDate: "1987-02"}
	BAD = Currency{Alpha3: "BAD", Numeric: 70, MinorUnits: -1, Name: "Dinar", WithdrawalDate: "1998-07"}
	BRB = Currency{Alpha3: "BRB", Numeric: 76, MinorUnits: -1, Name: "Cruzeiro", WithdrawalDate: "1986-03"}
	BRC = Currency{Alpha3: "BRC", Numeric: 76, MinorUnits: -1, Name: "Cruzado", WithdrawalDate: "1989-02"}
	BRE = Currency{Alpha3: "BRE", Numeric: 76, MinorUnits: -1, Name: "Cruzeiro", WithdrawalDate: "1993-03"}
	BRN = Currency{Alpha3: "BRN", Numeric: 76, MinorUnits: -1, Name: "New Cruzado", WithdrawalDate: "1990-03"}
	BRR = Currency{Alpha3: "BRR", Numeric: 987, MinorUnits: -1, Name: "Cruzeiro Real", WithdrawalDate: "1994-07"}
	BGJ = Currency{Alpha3: "BGJ", Numeric: 100, MinorUnits: -1, Name: "Lev A/52", WithdrawalDate: "1989 to 1990"}
	BGK = Currency{Alpha3: "BGK", Numeric: 100, MinorUnits: -1, Name: "Lev A/62", WithdrawalDate: "1989 to 1990"}
	BGL = Currency{Alpha3: "BGL", Numeric: 100, MinorUnits: -1, Name: "Lev", WithdrawalDate: "2003-11"}
	BUK = Currency{Alpha3: "BUK", Numeric: 104, MinorUnits: -1, Name: "Kyat", WithdrawalDate: "1990-02"}
	HRD = Currency{Alpha3: "HRD", Numeric: 191, MinorUnits: -1, Name: "Croatian Dinar", WithdrawalDate: "1995-01"}
	HRK = Currency{Alpha3: "HRK", Numeric: 191, MinorUnits: -1, Name: "Croatian Kuna", WithdrawalDate: "2015-06"}
	CUC = Currency{Alpha3: "CUC", Numeric: 931, MinorUnits: -1, Name: "Peso Convertible", WithdrawalDate: "2021-06"}
	ANG = Currency{Alpha3: "ANG", Numeric: 532, MinorUnits: -1, Name: "Netherlands Antillean Guilder", WithdrawalDate: "2025-03"}
	CYP = Currency{Alpha3: "CYP", Numeric: 196, MinorUnits: -1, Name: "Cyprus Pound", WithdrawalDate: "2008-01"}
	CSJ = Currency{Alpha3: "CSJ", Numeric: 203, MinorUnits: -1, Name: "Krona A/53", WithdrawalDate: "1989 to 1990"}
	CSK = Currency{Alpha3: "CSK", Numeric: 200, MinorUnits: -1, Name: "Koruna", WithdrawalDate: "1993-03"}
	ECS = Currency{Alpha3: "ECS", Numeric: 218, MinorUnits: -1, Name: "Sucre", WithdrawalDate: "2000-09"}
	ECV = Currency{Alpha3: "ECV", Numeric: 983, MinorUnits: -1, Name: "Unidad de Valor Constante (UVC)", WithdrawalDate: "2000-09"}
	GQE = Currency{Alpha3: "GQE", Numeric: 226, MinorUnits: -1, Name: "Ekwele", WithdrawalDate: "1986-06"}
	EEK = Currency{Alpha3: "EEK", Numeric: 233, MinorUnits: -1, Name: "Kroon", WithdrawalDate: "2011-01"}
	XEU = Currency{Alpha3: "XEU", Numeric: 954, MinorUnits: -1, Name: "European Currency Unit (E.C.U)", WithdrawalDate: "1999-01"}
	GEK = Currency{Alpha3: "GEK", Numeric: 268, MinorUnits: -1, Name: "Georgian Coupon", WithdrawalDate: "1995-10"}
	DDM = Currency{Alpha3: "DDM", Numeric: 278, MinorUnits: -1, Name: "Mark der DDR", WithdrawalDate: "1990-07 to 1990-09"}
	DEM = Currency{Alpha3: "DEM", Numeric: 276, MinorUnits: -1, Name: "Deutsche Mark", WithdrawalDate: "2002-03"}
	GHC = Currency{Alpha3: "GHC", Numeric: 288, MinorUnits: -1, Name: "Cedi", WithdrawalDate: "2008-01"}
	GHP = Currency{Alpha3: "GHP", Numeric: 939, MinorUnits: -1, Name: "Ghana Cedi", WithdrawalDate: "2007-06"}
	GRD = Currency{Alpha3: "GRD", Numeric: 300, MinorUnits: -1, Name: "Drachma", WithdrawalDate: "2002-03"}
	GNE = Currency{Alpha3: "GNE", Numeric: 324, MinorUnits: -1, Name: "Syli", WithdrawalDate: "1989-12"}
	GNS = Currency{Alpha3: "GNS", Numeric: 324, MinorUnits: -1, Name: "Syli", WithdrawalDate: "1986-02"}
	GWE = Currency{Alpha3: "GWE", Numeric: 624, MinorUnits: -1, Name: "Guinea Escudo", WithdrawalDate: "1978 to 1981"}
	GWP = Currency{Alpha3: "GWP", Numeric: 624, MinorUnits: -1, Name: "Guinea-Bissau Peso", WithdrawalDate: "1997-05"}
	ITL = Currency{Alpha3: "ITL", Numeric: 380, MinorUnits: -1, Name: "Italian Lira", WithdrawalDate: "2002-03"}
	ISJ = Currency{Alpha3: "ISJ", Numeric: 352, MinorUnits: -1, Name: "Old Krona", WithdrawalDate: "1989 to 1990"}
	IEP = Currency{Alpha3: "IEP", Numeric: 372, MinorUnits: -1, Name: "Irish Pound", WithdrawalDate: "2002-03"}
	ILP = Currency{Alpha3: "ILP", Numeric: 376, MinorUnits: -1, Name: "Pound", WithdrawalDate: "1978 to 1981"}
	ILR = Currency{Alpha3: "ILR", Numeric: 376, MinorUnits: -1, Name: "Old Shekel", WithdrawalDate: "1989 to 1990"}
	LAJ = Currency{Alpha3: "LAJ", Numeric: 418, MinorUnits: -1, Name: "Pathet Lao Kip", WithdrawalDate: "1979-12"}
	LVL = Currency{Alpha3: "LVL", Numeric: 428, MinorUnits: -1, Name: "Latvian Lats", WithdrawalDate: "2014-01"}
	LVR = Currency{Alpha3: "LVR", Numeric: 428, MinorUnits: -1, Name: "Latvian Ruble", WithdrawalDate: "1994-12"}
	LSM = Currency{Alpha3: "LSM", Numeric: 426, MinorUnits: -1, Name: "Loti", WithdrawalDate: "1985-05"}
	ZAL = Currency{Alpha3: "ZAL", Numeric: 991, MinorUnits: -1, Name: "Financial Rand", WithdrawalDate: "1995-03"}
	LTL = Currency{Alpha3: "LTL", Numeric: 440, MinorUnits: -1, Name: "Lithuanian Litas", WithdrawalDate: "2014-12"}
	LTT = Currency{Alpha3: "LTT", Numeric: 440, MinorUnits: -1, Name: "Talonas", WithdrawalDate: "1993-07"}
	LUC = Currency{Alpha3: "LUC", Numeric: 989, MinorUnits: -1, Name: "Luxembourg Convertible Franc", WithdrawalDate: "1990-03"}
	LUF = Currency{Alpha3: "LUF", Numeric: 442, MinorUnits: -1, Name: "Luxembourg Franc", WithdrawalDate: "2002-03"}
	LUL = Currency{Alpha3: "LUL", Numeric: 988, MinorUnits: -1, Name: "Luxembourg Financial Franc", WithdrawalDate: "1990-03"}
	MGF = Currency{Alpha3: "MGF", Numeric: 450, MinorUnits: -1, Name: "Malagasy Franc", WithdrawalDate: "2004-12"}
	MVQ = Currency{Alpha3: "MVQ", Numeric: 462, MinorUnits: -1, Name: "Maldive Rupee", WithdrawalDate: "1989-12"}
	MLF = Currency{Alpha3: "MLF", Numeric: 466, MinorUnits: -1, Name: "Mali Franc", WithdrawalDate: "1984-11"}
	MTL = Currency{Alpha3: "MTL", Numeric: 470, MinorUnits: -1, Name: "Maltese Lira", WithdrawalDate: "2008-01"}
	MTP = Currency{Alpha3: "MTP", Numeric: 470, MinorUnits: -1, Name: "Maltese Pound", WithdrawalDate: "1983-06"}
	MRO = Currency{Alpha3: "MRO", Numeric: 478, MinorUnits: -1, Name: "Ouguiya", WithdrawalDate: "2017-12"}
	MXP = Currency{Alpha3: "MXP", Numeric: 484, MinorUnits: -1, Name: "Mexican Peso", WithdrawalDate: "1993-01"}
	MZE = Currency{Alpha3: "MZE", Numeric: 508, MinorUnits: -1, Name: "Mozambique Escudo", WithdrawalDate: "1978 to 1981"}
	MZM = Currency{Alpha3: "MZM", Numeric: 508, MinorUnits: -1, Name: "Mozambique Metical", WithdrawalDate: "2006-06"}
	NLG = Currency{Alpha3: "NLG", Numeric: 528, MinorUnits: -1, Name: "Netherlands Guilder", WithdrawalDate: "2002-03"}
	NIC = Currency{Alpha3: "NIC", Numeric: 558, MinorUnits: -1, Name: "Cordoba", WithdrawalDate: "1990-10"}
	PEH = Currency{Alpha3: "PEH", Numeric: 604, MinorUnits: -1, Name: "Sol", WithdrawalDate: "1989 to 1990"}
	PEI = Currency{Alpha3: "PEI", Numeric: 604, MinorUnits: -1, Name: "Inti", WithdrawalDate: "1991-07"}
	PES = Currency{Alpha3: "PES", Numeric: 604, MinorUnits: -1, Name: "Sol", WithdrawalDate: "1986-02"}
	PLZ = Currency{Alpha3: "PLZ", Numeric: 616, MinorUnits: -1, Name: "Zloty", WithdrawalDate: "1997-01"}
	PTE = Currency{Alpha3: "PTE", Numeric: 620, MinorUnits: -1, Name: "Portuguese Escudo", WithdrawalDate: "2002-03"}
	ROK = Currency{Alpha3: "ROK", Numeric: 642, MinorUnits: -1, Name: "Leu A/52", WithdrawalDate: "1989 to 1990"}
	ROL = Currency{Alpha3: "ROL", Numeric: 642, MinorUnits: -1, Name: "Old Leu", WithdrawalDate: "2005-06"}
	STD = Currency{Alpha3: "STD", Numeric: 678, MinorUnits: -1, Name: "Dobra", WithdrawalDate: "2017-12"}
	CSD = Currency{Alpha3: "CSD", Numeric: 891, MinorUnits: -1, Name: "Serbian Dinar", WithdrawalDate: "2006-10"}
	SLL = Currency{Alpha3: "SLL", Numeric: 694, MinorUnits: -1, Name: "Leone", WithdrawalDate: "2023-12"}
	SKK = Currency{Alpha3: "SKK", Numeric: 703, MinorUnits: -1, Name: "Slovak Koruna", WithdrawalDate: "2009-01"}
	SIT = Currency{Alpha3: "SIT", Numeric: 705, MinorUnits: -1, Name: "Tolar", WithdrawalDate: "2007-01"}
	RHD = Currency{Alpha3: "RHD", Numeric: 716, MinorUnits: -1, Name: "Rhodesian Dollar", WithdrawalDate: "1978 to 1981"}
	ESA = Currency{Alpha3: "ESA", Numeric: 996, MinorUnits: -1, Name: "Spanish Peseta", WithdrawalDate: "1978 to 1981"}
	ESB = Currency{Alpha3: "ESB", Numeric: 995, MinorUnits: -1, Name: "\"A\" Account (convertible Peseta Account)", WithdrawalDate: "1994-12"}
	SDD = Currency{Alpha3: "SDD", Numeric: 736, MinorUnits: -1, Name: "Sudanese Dinar", WithdrawalDate: "2007-07"}
	SDP = Currency{Alpha3: "SDP", Numeric: 736, MinorUnits: -1, Name: "Sudanese Pound", WithdrawalDate: "1998-06"}
	SRG = Currency{Alpha3: "SRG", Numeric: 740, MinorUnits: -1, Name: "Surinam Guilder", WithdrawalDate: "2003-12"}
	CHC = Currency{Alpha3: "CHC", Numeric: 948, MinorUnits: -1, Name: "WIR Franc (for electronic)", WithdrawalDate: "2004-11"}
	TJR = Currency{Alpha3: "TJR", Numeric: 762, MinorUnits: -1, Name: "Tajik Ruble", WithdrawalDate: "2001-04"}
	TPE = Currency{Alpha3: "TPE", Numeric: 626, MinorUnits: -1, Name: "Timor Escudo", WithdrawalDate: "2002-11"}
	TRL = Currency{Alpha3: "TRL", Numeric: 792, MinorUnits: -1, Name: "Old Turkish Lira", WithdrawalDate: "2005-12"}
	TMM = Currency{Alpha3: "TMM", Numeric: 795, MinorUnits: -1, Name: "Turkmenistan Manat", WithdrawalDate: "2009-01"}
	UGS = Currency{Alpha3: "UGS", Numeric: 800, MinorUnits: -1, Name: "Uganda Shilling", WithdrawalDate: "1987-05"}
	UGW = Currency{Alpha3: "UGW", Numeric: 800, MinorUnits: -1, Name: "Old Shilling", WithdrawalDate: "1989 to 1990"}
	UAK = Currency{Alpha3: "UAK", Numeric: 804, MinorUnits: -1, Name: "Karbovanet", WithdrawalDate: "1996-09"}
	SUR = Currency{Alpha3: "SUR", Numeric: 810, MinorUnits: -1, Name: "Rouble", WithdrawalDate: "1990-12"}
	USS = Currency{Alpha3: "USS", Numeric: 998, MinorUnits: -1, Name: "US Dollar (Same day)", WithdrawalDate: "2014-03"}
	UYN = Currency{Alpha3: "UYN", Numeric: 858, MinorUnits: -1, Name: "Old Uruguay Peso", WithdrawalDate: "1989-12"}
	UYP = Currency{Alpha3: "UYP", Numeric: 858, MinorUnits: -1, Name: "Uruguayan Peso", WithdrawalDate: "1993-03"}
	VEB = Currency{Alpha3: "VEB", Numeric: 862, MinorUnits: -1, Name: "Bolivar", WithdrawalDate: "2008-01"}
	VEF = Currency{Alpha3: "VEF", Numeric: 937, MinorUnits: -1, Name: "Bolivar Fuerte", WithdrawalDate: "2011-12"}
	VNC = Currency{Alpha3: "VNC", Numeric: 704, MinorUnits: -1, Name: "Old Dong", WithdrawalDate: "1989-1990"}
	YDD = Currency{Alpha3: "YDD", Numeric: 720, MinorUnits: -1, Name: "Yemeni Dinar", WithdrawalDate: "1991-09"}
	YUD = Currency{Alpha3: "YUD", Numeric: 890, MinorUnits: -1, Name: "New Yugoslavian Dinar", WithdrawalDate: "1990-01"}
	YUM = Currency{Alpha3: "YUM", Numeric: 891, MinorUnits: -1, Name: "New Dinar", WithdrawalDate: "2003-07"}
	YUN = Currency{Alpha3: "YUN", Numeric: 890, MinorUnits: -1, Name: "Yugoslavian Dinar", WithdrawalDate: "1995-11"}
	ZRN = Currency{Alpha3: "ZRN", Numeric: 180, MinorUnits: -1, Name: "New Zaire", WithdrawalDate: "1999-06"}
	ZRZ = Currency{Alpha3: "ZRZ", Numeric: 180, MinorUnits: -1, Name: "Zaire", WithdrawalDate: "1994-02"}
	ZMK = Currency{Alpha3: "ZMK", Numeric: 894, MinorUnits: -1, Name: "Zambian Kwacha", WithdrawalDate: "2012-12"}
	ZWC = Currency{Alpha3: "ZWC", Numeric: 716, MinorUnits: -1, Name: "Rhodesian Dollar", WithdrawalDate: "1989-12"}
	ZWD = Currency{Alpha3: "ZWD", Numeric: 716, MinorUnits: -1, Name: "Zimbabwe Dollar (old)", WithdrawalDate: "2006-08"}
	ZWN = Currency{Alpha3: "ZWN", Numeric: 942, MinorUnits: -1, Name: "Zimbabwe Dollar (new)", WithdrawalDate: "2006-09"}
	ZWR = Currency{Alpha3: "ZWR", Numeric: 935, MinorUnits: -1, Name: "Zimbabwe Dollar", WithdrawalDate: "2009-06"}
	ZWL = Currency{Alpha3: "ZWL", Numeric: 932, MinorUnits: -1, Name: "Zimbabwe\u00a0Dollar", WithdrawalDate: "2024-09"}
	XFO = Currency{Alpha3: "XFO", Numeric: 0, MinorUnits: -1, Name: "Gold-Franc", WithdrawalDate: "2006-10"}
	XRE = Currency{Alpha3: "XRE", Numeric: 0, MinorUnits: -1, Name: "RINET Funds Code", WithdrawalDate: "1999-11"}
	XFU = Currency{Alpha3: "XFU", Numeric: 0, MinorUnits: -1, Name: "UIC-Franc", WithdrawalDate: "2013-11"}
)

var currenciesByNumeric = map[int]Currency{
	971: AFN,
	978: EUR,
	8: ALL,
	12: DZD,
	840: USD,
	973: AOA,
	951: XCD,
	396: XAD,
	32: ARS,
	51: AMD,
	533: AWG,
	36: AUD,
	944: AZN,
	44: BSD,
	48: BHD,
	50: BDT,
	52: BBD,
	933: BYN,
	84: BZD,
	952: XOF,
	60: BMD,
	356: INR,
	64: BTN,
	68: BOB,
	984: BOV,
	977: BAM,
	72: BWP,
	578: NOK,
	986: BRL,
	96: BND,
	975: BGN,
	108: BIF,
	132: CVE,
	116: KHR,
	950: XAF,
	124: CAD,
	136: KYD,
	152: CLP,
	990: CLF,
	156: CNY,
	170: COP,
	970: COU,
	174: KMF,
	976: CDF,
	554: NZD,
	188: CRC,
	192: CUP,
	532: XCG,
	203: CZK,
	208: DKK,
	262: DJF,
	214: DOP,
	818: EGP,
	222: SVC,
	232: ERN,
	748: SZL,
	230: ETB,
	238: FKP,
	242: FJD,
	953: XPF,
	270: GMD,
	981: GEL,
	936: GHS,
	292: GIP,
	320: GTQ,
	826: GBP,
	324: GNF,
	328: GYD,
	332: HTG,
	340: HNL,
	344: HKD,
	348: HUF,
	352: ISK,
	360: IDR,
	960: XDR,
	364: IRR,
	368: IQD,
	376: ILS,
	388: JMD,
	392: JPY,
	400: JOD,
	398: KZT,
	404: KES,
	408: KPW,
	410: KRW,
	414: KWD,
	417: KGS,
	418: LAK,
	422: LBP,
	426: LSL,
	710: ZAR,
	430: LRD,
	434: LYD,
	756: CHF,
	446: MOP,
	807: MKD,
	969: MGA,
	454: MWK,
	458: MYR,
	462: MVR,
	929: MRU,
	480: MUR,
	965: XUA,
	484: MXN,
	979: MXV,
	498: MDL,
	496: MNT,
	504: MAD,
	943: MZN,
	104: MMK,
	516: NAD,
	524: NPR,
	558: NIO,
	566: NGN,
	512: OMR,
	586: PKR,
	590: PAB,
	598: PGK,
	600: PYG,
	604: PEN,
	608: PHP,
	985: PLN,
	634: QAR,
	946: RON,
	643: RUB,
	646: RWF,
	654: SHP,
	882: WST,
	930: STN,
	682: SAR,
	941: RSD,
	690: SCR,
	925: SLE,
	702: SGD,
	994: XSU,
	90: SBD,
	706: SOS,
	728: SSP,
	144: LKR,
	938: SDG,
	968: SRD,
	752: SEK,
	947: CHE,
	948: CHW,
	760: SYP,
	901: TWD,
	972: TJS,
	834: TZS,
	764: THB,
	776: TOP,
	780: TTD,
	788: TND,
	949: TRY,
	934: TMT,
	800: UGX,
	980: UAH,
	784: AED,
	997: USN,
	858: UYU,
	940: UYI,
	927: UYW,
	860: UZS,
	548: VUV,
	928: VES,
	926: VED,
	704: VND,
	886: YER,
	967: ZMW,
	924: ZWG,
	955: XBA,
	956: XBB,
	957: XBC,
	958: XBD,
	963: XTS,
	999: XXX,
	959: XAU,
	964: XPD,
	962: XPT,
	961: XAG,
}

var currenciesByCountry = map[string]Currency{
	"IQ": IQD,
	"MA": MAD,
	"KN": XCD,
	"AE": AED,
	"DO": DOP,
	"EC": USD,
	"PT": EUR,
	"EH": MAD,
	"AZ": AZN,
	"CN": CNY,
	"CU": CUP,
	"SZ": SZL,
	"NZ": NZD,
	"QA": QAR,
	"AO": AOA,
	"AU": AUD,
	"BD": BDT,
	"CM": XAF,
	"KZ": KZT,
	"MP": USD,
	"NO": NOK,
	"AF": AFN,
	"KY": KYD,
	"HM": AUD,
	"GL": DKK,
	"IR": IRR,
	"EE": EUR,
	"GA": XAF,
	"SJ": NOK,
	"SV": USD,
	"IN": INR,
	"TC": USD,
	"LI": CHF,
	"MK": MKD,
	"MU": MUR,
	"VE": VED,
	"AI": XCD,
	"MV": MVR,
	"AX": EUR,
	"BT": BTN,
	"FO": DKK,
	"PF": XPF,
	"KG": KGS,
	"IS": ISK,
	"ID": IDR,
	"LY": LYD,
	"MH": USD,
	"ES": EUR,
	"IT": EUR,
	"JM": JMD,
	"SH": SHP,
	"SD": SDG,
	"VU": VUV,
	"BW": BWP,
	"DE": EUR,
	"GT": GTQ,
	"KW": KWD,
	"WS": WST,
	"LK": LKR,
	"TM": TMT,
	"DJ": DJF,
	"FK": FKP,
	"BI": BIF,
	"GU": USD,
	"GY": GYD,
	"MZ": MZN,
	"WF": XPF,
	"AM": AMD,
	"SN": XOF,
	"RS": RSD,
	"AS": USD,
	"BQ": USD,
	"GF": EUR,
	"GP": EUR,
	"PE": PEN,
	"SK": EUR,
	"TN": TND,
	"GE": GEL,
	"PH": PHP,
	"SS": SSP,
	"ZM": ZMW,
	"KP": KPW,
	"TH": THB,
	"VN": VND,
	"KM": KMF,
	"HU": HUF,
	"CL": CLP,
	"LR": LRD,
	"MM": MMK,
	"BB": BBD,
	"CG": XAF,
	"PG": PGK,
	"TV": AUD,
	"IL": ILS,
	"MN": MNT,
	"AR": ARS,
	"PN": NZD,
	"SA": SAR,
	"SO": SOS,
	"BY": BYN,
	"ML": XOF,
	"RO": RON,
	"AT": EUR,
	"PR": USD,
	"SY": SYP,
	"GB": GBP,
	"VI": USD,
	"BA": BAM,
	"CW": XCG,
	"ET": ETB,
	"FI": EUR,
	"LV": EUR,
	"TG": XOF,
	"CV": CVE,
	"NU": NZD,
	"LC": XCD,
	"BZ": BZD,
	"HR": EUR,
	"OM": OMR,
	"MD": MDL,
	"NL": EUR,
	"BV": NOK,
	"HN": HNL,
	"MG": MGA,
	"RW": RWF,
	"IM": GBP,
	"MT": EUR,
	"SC": SCR,
	"TZ": TZS,
	"UG": UGX,
	"CD": CDF,
	"DK": DKK,
	"MQ": EUR,
	"SX": XCG,
	"BM": BMD,
	"TD": XAF,
	"CI": XOF,
	"GH": GHS,
	"LT": EUR,
	"ME": EUR,
	"BL": EUR,
	"IE": EUR,
	"LU": EUR,
	"MY": MYR,
	"MC": EUR,
	"NI": NIO,
	"BF": XOF,
	"FJ": FJD,
	"HT": USD,
	"CH": CHF,
	"AW": AWG,
	"MX": MXN,
	"GG": GBP,
	"VA": EUR,
	"LA": LAK,
	"PL": PLN,
	"TL": USD,
	"CX": AUD,
	"NR": AUD,
	"EG": EGP,
	"FR": EUR,
	"NG": NGN,
	"UZ": UZS,
	"CY": EUR,
	"ER": ERN,
	"MO": MOP,
	"MR": MRU,
	"RU": RUB,
	"UM": USD,
	"CC": AUD,
	"MS": XCD,
	"VG": USD,
	"BH": BHD,
	"GI": GIP,
	"KE": KES,
	"LB": LBP,
	"US": USD,
	"GM": GMD,
	"GR": EUR,
	"MW": MWK,
	"VC": XCD,
	"ZA": ZAR,
	"UY": UYW,
	"ZW": ZWG,
	"DZ": DZD,
	"AD": EUR,
	"BE": EUR,
	"KR": KRW,
	"NE": XOF,
	"SR": SRD,
	"CR": CRC,
	"TF": EUR,
	"JO": JOD,
	"SI": EUR,
	"GD": XCD,
	"KI": AUD,
	"FM": USD,
	"TJ": TJS,
	"AG": XCD,
	"BS": BSD,
	"BG": BGN,
	"CZ": CZK,
	"GW": XOF,
	"NP": NPR,
	"PW": USD,
	"MF": EUR,
	"TW": TWD,
	"TR": TRY,
	"IO": USD,
	"HK": HKD,
	"NA": ZAR,
	"DM": XCD,
	"ST": STN,
	"AL": ALL,
	"NF": AUD,
	"RE": EUR,
	"PM": EUR,
	"SB": SBD,
	"YE": YER,
	"BN": BND,
	"LS": ZAR,
	"NC": XPF,
	"TT": TTD,
	"UA": UAH,
	"GN": GNF,
	"JP": JPY,
	"YT": EUR,
	"PY": PYG,
	"SE": SEK,
	"BO": BOB,
	"CF": XAF,
	"KH": KHR,
	"PK": PKR,
	"BJ": XOF,
	"BR": BRL,
	"JE": GBP,
	"TK": NZD,
	"TO": TOP,
	"CA": CAD,
	"CK": NZD,
	"GQ": XAF,
	"SL": SLE,
	"SG": SGD,
	"CO": COP,
	"PA": USD,
	"SM": EUR,
}

var countryList = map[Currency][]string{
	HNL: []string{"HN"},
	IQD: []string{"IQ"},
	JOD: []string{"JO"},
	VES: []string{"VE"},
	CVE: []string{"CV"},
	HKD: []string{"HK"},
	EUR: []string{"AX", "AD", "AT", "BE", "HR", "CY", "EE", "FI", "FR", "GF", "TF", "DE", "GR", "GP", "VA", "IE", "IT", "LV", "LT", "LU", "MT", "MQ", "YT", "MC", "ME", "NL", "PT", "RE", "BL", "MF", "PM", "SM", "SK", "SI", "ES"},
	AZN: []string{"AZ"},
	XOF: []string{"BJ", "BF", "CI", "GW", "ML", "NE", "SN", "TG"},
	BWP: []string{"BW"},
	TJS: []string{"TJ"},
	CNY: []string{"CN"},
	MKD: []string{"MK"},
	PEN: []string{"PE"},
	RUB: []string{"RU"},
	SYP: []string{"SY"},
	ETB: []string{"ET"},
	LSL: []string{"LS"},
	SCR: []string{"SC"},
	LKR: []string{"LK"},
	NIO: []string{"NI"},
	THB: []string{"TH"},
	BGN: []string{"BG"},
	DKK: []string{"DK", "FO", "GL"},
	FJD: []string{"FJ"},
	GYD: []string{"GY"},
	BYN: []string{"BY"},
	ERN: []string{"ER"},
	ISK: []string{"IS"},
	IRR: []string{"IR"},
	ZWG: []string{"ZW"},
	USD: []string{"AS", "BQ", "IO", "EC", "SV", "GU", "HT", "MH", "FM", "MP", "PW", "PA", "PR", "TL", "TC", "UM", "US", "VG", "VI"},
	XAF: []string{"CM", "CF", "TD", "CG", "GQ", "GA"},
	NZD: []string{"CK", "NZ", "NU", "PN", "TK"},
	ZMW: []string{"ZM"},
	SSP: []string{"SS"},
	TOP: []string{"TO"},
	VND: []string{"VN"},
	DZD: []string{"DZ"},
	MDL: []string{"MD"},
	NPR: []string{"NP"},
	PKR: []string{"PK"},
	MNT: []string{"MN"},
	QAR: []string{"QA"},
	SLE: []string{"SL"},
	ARS: []string{"AR"},
	CUP: []string{"CU"},
	SVC: []string{"SV"},
	KES: []string{"KE"},
	CAD: []string{"CA"},
	DOP: []string{"DO"},
	SEK: []string{"SE"},
	UYU: []string{"UY"},
	MVR: []string{"MV"},
	PHP: []string{"PH"},
	SRD: []string{"SR"},
	TTD: []string{"TT"},
	MXN: []string{"MX"},
	MZN: []string{"MZ"},
	SHP: []string{"SH"},
	SOS: []string{"SO"},
	BAM: []string{"BA"},
	LAK: []string{"LA"},
	LBP: []string{"LB"},
	LYD: []string{"LY"},
	CZK: []string{"CZ"},
	SZL: []string{"SZ"},
	KPW: []string{"KP"},
	MOP: []string{"MO"},
	CHF: []string{"LI", "CH"},
	RON: []string{"RO"},
	TZS: []string{"TZ"},
	BTN: []string{"BT"},
	BIF: []string{"BI"},
	CLP: []string{"CL"},
	CRC: []string{"CR"},
	MRU: []string{"MR"},
	SDG: []string{"SD"},
	TMT: []string{"TM"},
	VED: []string{"VE"},
	BMD: []string{"BM"},
	XPF: []string{"PF", "NC", "WF"},
	GBP: []string{"GG", "IM", "JE", "GB"},
	ZAR: []string{"LS", "NA", "ZA"},
	EGP: []string{"EG"},
	TRY: []string{"TR"},
	NGN: []string{"NG"},
	OMR: []string{"OM"},
	UYW: []string{"UY"},
	AUD: []string{"AU", "CX", "CC", "HM", "KI", "NR", "NF", "TV"},
	NOK: []string{"BV", "NO", "SJ"},
	GEL: []string{"GE"},
	HTG: []string{"HT"},
	BHD: []string{"BH"},
	KGS: []string{"KG"},
	AED: []string{"AE"},
	VUV: []string{"VU"},
	ALL: []string{"AL"},
	GIP: []string{"GI"},
	PAB: []string{"PA"},
	STN: []string{"ST"},
	SBD: []string{"SB"},
	UAH: []string{"UA"},
	UZS: []string{"UZ"},
	KHR: []string{"KH"},
	KYD: []string{"KY"},
	CDF: []string{"CD"},
	GHS: []string{"GH"},
	INR: []string{"BT", "IN"},
	KZT: []string{"KZ"},
	BND: []string{"BN"},
	DJF: []string{"DJ"},
	FKP: []string{"FK"},
	MGA: []string{"MG"},
	NAD: []string{"NA"},
	RWF: []string{"RW"},
	KWD: []string{"KW"},
	UGX: []string{"UG"},
	YER: []string{"YE"},
	XCD: []string{"AI", "AG", "DM", "GD", "MS", "KN", "LC", "VC"},
	GTQ: []string{"GT"},
	ILS: []string{"IL"},
	JMD: []string{"JM"},
	RSD: []string{"RS"},
	AFN: []string{"AF"},
	AMD: []string{"AM"},
	BBD: []string{"BB"},
	KMF: []string{"KM"},
	BZD: []string{"BZ"},
	KRW: []string{"KR"},
	SAR: []string{"SA"},
	SGD: []string{"SG"},
	PGK: []string{"PG"},
	GMD: []string{"GM"},
	IDR: []string{"ID"},
	JPY: []string{"JP"},
	MYR: []string{"MY"},
	WST: []string{"WS"},
	TWD: []string{"TW"},
	MAD: []string{"MA", "EH"},
	MMK: []string{"MM"},
	PYG: []string{"PY"},
	AOA: []string{"AO"},
	AWG: []string{"AW"},
	BSD: []string{"BS"},
	XCG: []string{"CW", "SX"},
	PLN: []string{"PL"},
	TND: []string{"TN"},
	BRL: []string{"BR"},
	COP: []string{"CO"},
	HUF: []string{"HU"},
	MUR: []string{"MU"},
	BDT: []string{"BD"},
	GNF: []string{"GN"},
	MWK: []string{"MW"},
	BOB: []string{"BO"},
	LRD: []string{"LR"},
}

var activeCurrencies = map[string]Currency{
	"KWD": KWD,
	"UYI": UYI,
	"ZWG": ZWG,
	"XUA": XUA,
	"XSU": XSU,
	"USN": USN,
	"XBA": XBA,
	"INR": INR,
	"NZD": NZD,
	"KES": KES,
	"PLN": PLN,
	"RUB": RUB,
	"VUV": VUV,
	"CNY": CNY,
	"GEL": GEL,
	"KGS": KGS,
	"TTD": TTD,
	"UYU": UYU,
	"ZMW": ZMW,
	"CAD": CAD,
	"GNF": GNF,
	"LRD": LRD,
	"MXN": MXN,
	"MAD": MAD,
	"SGD": SGD,
	"USD": USD,
	"BHD": BHD,
	"MZN": MZN,
	"NGN": NGN,
	"PEN": PEN,
	"AUD": AUD,
	"CZK": CZK,
	"JMD": JMD,
	"ZAR": ZAR,
	"MNT": MNT,
	"MMK": MMK,
	"SRD": SRD,
	"QAR": QAR,
	"UAH": UAH,
	"AMD": AMD,
	"BMD": BMD,
	"GMD": GMD,
	"GHS": GHS,
	"GIP": GIP,
	"LBP": LBP,
	"AFN": AFN,
	"BDT": BDT,
	"GTQ": GTQ,
	"TND": TND,
	"XBB": XBB,
	"DZD": DZD,
	"BOV": BOV,
	"CVE": CVE,
	"HTG": HTG,
	"RWF": RWF,
	"STN": STN,
	"GBP": GBP,
	"IDR": IDR,
	"KZT": KZT,
	"TZS": TZS,
	"MWK": MWK,
	"PKR": PKR,
	"TJS": TJS,
	"UYW": UYW,
	"BSD": BSD,
	"BND": BND,
	"XCG": XCG,
	"DJF": DJF,
	"CHW": CHW,
	"YER": YER,
	"MOP": MOP,
	"AOA": AOA,
	"BTN": BTN,
	"BGN": BGN,
	"KMF": KMF,
	"CUP": CUP,
	"IQD": IQD,
	"PYG": PYG,
	"SBD": SBD,
	"BYN": BYN,
	"DKK": DKK,
	"ISK": ISK,
	"KRW": KRW,
	"MXV": MXV,
	"PAB": PAB,
	"XBC": XBC,
	"ARS": ARS,
	"AZN": AZN,
	"CHF": CHF,
	"SDG": SDG,
	"THB": THB,
	"TRY": TRY,
	"ALL": ALL,
	"NOK": NOK,
	"SEK": SEK,
	"VES": VES,
	"WST": WST,
	"SSP": SSP,
	"VED": VED,
	"XAD": XAD,
	"BOB": BOB,
	"KHR": KHR,
	"COU": COU,
	"MVR": MVR,
	"NAD": NAD,
	"SHP": SHP,
	"CHE": CHE,
	"AWG": AWG,
	"BBD": BBD,
	"CRC": CRC,
	"GYD": GYD,
	"JOD": JOD,
	"OMR": OMR,
	"SYP": SYP,
	"TWD": TWD,
	"XBD": XBD,
	"NPR": NPR,
	"NIO": NIO,
	"EUR": EUR,
	"XOF": XOF,
	"BAM": BAM,
	"FJD": FJD,
	"XDR": XDR,
	"MRU": MRU,
	"SAR": SAR,
	"XTS": XTS,
	"DOP": DOP,
	"SOS": SOS,
	"UZS": UZS,
	"XAU": XAU,
	"XPD": XPD,
	"FKP": FKP,
	"MKD": MKD,
	"MGA": MGA,
	"MUR": MUR,
	"PHP": PHP,
	"RON": RON,
	"HKD": HKD,
	"AED": AED,
	"COP": COP,
	"HNL": HNL,
	"HUF": HUF,
	"LSL": LSL,
	"SCR": SCR,
	"VND": VND,
	"CLP": CLP,
	"ERN": ERN,
	"TMT": TMT,
	"UGX": UGX,
	"XAG": XAG,
	"CLF": CLF,
	"IRR": IRR,
	"LAK": LAK,
	"MDL": MDL,
	"TOP": TOP,
	"XPT": XPT,
	"SLE": SLE,
	"LKR": LKR,
	"BZD": BZD,
	"BIF": BIF,
	"KYD": KYD,
	"XPF": XPF,
	"JPY": JPY,
	"LYD": LYD,
	"XCD": XCD,
	"XAF": XAF,
	"ILS": ILS,
	"KPW": KPW,
	"RSD": RSD,
	"SZL": SZL,
	"PGK": PGK,
	"BWP": BWP,
	"CDF": CDF,
	"SVC": SVC,
	"ETB": ETB,
	"XXX": XXX,
	"BRL": BRL,
	"EGP": EGP,
	"MYR": MYR,
}

var inactiveCurrencies = map[string]Currency{
	"BEC": BEC,
	"CSK": CSK,
	"LVL": LVL,
	"BYB": BYB,
	"SLL": SLL,
	"AOK": AOK,
	"CSD": CSD,
	"ZRN": ZRN,
	"ZWC": ZWC,
	"BRN": BRN,
	"CUC": CUC,
	"GWE": GWE,
	"MTP": MTP,
	"UYN": UYN,
	"ESP": ESP,
	"FRF": FRF,
	"ECS": ECS,
	"GQE": GQE,
	"PES": PES,
	"YUD": YUD,
	"LTL": LTL,
	"LUC": LUC,
	"PEI": PEI,
	"ESB": ESB,
	"VNC": VNC,
	"ATS": ATS,
	"BRE": BRE,
	"BGJ": BGJ,
	"CSJ": CSJ,
	"UYP": UYP,
	"HRK": HRK,
	"GRD": GRD,
	"MRO": MRO,
	"VEF": VEF,
	"ZWL": ZWL,
	"CYP": CYP,
	"ITL": ITL,
	"ILR": ILR,
	"AON": AON,
	"LVR": LVR,
	"ZAL": ZAL,
	"MZE": MZE,
	"SKK": SKK,
	"TRL": TRL,
	"VEB": VEB,
	"ALK": ALK,
	"ARY": ARY,
	"GWP": GWP,
	"MGF": MGF,
	"ROK": ROK,
	"BRC": BRC,
	"HRD": HRD,
	"LUF": LUF,
	"MXP": MXP,
	"TMM": TMM,
	"MLF": MLF,
	"MTL": MTL,
	"BAD": BAD,
	"GNS": GNS,
	"YDD": YDD,
	"CHC": CHC,
	"TJR": TJR,
	"AOR": AOR,
	"ARA": ARA,
	"XEU": XEU,
	"LAJ": LAJ,
	"MVQ": MVQ,
	"SDP": SDP,
	"ZMK": ZMK,
	"RUR": RUR,
	"BGK": BGK,
	"YUM": YUM,
	"ARP": ARP,
	"RHD": RHD,
	"SDD": SDD,
	"MZM": MZM,
	"ADP": ADP,
	"BUK": BUK,
	"NLG": NLG,
	"UGS": UGS,
	"GHC": GHC,
	"LUL": LUL,
	"SIT": SIT,
	"USS": USS,
	"ZWN": ZWN,
	"BEL": BEL,
	"DDM": DDM,
	"IEP": IEP,
	"PTE": PTE,
	"TPE": TPE,
	"ILP": ILP,
	"ZRZ": ZRZ,
	"STD": STD,
	"UAK": UAK,
	"ISJ": ISJ,
	"ZWD": ZWD,
	"BEF": BEF,
	"BGL": BGL,
	"LTT": LTT,
	"PEH": PEH,
	"ESA": ESA,
	"SUR": SUR,
	"AYM": AYM,
	"BRB": BRB,
	"BRR": BRR,
	"EEK": EEK,
	"XFO": XFO,
	"SRG": SRG,
	"UGW": UGW,
	"ZWR": ZWR,
	"AZM": AZM,
	"ANG": ANG,
	"DEM": DEM,
	"GHP": GHP,
	"LSM": LSM,
	"GEK": GEK,
	"ROL": ROL,
	"XFU": XFU,
	"BYR": BYR,
	"GNE": GNE,
	"PLZ": PLZ,
	"YUN": YUN,
	"XRE": XRE,
	"AFA": AFA,
	"FIM": FIM,
	"BOP": BOP,
	"ECV": ECV,
	"NIC": NIC,
}

