package timezone

// Timezone holds data gathered from CSV files on http://timezonedb.com/download.
type Timezone struct {
	Location string
	Code     string
	Country  string
}

// Locations is a list of all IANA timezone locations and country codes.
var Locations = []Timezone{
	{"Africa/Abidjan", "CI", "Ivory Coast"},
	{"Africa/Accra", "GH", "Ghana"},
	{"Africa/Addis_Ababa", "ET", "Ethiopia"},
	{"Africa/Algiers", "DZ", "Algeria"},
	{"Africa/Asmara", "ER", "Eritrea"},
	{"Africa/Bamako", "ML", "Mali"},
	{"Africa/Bangui", "CF", "Central African Republic"},
	{"Africa/Banjul", "GM", "Gambia"},
	{"Africa/Bissau", "GW", "Guinea-Bissau"},
	{"Africa/Blantyre", "MW", "Malawi"},
	{"Africa/Brazzaville", "CG", "Republic of the Congo"},
	{"Africa/Bujumbura", "BI", "Burundi"},
	{"Africa/Cairo", "EG", "Egypt"},
	{"Africa/Casablanca", "MA", "Morocco"},
	{"Africa/Ceuta", "ES", "Spain"},
	{"Africa/Conakry", "GN", "Guinea"},
	{"Africa/Dakar", "SN", "Senegal"},
	{"Africa/Dar_es_Salaam", "TZ", "Tanzania"},
	{"Africa/Djibouti", "DJ", "Djibouti"},
	{"Africa/Douala", "CM", "Cameroon"},
	{"Africa/El_Aaiun", "EH", "Western Sahara"},
	{"Africa/Freetown", "SL", "Sierra Leone"},
	{"Africa/Gaborone", "BW", "Botswana"},
	{"Africa/Harare", "ZW", "Zimbabwe"},
	{"Africa/Johannesburg", "ZA", "South Africa"},
	{"Africa/Juba", "SS", "South Sudan"},
	{"Africa/Kampala", "UG", "Uganda"},
	{"Africa/Khartoum", "SD", "Sudan"},
	{"Africa/Kigali", "RW", "Rwanda"},
	{"Africa/Kinshasa", "CD", "Democratic Republic of the Congo"},
	{"Africa/Lagos", "NG", "Nigeria"},
	{"Africa/Libreville", "GA", "Gabon"},
	{"Africa/Lome", "TG", "Togo"},
	{"Africa/Luanda", "AO", "Angola"},
	{"Africa/Lubumbashi", "CD", "Democratic Republic of the Congo"},
	{"Africa/Lusaka", "ZM", "Zambia"},
	{"Africa/Malabo", "GQ", "Equatorial Guinea"},
	{"Africa/Maputo", "MZ", "Mozambique"},
	{"Africa/Maseru", "LS", "Lesotho"},
	{"Africa/Mbabane", "SZ", "Swaziland"},
	{"Africa/Mogadishu", "SO", "Somalia"},
	{"Africa/Monrovia", "LR", "Liberia"},
	{"Africa/Nairobi", "KE", "Kenya"},
	{"Africa/Ndjamena", "TD", "Chad"},
	{"Africa/Niamey", "NE", "Niger"},
	{"Africa/Nouakchott", "MR", "Mauritania"},
	{"Africa/Ouagadougou", "BF", "Burkina Faso"},
	{"Africa/Porto-Novo", "BJ", "Benin"},
	{"Africa/Sao_Tome", "ST", "Sao Tome and Principe"},
	{"Africa/Tripoli", "LY", "Libya"},
	{"Africa/Tunis", "TN", "Tunisia"},
	{"Africa/Windhoek", "NA", "Namibia"},
	{"America/Adak", "US", "United States"},
	{"America/Anchorage", "US", "United States"},
	{"America/Anguilla", "AI", "Anguilla"},
	{"America/Antigua", "AG", "Antigua and Barbuda"},
	{"America/Araguaina", "BR", "Brazil"},
	{"America/Argentina/Buenos_Aires", "AR", "Argentina"},
	{"America/Argentina/Catamarca", "AR", "Argentina"},
	{"America/Argentina/Cordoba", "AR", "Argentina"},
	{"America/Argentina/Jujuy", "AR", "Argentina"},
	{"America/Argentina/La_Rioja", "AR", "Argentina"},
	{"America/Argentina/Mendoza", "AR", "Argentina"},
	{"America/Argentina/Rio_Gallegos", "AR", "Argentina"},
	{"America/Argentina/Salta", "AR", "Argentina"},
	{"America/Argentina/San_Juan", "AR", "Argentina"},
	{"America/Argentina/San_Luis", "AR", "Argentina"},
	{"America/Argentina/Tucuman", "AR", "Argentina"},
	{"America/Argentina/Ushuaia", "AR", "Argentina"},
	{"America/Aruba", "AW", "Aruba"},
	{"America/Asuncion", "PY", "Paraguay"},
	{"America/Atikokan", "CA", "Canada"},
	{"America/Bahia", "BR", "Brazil"},
	{"America/Bahia_Banderas", "MX", "Mexico"},
	{"America/Barbados", "BB", "Barbados"},
	{"America/Belem", "BR", "Brazil"},
	{"America/Belize", "BZ", "Belize"},
	{"America/Blanc-Sablon", "CA", "Canada"},
	{"America/Boa_Vista", "BR", "Brazil"},
	{"America/Bogota", "CO", "Colombia"},
	{"America/Boise", "US", "United States"},
	{"America/Cambridge_Bay", "CA", "Canada"},
	{"America/Campo_Grande", "BR", "Brazil"},
	{"America/Cancun", "MX", "Mexico"},
	{"America/Caracas", "VE", "Venezuela"},
	{"America/Cayenne", "GF", "French Guiana"},
	{"America/Cayman", "KY", "Cayman Islands"},
	{"America/Chicago", "US", "United States"},
	{"America/Chihuahua", "MX", "Mexico"},
	{"America/Costa_Rica", "CR", "Costa Rica"},
	{"America/Creston", "CA", "Canada"},
	{"America/Cuiaba", "BR", "Brazil"},
	{"America/Curacao", "CW", "Curaçao"},
	{"America/Danmarkshavn", "GL", "Greenland"},
	{"America/Dawson", "CA", "Canada"},
	{"America/Dawson_Creek", "CA", "Canada"},
	{"America/Denver", "US", "United States"},
	{"America/Detroit", "US", "United States"},
	{"America/Dominica", "DM", "Dominica"},
	{"America/Edmonton", "CA", "Canada"},
	{"America/Eirunepe", "BR", "Brazil"},
	{"America/El_Salvador", "SV", "El Salvador"},
	{"America/Fortaleza", "BR", "Brazil"},
	{"America/Glace_Bay", "CA", "Canada"},
	{"America/Godthab", "GL", "Greenland"},
	{"America/Goose_Bay", "CA", "Canada"},
	{"America/Grand_Turk", "TC", "Turks and Caicos Islands"},
	{"America/Grenada", "GD", "Grenada"},
	{"America/Guadeloupe", "GP", "Guadeloupe"},
	{"America/Guatemala", "GT", "Guatemala"},
	{"America/Guayaquil", "EC", "Ecuador"},
	{"America/Guyana", "GY", "Guyana"},
	{"America/Halifax", "CA", "Canada"},
	{"America/Havana", "CU", "Cuba"},
	{"America/Hermosillo", "MX", "Mexico"},
	{"America/Indiana/Indianapolis", "US", "United States"},
	{"America/Indiana/Knox", "US", "United States"},
	{"America/Indiana/Marengo", "US", "United States"},
	{"America/Indiana/Petersburg", "US", "United States"},
	{"America/Indiana/Tell_City", "US", "United States"},
	{"America/Indiana/Vevay", "US", "United States"},
	{"America/Indiana/Vincennes", "US", "United States"},
	{"America/Indiana/Winamac", "US", "United States"},
	{"America/Inuvik", "CA", "Canada"},
	{"America/Iqaluit", "CA", "Canada"},
	{"America/Jamaica", "JM", "Jamaica"},
	{"America/Juneau", "US", "United States"},
	{"America/Kentucky/Louisville", "US", "United States"},
	{"America/Kentucky/Monticello", "US", "United States"},
	{"America/Kralendijk", "BQ", "Bonaire, Saint Eustatius and Saba"},
	{"America/La_Paz", "BO", "Bolivia"},
	{"America/Lima", "PE", "Peru"},
	{"America/Los_Angeles", "US", "United States"},
	{"America/Lower_Princes", "SX", "Sint Maarten"},
	{"America/Maceio", "BR", "Brazil"},
	{"America/Managua", "NI", "Nicaragua"},
	{"America/Manaus", "BR", "Brazil"},
	{"America/Marigot", "MF", "Saint Martin"},
	{"America/Martinique", "MQ", "Martinique"},
	{"America/Matamoros", "MX", "Mexico"},
	{"America/Mazatlan", "MX", "Mexico"},
	{"America/Menominee", "US", "United States"},
	{"America/Merida", "MX", "Mexico"},
	{"America/Metlakatla", "US", "United States"},
	{"America/Mexico_City", "MX", "Mexico"},
	{"America/Miquelon", "PM", "Saint Pierre and Miquelon"},
	{"America/Moncton", "CA", "Canada"},
	{"America/Monterrey", "MX", "Mexico"},
	{"America/Montevideo", "UY", "Uruguay"},
	{"America/Montserrat", "MS", "Montserrat"},
	{"America/Nassau", "BS", "Bahamas"},
	{"America/New_York", "US", "United States"},
	{"America/Nipigon", "CA", "Canada"},
	{"America/Nome", "US", "United States"},
	{"America/Noronha", "BR", "Brazil"},
	{"America/North_Dakota/Beulah", "US", "United States"},
	{"America/North_Dakota/Center", "US", "United States"},
	{"America/North_Dakota/New_Salem", "US", "United States"},
	{"America/Ojinaga", "MX", "Mexico"},
	{"America/Panama", "PA", "Panama"},
	{"America/Pangnirtung", "CA", "Canada"},
	{"America/Paramaribo", "SR", "Suriname"},
	{"America/Phoenix", "US", "United States"},
	{"America/Port-au-Prince", "HT", "Haiti"},
	{"America/Port_of_Spain", "TT", "Trinidad and Tobago"},
	{"America/Porto_Velho", "BR", "Brazil"},
	{"America/Puerto_Rico", "PR", "Puerto Rico"},
	{"America/Rainy_River", "CA", "Canada"},
	{"America/Rankin_Inlet", "CA", "Canada"},
	{"America/Recife", "BR", "Brazil"},
	{"America/Regina", "CA", "Canada"},
	{"America/Resolute", "CA", "Canada"},
	{"America/Rio_Branco", "BR", "Brazil"},
	{"America/Santa_Isabel", "MX", "Mexico"},
	{"America/Santarem", "BR", "Brazil"},
	{"America/Santiago", "CL", "Chile"},
	{"America/Santo_Domingo", "DO", "Dominican Republic"},
	{"America/Sao_Paulo", "BR", "Brazil"},
	{"America/Scoresbysund", "GL", "Greenland"},
	{"America/Sitka", "US", "United States"},
	{"America/St_Barthelemy", "BL", "Saint Barthélemy"},
	{"America/St_Johns", "CA", "Canada"},
	{"America/St_Kitts", "KN", "Saint Kitts and Nevis"},
	{"America/St_Lucia", "LC", "Saint Lucia"},
	{"America/St_Thomas", "VI", "U.S. Virgin Islands"},
	{"America/St_Vincent", "VC", "Saint Vincent and the Grenadines"},
	{"America/Swift_Current", "CA", "Canada"},
	{"America/Tegucigalpa", "HN", "Honduras"},
	{"America/Thule", "GL", "Greenland"},
	{"America/Thunder_Bay", "CA", "Canada"},
	{"America/Tijuana", "MX", "Mexico"},
	{"America/Toronto", "CA", "Canada"},
	{"America/Tortola", "VG", "British Virgin Islands"},
	{"America/Vancouver", "CA", "Canada"},
	{"America/Whitehorse", "CA", "Canada"},
	{"America/Winnipeg", "CA", "Canada"},
	{"America/Yakutat", "US", "United States"},
	{"America/Yellowknife", "CA", "Canada"},
	{"Antarctica/Casey", "AQ", "Antarctica"},
	{"Antarctica/Davis", "AQ", "Antarctica"},
	{"Antarctica/DumontDUrville", "AQ", "Antarctica"},
	{"Antarctica/Macquarie", "AU", "Australia"},
	{"Antarctica/Mawson", "AQ", "Antarctica"},
	{"Antarctica/McMurdo", "AQ", "Antarctica"},
	{"Antarctica/Palmer", "AQ", "Antarctica"},
	{"Antarctica/Rothera", "AQ", "Antarctica"},
	{"Antarctica/Syowa", "AQ", "Antarctica"},
	{"Antarctica/Troll", "AQ", "Antarctica"},
	{"Antarctica/Vostok", "AQ", "Antarctica"},
	{"Arctic/Longyearbyen", "SJ", "Svalbard and Jan Mayen"},
	{"Asia/Aden", "YE", "Yemen"},
	{"Asia/Almaty", "KZ", "Kazakhstan"},
	{"Asia/Amman", "JO", "Jordan"},
	{"Asia/Anadyr", "RU", "Russia"},
	{"Asia/Aqtau", "KZ", "Kazakhstan"},
	{"Asia/Aqtobe", "KZ", "Kazakhstan"},
	{"Asia/Ashgabat", "TM", "Turkmenistan"},
	{"Asia/Baghdad", "IQ", "Iraq"},
	{"Asia/Bahrain", "BH", "Bahrain"},
	{"Asia/Baku", "AZ", "Azerbaijan"},
	{"Asia/Bangkok", "TH", "Thailand"},
	{"Asia/Beirut", "LB", "Lebanon"},
	{"Asia/Bishkek", "KG", "Kyrgyzstan"},
	{"Asia/Brunei", "BN", "Brunei"},
	{"Asia/Chita", "RU", "Russia"},
	{"Asia/Choibalsan", "MN", "Mongolia"},
	{"Asia/Colombo", "LK", "Sri Lanka"},
	{"Asia/Damascus", "SY", "Syria"},
	{"Asia/Dhaka", "BD", "Bangladesh"},
	{"Asia/Dili", "TL", "East Timor"},
	{"Asia/Dubai", "AE", "United Arab Emirates"},
	{"Asia/Dushanbe", "TJ", "Tajikistan"},
	{"Asia/Gaza", "PS", "Palestinian Territory"},
	{"Asia/Hebron", "PS", "Palestinian Territory"},
	{"Asia/Ho_Chi_Minh", "VN", "Vietnam"},
	{"Asia/Hong_Kong", "HK", "Hong Kong"},
	{"Asia/Hovd", "MN", "Mongolia"},
	{"Asia/Irkutsk", "RU", "Russia"},
	{"Asia/Jakarta", "ID", "Indonesia"},
	{"Asia/Jayapura", "ID", "Indonesia"},
	{"Asia/Jerusalem", "IL", "Israel"},
	{"Asia/Kabul", "AF", "Afghanistan"},
	{"Asia/Kamchatka", "RU", "Russia"},
	{"Asia/Karachi", "PK", "Pakistan"},
	{"Asia/Kathmandu", "NP", "Nepal"},
	{"Asia/Khandyga", "RU", "Russia"},
	{"Asia/Kolkata", "IN", "India"},
	{"Asia/Krasnoyarsk", "RU", "Russia"},
	{"Asia/Kuala_Lumpur", "MY", "Malaysia"},
	{"Asia/Kuching", "MY", "Malaysia"},
	{"Asia/Kuwait", "KW", "Kuwait"},
	{"Asia/Macau", "MO", "Macao"},
	{"Asia/Magadan", "RU", "Russia"},
	{"Asia/Makassar", "ID", "Indonesia"},
	{"Asia/Manila", "PH", "Philippines"},
	{"Asia/Muscat", "OM", "Oman"},
	{"Asia/Nicosia", "CY", "Cyprus"},
	{"Asia/Novokuznetsk", "RU", "Russia"},
	{"Asia/Novosibirsk", "RU", "Russia"},
	{"Asia/Omsk", "RU", "Russia"},
	{"Asia/Oral", "KZ", "Kazakhstan"},
	{"Asia/Phnom_Penh", "KH", "Cambodia"},
	{"Asia/Pontianak", "ID", "Indonesia"},
	{"Asia/Pyongyang", "KP", "North Korea"},
	{"Asia/Qatar", "QA", "Qatar"},
	{"Asia/Qyzylorda", "KZ", "Kazakhstan"},
	{"Asia/Rangoon", "MM", "Myanmar"},
	{"Asia/Riyadh", "SA", "Saudi Arabia"},
	{"Asia/Sakhalin", "RU", "Russia"},
	{"Asia/Samarkand", "UZ", "Uzbekistan"},
	{"Asia/Seoul", "KR", "South Korea"},
	{"Asia/Shanghai", "CN", "China"},
	{"Asia/Singapore", "SG", "Singapore"},
	{"Asia/Srednekolymsk", "RU", "Russia"},
	{"Asia/Taipei", "TW", "Taiwan"},
	{"Asia/Tashkent", "UZ", "Uzbekistan"},
	{"Asia/Tbilisi", "GE", "Georgia"},
	{"Asia/Tehran", "IR", "Iran"},
	{"Asia/Thimphu", "BT", "Bhutan"},
	{"Asia/Tokyo", "JP", "Japan"},
	{"Asia/Ulaanbaatar", "MN", "Mongolia"},
	{"Asia/Urumqi", "CN", "China"},
	{"Asia/Ust-Nera", "RU", "Russia"},
	{"Asia/Vientiane", "LA", "Laos"},
	{"Asia/Vladivostok", "RU", "Russia"},
	{"Asia/Yakutsk", "RU", "Russia"},
	{"Asia/Yekaterinburg", "RU", "Russia"},
	{"Asia/Yerevan", "AM", "Armenia"},
	{"Atlantic/Azores", "PT", "Portugal"},
	{"Atlantic/Bermuda", "BM", "Bermuda"},
	{"Atlantic/Canary", "ES", "Spain"},
	{"Atlantic/Cape_Verde", "CV", "Cape Verde"},
	{"Atlantic/Faroe", "FO", "Faroe Islands"},
	{"Atlantic/Madeira", "PT", "Portugal"},
	{"Atlantic/Reykjavik", "IS", "Iceland"},
	{"Atlantic/South_Georgia", "GS", "South Georgia and the South Sandwich Islands"},
	{"Atlantic/St_Helena", "SH", "Saint Helena"},
	{"Atlantic/Stanley", "FK", "Falkland Islands"},
	{"Australia/Adelaide", "AU", "Australia"},
	{"Australia/Brisbane", "AU", "Australia"},
	{"Australia/Broken_Hill", "AU", "Australia"},
	{"Australia/Currie", "AU", "Australia"},
	{"Australia/Darwin", "AU", "Australia"},
	{"Australia/Eucla", "AU", "Australia"},
	{"Australia/Hobart", "AU", "Australia"},
	{"Australia/Lindeman", "AU", "Australia"},
	{"Australia/Lord_Howe", "AU", "Australia"},
	{"Australia/Melbourne", "AU", "Australia"},
	{"Australia/Perth", "AU", "Australia"},
	{"Australia/Sydney", "AU", "Australia"},
	{"Europe/Amsterdam", "NL", "Netherlands"},
	{"Europe/Andorra", "AD", "Andorra"},
	{"Europe/Athens", "GR", "Greece"},
	{"Europe/Belgrade", "RS", "Serbia"},
	{"Europe/Berlin", "DE", "Germany"},
	{"Europe/Bratislava", "SK", "Slovakia"},
	{"Europe/Brussels", "BE", "Belgium"},
	{"Europe/Bucharest", "RO", "Romania"},
	{"Europe/Budapest", "HU", "Hungary"},
	{"Europe/Busingen", "DE", "Germany"},
	{"Europe/Chisinau", "MD", "Moldova"},
	{"Europe/Copenhagen", "DK", "Denmark"},
	{"Europe/Dublin", "IE", "Ireland"},
	{"Europe/Gibraltar", "GI", "Gibraltar"},
	{"Europe/Guernsey", "GG", "Guernsey"},
	{"Europe/Helsinki", "FI", "Finland"},
	{"Europe/Isle_of_Man", "IM", "Isle of Man"},
	{"Europe/Istanbul", "TR", "Turkey"},
	{"Europe/Jersey", "JE", "Jersey"},
	{"Europe/Kaliningrad", "RU", "Russia"},
	{"Europe/Kiev", "UA", "Ukraine"},
	{"Europe/Lisbon", "PT", "Portugal"},
	{"Europe/Ljubljana", "SI", "Slovenia"},
	{"Europe/London", "GB", "United Kingdom"},
	{"Europe/Luxembourg", "LU", "Luxembourg"},
	{"Europe/Madrid", "ES", "Spain"},
	{"Europe/Malta", "MT", "Malta"},
	{"Europe/Mariehamn", "AX", "Aland Islands"},
	{"Europe/Minsk", "BY", "Belarus"},
	{"Europe/Monaco", "MC", "Monaco"},
	{"Europe/Moscow", "RU", "Russia"},
	{"Europe/Oslo", "NO", "Norway"},
	{"Europe/Paris", "FR", "France"},
	{"Europe/Podgorica", "ME", "Montenegro"},
	{"Europe/Prague", "CZ", "Czech Republic"},
	{"Europe/Riga", "LV", "Latvia"},
	{"Europe/Rome", "IT", "Italy"},
	{"Europe/Samara", "RU", "Russia"},
	{"Europe/San_Marino", "SM", "San Marino"},
	{"Europe/Sarajevo", "BA", "Bosnia and Herzegovina"},
	{"Europe/Simferopol", "RU", "Russia"},
	{"Europe/Skopje", "MK", "Macedonia"},
	{"Europe/Sofia", "BG", "Bulgaria"},
	{"Europe/Stockholm", "SE", "Sweden"},
	{"Europe/Tallinn", "EE", "Estonia"},
	{"Europe/Tirane", "AL", "Albania"},
	{"Europe/Uzhgorod", "UA", "Ukraine"},
	{"Europe/Vaduz", "LI", "Liechtenstein"},
	{"Europe/Vatican", "VA", "Vatican"},
	{"Europe/Vienna", "AT", "Austria"},
	{"Europe/Vilnius", "LT", "Lithuania"},
	{"Europe/Volgograd", "RU", "Russia"},
	{"Europe/Warsaw", "PL", "Poland"},
	{"Europe/Zagreb", "HR", "Croatia"},
	{"Europe/Zaporozhye", "UA", "Ukraine"},
	{"Europe/Zurich", "CH", "Switzerland"},
	{"Indian/Antananarivo", "MG", "Madagascar"},
	{"Indian/Chagos", "IO", "British Indian Ocean Territory"},
	{"Indian/Christmas", "CX", "Christmas Island"},
	{"Indian/Cocos", "CC", "Cocos Islands"},
	{"Indian/Comoro", "KM", "Comoros"},
	{"Indian/Kerguelen", "TF", "French Southern Territories"},
	{"Indian/Mahe", "SC", "Seychelles"},
	{"Indian/Maldives", "MV", "Maldives"},
	{"Indian/Mauritius", "MU", "Mauritius"},
	{"Indian/Mayotte", "YT", "Mayotte"},
	{"Indian/Reunion", "RE", "Reunion"},
	{"Pacific/Apia", "WS", "Samoa"},
	{"Pacific/Auckland", "NZ", "New Zealand"},
	{"Pacific/Bougainville", "PG", "Papua New Guinea"},
	{"Pacific/Chatham", "NZ", "New Zealand"},
	{"Pacific/Chuuk", "FM", "Micronesia"},
	{"Pacific/Easter", "CL", "Chile"},
	{"Pacific/Efate", "VU", "Vanuatu"},
	{"Pacific/Enderbury", "KI", "Kiribati"},
	{"Pacific/Fakaofo", "TK", "Tokelau"},
	{"Pacific/Fiji", "FJ", "Fiji"},
	{"Pacific/Funafuti", "TV", "Tuvalu"},
	{"Pacific/Galapagos", "EC", "Ecuador"},
	{"Pacific/Gambier", "PF", "French Polynesia"},
	{"Pacific/Guadalcanal", "SB", "Solomon Islands"},
	{"Pacific/Guam", "GU", "Guam"},
	{"Pacific/Honolulu", "US", "United States"},
	{"Pacific/Johnston", "UM", "United States Minor Outlying Islands"},
	{"Pacific/Kiritimati", "KI", "Kiribati"},
	{"Pacific/Kosrae", "FM", "Micronesia"},
	{"Pacific/Kwajalein", "MH", "Marshall Islands"},
	{"Pacific/Majuro", "MH", "Marshall Islands"},
	{"Pacific/Marquesas", "PF", "French Polynesia"},
	{"Pacific/Midway", "UM", "United States Minor Outlying Islands"},
	{"Pacific/Nauru", "NR", "Nauru"},
	{"Pacific/Niue", "NU", "Niue"},
	{"Pacific/Norfolk", "NF", "Norfolk Island"},
	{"Pacific/Noumea", "NC", "New Caledonia"},
	{"Pacific/Pago_Pago", "AS", "American Samoa"},
	{"Pacific/Palau", "PW", "Palau"},
	{"Pacific/Pitcairn", "PN", "Pitcairn"},
	{"Pacific/Pohnpei", "FM", "Micronesia"},
	{"Pacific/Port_Moresby", "PG", "Papua New Guinea"},
	{"Pacific/Rarotonga", "CK", "Cook Islands"},
	{"Pacific/Saipan", "MP", "Northern Mariana Islands"},
	{"Pacific/Tahiti", "PF", "French Polynesia"},
	{"Pacific/Tarawa", "KI", "Kiribati"},
	{"Pacific/Tongatapu", "TO", "Tonga"},
	{"Pacific/Wake", "UM", "United States Minor Outlying Islands"},
	{"Pacific/Wallis", "WF", "Wallis and Futuna"},
}
