package iso4217

type Currency struct {
	Alpha3     string
	Numeric    int
	MinorUnits int
	Name       string
}

var (
	AFN = Currency{Alpha3: "AFN", Numeric: 971, MinorUnits: 2, Name: "Afghani"}
	EUR = Currency{Alpha3: "EUR", Numeric: 978, MinorUnits: 2, Name: "Euro"}
	ALL = Currency{Alpha3: "ALL", Numeric: 8, MinorUnits: 2, Name: "Lek"}
	DZD = Currency{Alpha3: "DZD", Numeric: 12, MinorUnits: 2, Name: "Algerian Dinar"}
	USD = Currency{Alpha3: "USD", Numeric: 840, MinorUnits: 2, Name: "US Dollar"}
	AOA = Currency{Alpha3: "AOA", Numeric: 973, MinorUnits: 2, Name: "Kwanza"}
	XCD = Currency{Alpha3: "XCD", Numeric: 951, MinorUnits: 2, Name: "East Caribbean Dollar"}
	XAD = Currency{Alpha3: "XAD", Numeric: 396, MinorUnits: 2, Name: "Arab Accounting Dinar"}
	ARS = Currency{Alpha3: "ARS", Numeric: 32, MinorUnits: 2, Name: "Argentine Peso"}
	AMD = Currency{Alpha3: "AMD", Numeric: 51, MinorUnits: 2, Name: "Armenian Dram"}
	AWG = Currency{Alpha3: "AWG", Numeric: 533, MinorUnits: 2, Name: "Aruban Florin"}
	AUD = Currency{Alpha3: "AUD", Numeric: 36, MinorUnits: 2, Name: "Australian Dollar"}
	AZN = Currency{Alpha3: "AZN", Numeric: 944, MinorUnits: 2, Name: "Azerbaijan Manat"}
	BSD = Currency{Alpha3: "BSD", Numeric: 44, MinorUnits: 2, Name: "Bahamian Dollar"}
	BHD = Currency{Alpha3: "BHD", Numeric: 48, MinorUnits: 3, Name: "Bahraini Dinar"}
	BDT = Currency{Alpha3: "BDT", Numeric: 50, MinorUnits: 2, Name: "Taka"}
	BBD = Currency{Alpha3: "BBD", Numeric: 52, MinorUnits: 2, Name: "Barbados Dollar"}
	BYN = Currency{Alpha3: "BYN", Numeric: 933, MinorUnits: 2, Name: "Belarusian Ruble"}
	BZD = Currency{Alpha3: "BZD", Numeric: 84, MinorUnits: 2, Name: "Belize Dollar"}
	XOF = Currency{Alpha3: "XOF", Numeric: 952, MinorUnits: 0, Name: "CFA Franc BCEAO"}
	BMD = Currency{Alpha3: "BMD", Numeric: 60, MinorUnits: 2, Name: "Bermudian Dollar"}
	INR = Currency{Alpha3: "INR", Numeric: 356, MinorUnits: 2, Name: "Indian Rupee"}
	BTN = Currency{Alpha3: "BTN", Numeric: 64, MinorUnits: 2, Name: "Ngultrum"}
	BOB = Currency{Alpha3: "BOB", Numeric: 68, MinorUnits: 2, Name: "Boliviano"}
	BOV = Currency{Alpha3: "BOV", Numeric: 984, MinorUnits: 2, Name: "Mvdol"}
	BAM = Currency{Alpha3: "BAM", Numeric: 977, MinorUnits: 2, Name: "Convertible Mark"}
	BWP = Currency{Alpha3: "BWP", Numeric: 72, MinorUnits: 2, Name: "Pula"}
	NOK = Currency{Alpha3: "NOK", Numeric: 578, MinorUnits: 2, Name: "Norwegian Krone"}
	BRL = Currency{Alpha3: "BRL", Numeric: 986, MinorUnits: 2, Name: "Brazilian Real"}
	BND = Currency{Alpha3: "BND", Numeric: 96, MinorUnits: 2, Name: "Brunei Dollar"}
	BGN = Currency{Alpha3: "BGN", Numeric: 975, MinorUnits: 2, Name: "Bulgarian Lev"}
	BIF = Currency{Alpha3: "BIF", Numeric: 108, MinorUnits: 0, Name: "Burundi Franc"}
	CVE = Currency{Alpha3: "CVE", Numeric: 132, MinorUnits: 2, Name: "Cabo Verde Escudo"}
	KHR = Currency{Alpha3: "KHR", Numeric: 116, MinorUnits: 2, Name: "Riel"}
	XAF = Currency{Alpha3: "XAF", Numeric: 950, MinorUnits: 0, Name: "CFA Franc BEAC"}
	CAD = Currency{Alpha3: "CAD", Numeric: 124, MinorUnits: 2, Name: "Canadian Dollar"}
	KYD = Currency{Alpha3: "KYD", Numeric: 136, MinorUnits: 2, Name: "Cayman Islands Dollar"}
	CLP = Currency{Alpha3: "CLP", Numeric: 152, MinorUnits: 0, Name: "Chilean Peso"}
	CLF = Currency{Alpha3: "CLF", Numeric: 990, MinorUnits: 4, Name: "Unidad de Fomento"}
	CNY = Currency{Alpha3: "CNY", Numeric: 156, MinorUnits: 2, Name: "Yuan Renminbi"}
	COP = Currency{Alpha3: "COP", Numeric: 170, MinorUnits: 2, Name: "Colombian Peso"}
	COU = Currency{Alpha3: "COU", Numeric: 970, MinorUnits: 2, Name: "Unidad de Valor Real"}
	KMF = Currency{Alpha3: "KMF", Numeric: 174, MinorUnits: 0, Name: "Comorian Franc "}
	CDF = Currency{Alpha3: "CDF", Numeric: 976, MinorUnits: 2, Name: "Congolese Franc"}
	NZD = Currency{Alpha3: "NZD", Numeric: 554, MinorUnits: 2, Name: "New Zealand Dollar"}
	CRC = Currency{Alpha3: "CRC", Numeric: 188, MinorUnits: 2, Name: "Costa Rican Colon"}
	CUP = Currency{Alpha3: "CUP", Numeric: 192, MinorUnits: 2, Name: "Cuban Peso"}
	XCG = Currency{Alpha3: "XCG", Numeric: 532, MinorUnits: 2, Name: "Caribbean Guilder"}
	CZK = Currency{Alpha3: "CZK", Numeric: 203, MinorUnits: 2, Name: "Czech Koruna"}
	DKK = Currency{Alpha3: "DKK", Numeric: 208, MinorUnits: 2, Name: "Danish Krone"}
	DJF = Currency{Alpha3: "DJF", Numeric: 262, MinorUnits: 0, Name: "Djibouti Franc"}
	DOP = Currency{Alpha3: "DOP", Numeric: 214, MinorUnits: 2, Name: "Dominican Peso"}
	EGP = Currency{Alpha3: "EGP", Numeric: 818, MinorUnits: 2, Name: "Egyptian Pound"}
	SVC = Currency{Alpha3: "SVC", Numeric: 222, MinorUnits: 2, Name: "El Salvador Colon"}
	ERN = Currency{Alpha3: "ERN", Numeric: 232, MinorUnits: 2, Name: "Nakfa"}
	SZL = Currency{Alpha3: "SZL", Numeric: 748, MinorUnits: 2, Name: "Lilangeni"}
	ETB = Currency{Alpha3: "ETB", Numeric: 230, MinorUnits: 2, Name: "Ethiopian Birr"}
	FKP = Currency{Alpha3: "FKP", Numeric: 238, MinorUnits: 2, Name: "Falkland Islands Pound"}
	FJD = Currency{Alpha3: "FJD", Numeric: 242, MinorUnits: 2, Name: "Fiji Dollar"}
	XPF = Currency{Alpha3: "XPF", Numeric: 953, MinorUnits: 0, Name: "CFP Franc"}
	GMD = Currency{Alpha3: "GMD", Numeric: 270, MinorUnits: 2, Name: "Dalasi"}
	GEL = Currency{Alpha3: "GEL", Numeric: 981, MinorUnits: 2, Name: "Lari"}
	GHS = Currency{Alpha3: "GHS", Numeric: 936, MinorUnits: 2, Name: "Ghana Cedi"}
	GIP = Currency{Alpha3: "GIP", Numeric: 292, MinorUnits: 2, Name: "Gibraltar Pound"}
	GTQ = Currency{Alpha3: "GTQ", Numeric: 320, MinorUnits: 2, Name: "Quetzal"}
	GBP = Currency{Alpha3: "GBP", Numeric: 826, MinorUnits: 2, Name: "Pound Sterling"}
	GNF = Currency{Alpha3: "GNF", Numeric: 324, MinorUnits: 0, Name: "Guinean Franc"}
	GYD = Currency{Alpha3: "GYD", Numeric: 328, MinorUnits: 2, Name: "Guyana Dollar"}
	HTG = Currency{Alpha3: "HTG", Numeric: 332, MinorUnits: 2, Name: "Gourde"}
	HNL = Currency{Alpha3: "HNL", Numeric: 340, MinorUnits: 2, Name: "Lempira"}
	HKD = Currency{Alpha3: "HKD", Numeric: 344, MinorUnits: 2, Name: "Hong Kong Dollar"}
	HUF = Currency{Alpha3: "HUF", Numeric: 348, MinorUnits: 2, Name: "Forint"}
	ISK = Currency{Alpha3: "ISK", Numeric: 352, MinorUnits: 0, Name: "Iceland Krona"}
	IDR = Currency{Alpha3: "IDR", Numeric: 360, MinorUnits: 2, Name: "Rupiah"}
	XDR = Currency{Alpha3: "XDR", Numeric: 960, MinorUnits: -1, Name: "SDR (Special Drawing Right)"}
	IRR = Currency{Alpha3: "IRR", Numeric: 364, MinorUnits: 2, Name: "Iranian Rial"}
	IQD = Currency{Alpha3: "IQD", Numeric: 368, MinorUnits: 3, Name: "Iraqi Dinar"}
	ILS = Currency{Alpha3: "ILS", Numeric: 376, MinorUnits: 2, Name: "New Israeli Sheqel"}
	JMD = Currency{Alpha3: "JMD", Numeric: 388, MinorUnits: 2, Name: "Jamaican Dollar"}
	JPY = Currency{Alpha3: "JPY", Numeric: 392, MinorUnits: 0, Name: "Yen"}
	JOD = Currency{Alpha3: "JOD", Numeric: 400, MinorUnits: 3, Name: "Jordanian Dinar"}
	KZT = Currency{Alpha3: "KZT", Numeric: 398, MinorUnits: 2, Name: "Tenge"}
	KES = Currency{Alpha3: "KES", Numeric: 404, MinorUnits: 2, Name: "Kenyan Shilling"}
	KPW = Currency{Alpha3: "KPW", Numeric: 408, MinorUnits: 2, Name: "North Korean Won"}
	KRW = Currency{Alpha3: "KRW", Numeric: 410, MinorUnits: 0, Name: "Won"}
	KWD = Currency{Alpha3: "KWD", Numeric: 414, MinorUnits: 3, Name: "Kuwaiti Dinar"}
	KGS = Currency{Alpha3: "KGS", Numeric: 417, MinorUnits: 2, Name: "Som"}
	LAK = Currency{Alpha3: "LAK", Numeric: 418, MinorUnits: 2, Name: "Lao Kip"}
	LBP = Currency{Alpha3: "LBP", Numeric: 422, MinorUnits: 2, Name: "Lebanese Pound"}
	LSL = Currency{Alpha3: "LSL", Numeric: 426, MinorUnits: 2, Name: "Loti"}
	ZAR = Currency{Alpha3: "ZAR", Numeric: 710, MinorUnits: 2, Name: "Rand"}
	LRD = Currency{Alpha3: "LRD", Numeric: 430, MinorUnits: 2, Name: "Liberian Dollar"}
	LYD = Currency{Alpha3: "LYD", Numeric: 434, MinorUnits: 3, Name: "Libyan Dinar"}
	CHF = Currency{Alpha3: "CHF", Numeric: 756, MinorUnits: 2, Name: "Swiss Franc"}
	MOP = Currency{Alpha3: "MOP", Numeric: 446, MinorUnits: 2, Name: "Pataca"}
	MKD = Currency{Alpha3: "MKD", Numeric: 807, MinorUnits: 2, Name: "Denar"}
	MGA = Currency{Alpha3: "MGA", Numeric: 969, MinorUnits: 2, Name: "Malagasy Ariary"}
	MWK = Currency{Alpha3: "MWK", Numeric: 454, MinorUnits: 2, Name: "Malawi Kwacha"}
	MYR = Currency{Alpha3: "MYR", Numeric: 458, MinorUnits: 2, Name: "Malaysian Ringgit"}
	MVR = Currency{Alpha3: "MVR", Numeric: 462, MinorUnits: 2, Name: "Rufiyaa"}
	MRU = Currency{Alpha3: "MRU", Numeric: 929, MinorUnits: 2, Name: "Ouguiya"}
	MUR = Currency{Alpha3: "MUR", Numeric: 480, MinorUnits: 2, Name: "Mauritius Rupee"}
	XUA = Currency{Alpha3: "XUA", Numeric: 965, MinorUnits: -1, Name: "ADB Unit of Account"}
	MXN = Currency{Alpha3: "MXN", Numeric: 484, MinorUnits: 2, Name: "Mexican Peso"}
	MXV = Currency{Alpha3: "MXV", Numeric: 979, MinorUnits: 2, Name: "Mexican Unidad de Inversion (UDI)"}
	MDL = Currency{Alpha3: "MDL", Numeric: 498, MinorUnits: 2, Name: "Moldovan Leu"}
	MNT = Currency{Alpha3: "MNT", Numeric: 496, MinorUnits: 2, Name: "Tugrik"}
	MAD = Currency{Alpha3: "MAD", Numeric: 504, MinorUnits: 2, Name: "Moroccan Dirham"}
	MZN = Currency{Alpha3: "MZN", Numeric: 943, MinorUnits: 2, Name: "Mozambique Metical"}
	MMK = Currency{Alpha3: "MMK", Numeric: 104, MinorUnits: 2, Name: "Kyat"}
	NAD = Currency{Alpha3: "NAD", Numeric: 516, MinorUnits: 2, Name: "Namibia Dollar"}
	NPR = Currency{Alpha3: "NPR", Numeric: 524, MinorUnits: 2, Name: "Nepalese Rupee"}
	NIO = Currency{Alpha3: "NIO", Numeric: 558, MinorUnits: 2, Name: "Cordoba Oro"}
	NGN = Currency{Alpha3: "NGN", Numeric: 566, MinorUnits: 2, Name: "Naira"}
	OMR = Currency{Alpha3: "OMR", Numeric: 512, MinorUnits: 3, Name: "Rial Omani"}
	PKR = Currency{Alpha3: "PKR", Numeric: 586, MinorUnits: 2, Name: "Pakistan Rupee"}
	PAB = Currency{Alpha3: "PAB", Numeric: 590, MinorUnits: 2, Name: "Balboa"}
	PGK = Currency{Alpha3: "PGK", Numeric: 598, MinorUnits: 2, Name: "Kina"}
	PYG = Currency{Alpha3: "PYG", Numeric: 600, MinorUnits: 0, Name: "Guarani"}
	PEN = Currency{Alpha3: "PEN", Numeric: 604, MinorUnits: 2, Name: "Sol"}
	PHP = Currency{Alpha3: "PHP", Numeric: 608, MinorUnits: 2, Name: "Philippine Peso"}
	PLN = Currency{Alpha3: "PLN", Numeric: 985, MinorUnits: 2, Name: "Zloty"}
	QAR = Currency{Alpha3: "QAR", Numeric: 634, MinorUnits: 2, Name: "Qatari Rial"}
	RON = Currency{Alpha3: "RON", Numeric: 946, MinorUnits: 2, Name: "Romanian Leu"}
	RUB = Currency{Alpha3: "RUB", Numeric: 643, MinorUnits: 2, Name: "Russian Ruble"}
	RWF = Currency{Alpha3: "RWF", Numeric: 646, MinorUnits: 0, Name: "Rwanda Franc"}
	SHP = Currency{Alpha3: "SHP", Numeric: 654, MinorUnits: 2, Name: "Saint Helena Pound"}
	WST = Currency{Alpha3: "WST", Numeric: 882, MinorUnits: 2, Name: "Tala"}
	STN = Currency{Alpha3: "STN", Numeric: 930, MinorUnits: 2, Name: "Dobra"}
	SAR = Currency{Alpha3: "SAR", Numeric: 682, MinorUnits: 2, Name: "Saudi Riyal"}
	RSD = Currency{Alpha3: "RSD", Numeric: 941, MinorUnits: 2, Name: "Serbian Dinar"}
	SCR = Currency{Alpha3: "SCR", Numeric: 690, MinorUnits: 2, Name: "Seychelles Rupee"}
	SLE = Currency{Alpha3: "SLE", Numeric: 925, MinorUnits: 2, Name: "Leone"}
	SGD = Currency{Alpha3: "SGD", Numeric: 702, MinorUnits: 2, Name: "Singapore Dollar"}
	XSU = Currency{Alpha3: "XSU", Numeric: 994, MinorUnits: -1, Name: "Sucre"}
	SBD = Currency{Alpha3: "SBD", Numeric: 90, MinorUnits: 2, Name: "Solomon Islands Dollar"}
	SOS = Currency{Alpha3: "SOS", Numeric: 706, MinorUnits: 2, Name: "Somali Shilling"}
	SSP = Currency{Alpha3: "SSP", Numeric: 728, MinorUnits: 2, Name: "South Sudanese Pound"}
	LKR = Currency{Alpha3: "LKR", Numeric: 144, MinorUnits: 2, Name: "Sri Lanka Rupee"}
	SDG = Currency{Alpha3: "SDG", Numeric: 938, MinorUnits: 2, Name: "Sudanese Pound"}
	SRD = Currency{Alpha3: "SRD", Numeric: 968, MinorUnits: 2, Name: "Surinam Dollar"}
	SEK = Currency{Alpha3: "SEK", Numeric: 752, MinorUnits: 2, Name: "Swedish Krona"}
	CHE = Currency{Alpha3: "CHE", Numeric: 947, MinorUnits: 2, Name: "WIR Euro"}
	CHW = Currency{Alpha3: "CHW", Numeric: 948, MinorUnits: 2, Name: "WIR Franc"}
	SYP = Currency{Alpha3: "SYP", Numeric: 760, MinorUnits: 2, Name: "Syrian Pound"}
	TWD = Currency{Alpha3: "TWD", Numeric: 901, MinorUnits: 2, Name: "New Taiwan Dollar"}
	TJS = Currency{Alpha3: "TJS", Numeric: 972, MinorUnits: 2, Name: "Somoni"}
	TZS = Currency{Alpha3: "TZS", Numeric: 834, MinorUnits: 2, Name: "Tanzanian Shilling"}
	THB = Currency{Alpha3: "THB", Numeric: 764, MinorUnits: 2, Name: "Baht"}
	TOP = Currency{Alpha3: "TOP", Numeric: 776, MinorUnits: 2, Name: "Pa’anga"}
	TTD = Currency{Alpha3: "TTD", Numeric: 780, MinorUnits: 2, Name: "Trinidad and Tobago Dollar"}
	TND = Currency{Alpha3: "TND", Numeric: 788, MinorUnits: 3, Name: "Tunisian Dinar"}
	TRY = Currency{Alpha3: "TRY", Numeric: 949, MinorUnits: 2, Name: "Turkish Lira"}
	TMT = Currency{Alpha3: "TMT", Numeric: 934, MinorUnits: 2, Name: "Turkmenistan New Manat"}
	UGX = Currency{Alpha3: "UGX", Numeric: 800, MinorUnits: 0, Name: "Uganda Shilling"}
	UAH = Currency{Alpha3: "UAH", Numeric: 980, MinorUnits: 2, Name: "Hryvnia"}
	AED = Currency{Alpha3: "AED", Numeric: 784, MinorUnits: 2, Name: "UAE Dirham"}
	USN = Currency{Alpha3: "USN", Numeric: 997, MinorUnits: 2, Name: "US Dollar (Next day)"}
	UYU = Currency{Alpha3: "UYU", Numeric: 858, MinorUnits: 2, Name: "Peso Uruguayo"}
	UYI = Currency{Alpha3: "UYI", Numeric: 940, MinorUnits: 0, Name: "Uruguay Peso en Unidades Indexadas (UI)"}
	UYW = Currency{Alpha3: "UYW", Numeric: 927, MinorUnits: 4, Name: "Unidad Previsional"}
	UZS = Currency{Alpha3: "UZS", Numeric: 860, MinorUnits: 2, Name: "Uzbekistan Sum"}
	VUV = Currency{Alpha3: "VUV", Numeric: 548, MinorUnits: 0, Name: "Vatu"}
	VES = Currency{Alpha3: "VES", Numeric: 928, MinorUnits: 2, Name: "Bolívar Soberano"}
	VED = Currency{Alpha3: "VED", Numeric: 926, MinorUnits: 2, Name: "Bolívar Soberano"}
	VND = Currency{Alpha3: "VND", Numeric: 704, MinorUnits: 0, Name: "Dong"}
	YER = Currency{Alpha3: "YER", Numeric: 886, MinorUnits: 2, Name: "Yemeni Rial"}
	ZMW = Currency{Alpha3: "ZMW", Numeric: 967, MinorUnits: 2, Name: "Zambian Kwacha"}
	ZWG = Currency{Alpha3: "ZWG", Numeric: 924, MinorUnits: 2, Name: "Zimbabwe Gold"}
	XBA = Currency{Alpha3: "XBA", Numeric: 955, MinorUnits: -1, Name: "Bond Markets Unit European Composite Unit (EURCO)"}
	XBB = Currency{Alpha3: "XBB", Numeric: 956, MinorUnits: -1, Name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)"}
	XBC = Currency{Alpha3: "XBC", Numeric: 957, MinorUnits: -1, Name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)"}
	XBD = Currency{Alpha3: "XBD", Numeric: 958, MinorUnits: -1, Name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)"}
	XTS = Currency{Alpha3: "XTS", Numeric: 963, MinorUnits: -1, Name: "Codes specifically reserved for testing purposes"}
	XXX = Currency{Alpha3: "XXX", Numeric: 999, MinorUnits: -1, Name: "The codes assigned for transactions where no currency is involved"}
	XAU = Currency{Alpha3: "XAU", Numeric: 959, MinorUnits: -1, Name: "Gold"}
	XPD = Currency{Alpha3: "XPD", Numeric: 964, MinorUnits: -1, Name: "Palladium"}
	XPT = Currency{Alpha3: "XPT", Numeric: 962, MinorUnits: -1, Name: "Platinum"}
	XAG = Currency{Alpha3: "XAG", Numeric: 961, MinorUnits: -1, Name: "Silver"}
)

var currenciesByAlpha3 = map[string]Currency{
	"AFN": AFN,
	"EUR": EUR,
	"ALL": ALL,
	"DZD": DZD,
	"USD": USD,
	"AOA": AOA,
	"XCD": XCD,
	"XAD": XAD,
	"ARS": ARS,
	"AMD": AMD,
	"AWG": AWG,
	"AUD": AUD,
	"AZN": AZN,
	"BSD": BSD,
	"BHD": BHD,
	"BDT": BDT,
	"BBD": BBD,
	"BYN": BYN,
	"BZD": BZD,
	"XOF": XOF,
	"BMD": BMD,
	"INR": INR,
	"BTN": BTN,
	"BOB": BOB,
	"BOV": BOV,
	"BAM": BAM,
	"BWP": BWP,
	"NOK": NOK,
	"BRL": BRL,
	"BND": BND,
	"BGN": BGN,
	"BIF": BIF,
	"CVE": CVE,
	"KHR": KHR,
	"XAF": XAF,
	"CAD": CAD,
	"KYD": KYD,
	"CLP": CLP,
	"CLF": CLF,
	"CNY": CNY,
	"COP": COP,
	"COU": COU,
	"KMF": KMF,
	"CDF": CDF,
	"NZD": NZD,
	"CRC": CRC,
	"CUP": CUP,
	"XCG": XCG,
	"CZK": CZK,
	"DKK": DKK,
	"DJF": DJF,
	"DOP": DOP,
	"EGP": EGP,
	"SVC": SVC,
	"ERN": ERN,
	"SZL": SZL,
	"ETB": ETB,
	"FKP": FKP,
	"FJD": FJD,
	"XPF": XPF,
	"GMD": GMD,
	"GEL": GEL,
	"GHS": GHS,
	"GIP": GIP,
	"GTQ": GTQ,
	"GBP": GBP,
	"GNF": GNF,
	"GYD": GYD,
	"HTG": HTG,
	"HNL": HNL,
	"HKD": HKD,
	"HUF": HUF,
	"ISK": ISK,
	"IDR": IDR,
	"XDR": XDR,
	"IRR": IRR,
	"IQD": IQD,
	"ILS": ILS,
	"JMD": JMD,
	"JPY": JPY,
	"JOD": JOD,
	"KZT": KZT,
	"KES": KES,
	"KPW": KPW,
	"KRW": KRW,
	"KWD": KWD,
	"KGS": KGS,
	"LAK": LAK,
	"LBP": LBP,
	"LSL": LSL,
	"ZAR": ZAR,
	"LRD": LRD,
	"LYD": LYD,
	"CHF": CHF,
	"MOP": MOP,
	"MKD": MKD,
	"MGA": MGA,
	"MWK": MWK,
	"MYR": MYR,
	"MVR": MVR,
	"MRU": MRU,
	"MUR": MUR,
	"XUA": XUA,
	"MXN": MXN,
	"MXV": MXV,
	"MDL": MDL,
	"MNT": MNT,
	"MAD": MAD,
	"MZN": MZN,
	"MMK": MMK,
	"NAD": NAD,
	"NPR": NPR,
	"NIO": NIO,
	"NGN": NGN,
	"OMR": OMR,
	"PKR": PKR,
	"PAB": PAB,
	"PGK": PGK,
	"PYG": PYG,
	"PEN": PEN,
	"PHP": PHP,
	"PLN": PLN,
	"QAR": QAR,
	"RON": RON,
	"RUB": RUB,
	"RWF": RWF,
	"SHP": SHP,
	"WST": WST,
	"STN": STN,
	"SAR": SAR,
	"RSD": RSD,
	"SCR": SCR,
	"SLE": SLE,
	"SGD": SGD,
	"XSU": XSU,
	"SBD": SBD,
	"SOS": SOS,
	"SSP": SSP,
	"LKR": LKR,
	"SDG": SDG,
	"SRD": SRD,
	"SEK": SEK,
	"CHE": CHE,
	"CHW": CHW,
	"SYP": SYP,
	"TWD": TWD,
	"TJS": TJS,
	"TZS": TZS,
	"THB": THB,
	"TOP": TOP,
	"TTD": TTD,
	"TND": TND,
	"TRY": TRY,
	"TMT": TMT,
	"UGX": UGX,
	"UAH": UAH,
	"AED": AED,
	"USN": USN,
	"UYU": UYU,
	"UYI": UYI,
	"UYW": UYW,
	"UZS": UZS,
	"VUV": VUV,
	"VES": VES,
	"VED": VED,
	"VND": VND,
	"YER": YER,
	"ZMW": ZMW,
	"ZWG": ZWG,
	"XBA": XBA,
	"XBB": XBB,
	"XBC": XBC,
	"XBD": XBD,
	"XTS": XTS,
	"XXX": XXX,
	"XAU": XAU,
	"XPD": XPD,
	"XPT": XPT,
	"XAG": XAG,
}

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
	"BV": NOK,
	"EG": EGP,
	"HN": HNL,
	"KN": XCD,
	"TJ": TJS,
	"TZ": TZS,
	"BJ": XOF,
	"BA": BAM,
	"BQ": USD,
	"MS": XCD,
	"RU": RUB,
	"CL": CLF,
	"GN": GNF,
	"GM": GMD,
	"GW": XOF,
	"PL": PLN,
	"RO": RON,
	"SG": SGD,
	"SS": SSP,
	"TL": USD,
	"TG": XOF,
	"HU": HUF,
	"NR": AUD,
	"AE": AED,
	"ML": XOF,
	"SR": SRD,
	"VG": USD,
	"AX": EUR,
	"JP": JPY,
	"NA": ZAR,
	"PG": PGK,
	"SA": SAR,
	"UG": UGX,
	"NC": XPF,
	"PN": NZD,
	"MP": USD,
	"SO": SOS,
	"SJ": NOK,
	"IE": EUR,
	"MT": EUR,
	"AR": ARS,
	"TW": TWD,
	"VU": VUV,
	"MU": MUR,
	"MA": MAD,
	"UY": UYW,
	"GT": GTQ,
	"MH": USD,
	"SD": SDG,
	"AM": AMD,
	"EE": EUR,
	"GE": GEL,
	"CK": NZD,
	"CU": CUP,
	"CI": XOF,
	"MC": EUR,
	"DO": DOP,
	"UA": UAH,
	"HK": HKD,
	"PK": PKR,
	"BS": BSD,
	"KH": KHR,
	"PT": EUR,
	"CF": XAF,
	"MM": MMK,
	"FK": FKP,
	"BF": XOF,
	"ET": ETB,
	"AO": AOA,
	"AW": AWG,
	"PW": USD,
	"TF": EUR,
	"LI": CHF,
	"KW": KWD,
	"KG": KGS,
	"EC": USD,
	"GI": GIP,
	"CW": XCG,
	"ID": IDR,
	"LB": LBP,
	"LT": EUR,
	"MW": MWK,
	"BL": EUR,
	"IO": USD,
	"CO": COU,
	"NO": NOK,
	"ZA": ZAR,
	"TT": TTD,
	"WF": XPF,
	"CR": CRC,
	"FO": DKK,
	"IL": ILS,
	"SE": SEK,
	"RS": RSD,
	"LV": EUR,
	"LR": LRD,
	"QA": QAR,
	"YE": YER,
	"TD": XAF,
	"HR": EUR,
	"NG": NGN,
	"RE": EUR,
	"TO": TOP,
	"PF": XPF,
	"JM": JMD,
	"UM": USD,
	"DK": DKK,
	"IN": INR,
	"KR": KRW,
	"MN": MNT,
	"SH": SHP,
	"VC": XCD,
	"AU": AUD,
	"CM": XAF,
	"KI": AUD,
	"CH": CHW,
	"AL": ALL,
	"GD": XCD,
	"NZ": NZD,
	"IQ": IQD,
	"MR": MRU,
	"TH": THB,
	"BW": BWP,
	"YT": EUR,
	"LS": ZAR,
	"MQ": EUR,
	"NE": XOF,
	"PH": PHP,
	"BD": BDT,
	"LY": LYD,
	"DJ": DJF,
	"FJ": FJD,
	"JO": JOD,
	"SI": EUR,
	"AT": EUR,
	"CY": EUR,
	"ER": ERN,
	"BZ": BZD,
	"CA": CAD,
	"FI": EUR,
	"LK": LKR,
	"SC": SCR,
	"GR": EUR,
	"KE": KES,
	"MD": MDL,
	"MZ": MZN,
	"NU": NZD,
	"RW": RWF,
	"VI": USD,
	"AS": USD,
	"CN": CNY,
	"OM": OMR,
	"LC": XCD,
	"WS": WST,
	"CX": AUD,
	"KZ": KZT,
	"NL": EUR,
	"MF": EUR,
	"DM": XCD,
	"FR": EUR,
	"DZ": DZD,
	"SK": EUR,
	"DE": EUR,
	"MV": MVR,
	"MX": MXV,
	"BN": BND,
	"GQ": XAF,
	"BG": BGN,
	"VN": VND,
	"PA": USD,
	"PY": PYG,
	"AZ": AZN,
	"GY": GYD,
	"TK": NZD,
	"NI": NIO,
	"SN": XOF,
	"AG": XCD,
	"SV": USD,
	"GH": GHS,
	"GP": EUR,
	"IT": EUR,
	"TC": USD,
	"BB": BBD,
	"CC": AUD,
	"MY": MYR,
	"TN": TND,
	"ZW": ZWG,
	"MO": MOP,
	"NP": NPR,
	"IM": GBP,
	"PM": EUR,
	"SX": XCG,
	"UZ": UZS,
	"BY": BYN,
	"BR": BRL,
	"BT": BTN,
	"GF": EUR,
	"LU": EUR,
	"AF": AFN,
	"GG": GBP,
	"ST": STN,
	"ZM": ZMW,
	"AD": EUR,
	"GA": XAF,
	"NF": AUD,
	"ES": EUR,
	"BM": BMD,
	"HT": USD,
	"BI": BIF,
	"KY": KYD,
	"CD": CDF,
	"HM": AUD,
	"IS": ISK,
	"MG": MGA,
	"AI": XCD,
	"BE": EUR,
	"TV": AUD,
	"PR": USD,
	"TM": TMT,
	"JE": GBP,
	"PE": PEN,
	"GU": USD,
	"ME": EUR,
	"SM": EUR,
	"SL": SLE,
	"BH": BHD,
	"GL": DKK,
}

var countryList = map[Currency][]string{
	TMT: []string{"TM"},
	TND: []string{"TN"},
	DOP: []string{"DO"},
	MXN: []string{"MX"},
	BAM: []string{"BA"},
	USD: []string{"AS", "BQ", "IO", "EC", "SV", "GU", "HT", "MH", "MP", "PW", "PA", "PR", "TL", "TC", "UM", "VG", "VI"},
	AMD: []string{"AM"},
	GTQ: []string{"GT"},
	CHW: []string{"CH"},
	DZD: []string{"DZ"},
	HKD: []string{"HK"},
	ISK: []string{"IS"},
	LBP: []string{"LB"},
	RSD: []string{"RS"},
	UYU: []string{"UY"},
	BND: []string{"BN"},
	NZD: []string{"CK", "NZ", "NU", "PN", "TK"},
	SVC: []string{"SV"},
	MWK: []string{"MW"},
	MYR: []string{"MY"},
	MUR: []string{"MU"},
	VUV: []string{"VU"},
	VND: []string{"VN"},
	ARS: []string{"AR"},
	SRD: []string{"SR"},
	RON: []string{"RO"},
	AOA: []string{"AO"},
	AUD: []string{"AU", "CX", "CC", "HM", "KI", "NR", "NF", "TV"},
	BGN: []string{"BG"},
	NGN: []string{"NG"},
	WST: []string{"WS"},
	ALL: []string{"AL"},
	GEL: []string{"GE"},
	ZAR: []string{"LS", "NA", "ZA"},
	MGA: []string{"MG"},
	SGD: []string{"SG"},
	BDT: []string{"BD"},
	IDR: []string{"ID"},
	MOP: []string{"MO"},
	MRU: []string{"MR"},
	TJS: []string{"TJ"},
	TZS: []string{"TZ"},
	HUF: []string{"HU"},
	CNY: []string{"CN"},
	PHP: []string{"PH"},
	ZWG: []string{"ZW"},
	BMD: []string{"BM"},
	JMD: []string{"JM"},
	NIO: []string{"NI"},
	THB: []string{"TH"},
	UAH: []string{"UA"},
	COU: []string{"CO"},
	SAR: []string{"SA"},
	PAB: []string{"PA"},
	RWF: []string{"RW"},
	SOS: []string{"SO"},
	UYW: []string{"UY"},
	EUR: []string{"AX", "AD", "AT", "BE", "HR", "CY", "EE", "FI", "FR", "GF", "TF", "DE", "GR", "GP", "IE", "IT", "LV", "LT", "LU", "MT", "MQ", "YT", "MC", "ME", "NL", "PT", "RE", "BL", "MF", "PM", "SM", "SK", "SI", "ES"},
	FJD: []string{"FJ"},
	ILS: []string{"IL"},
	JPY: []string{"JP"},
	JOD: []string{"JO"},
	CHF: []string{"LI", "CH"},
	PYG: []string{"PY"},
	QAR: []string{"QA"},
	XCG: []string{"CW", "SX"},
	SHP: []string{"SH"},
	SEK: []string{"SE"},
	RUB: []string{"RU"},
	SLE: []string{"SL"},
	PGK: []string{"PG"},
	BRL: []string{"BR"},
	CLP: []string{"CL"},
	GIP: []string{"GI"},
	KWD: []string{"KW"},
	MZN: []string{"MZ"},
	SCR: []string{"SC"},
	BBD: []string{"BB"},
	NAD: []string{"NA"},
	LKR: []string{"LK"},
	COP: []string{"CO"},
	KGS: []string{"KG"},
	MVR: []string{"MV"},
	BHD: []string{"BH"},
	CAD: []string{"CA"},
	GBP: []string{"GG", "IM", "JE"},
	LRD: []string{"LR"},
	PKR: []string{"PK"},
	PEN: []string{"PE"},
	UZS: []string{"UZ"},
	XCD: []string{"AI", "AG", "DM", "GD", "MS", "KN", "LC", "VC"},
	HTG: []string{"HT"},
	CUP: []string{"CU"},
	KZT: []string{"KZ"},
	STN: []string{"ST"},
	TOP: []string{"TO"},
	KHR: []string{"KH"},
	INR: []string{"BT", "IN"},
	ETB: []string{"ET"},
	LYD: []string{"LY"},
	BYN: []string{"BY"},
	MDL: []string{"MD"},
	DJF: []string{"DJ"},
	CLF: []string{"CL"},
	GHS: []string{"GH"},
	GYD: []string{"GY"},
	HNL: []string{"HN"},
	MXV: []string{"MX"},
	BZD: []string{"BZ"},
	BIF: []string{"BI"},
	GMD: []string{"GM"},
	LSL: []string{"LS"},
	SSP: []string{"SS"},
	UYI: []string{"UY"},
	NOK: []string{"BV", "NO", "SJ"},
	CDF: []string{"CD"},
	CRC: []string{"CR"},
	MNT: []string{"MN"},
	MMK: []string{"MM"},
	AFN: []string{"AF"},
	XAF: []string{"CM", "CF", "TD", "GQ", "GA"},
	EGP: []string{"EG"},
	ERN: []string{"ER"},
	FKP: []string{"FK"},
	MAD: []string{"MA"},
	OMR: []string{"OM"},
	TWD: []string{"TW"},
	BSD: []string{"BS"},
	AED: []string{"AE"},
	KYD: []string{"KY"},
	GNF: []string{"GN"},
	AWG: []string{"AW"},
	DKK: []string{"DK", "FO", "GL"},
	NPR: []string{"NP"},
	SDG: []string{"SD"},
	UGX: []string{"UG"},
	XOF: []string{"BJ", "BF", "CI", "GW", "ML", "NE", "SN", "TG"},
	BWP: []string{"BW"},
	KES: []string{"KE"},
	KRW: []string{"KR"},
	PLN: []string{"PL"},
	CHE: []string{"CH"},
	TTD: []string{"TT"},
	YER: []string{"YE"},
	AZN: []string{"AZ"},
	ZMW: []string{"ZM"},
	XPF: []string{"PF", "NC", "WF"},
	IQD: []string{"IQ"},
	BTN: []string{"BT"},
}

