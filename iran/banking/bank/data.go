package bank

// byCode is a partial map of Sheba bank codes from Bank Keshavarzi documentation
// and subsequent commonly listed codes. Not exhaustive for all current banks.
var byCode = map[string]Bank{
	"010": {Code: "010", Name: "Central Bank of Iran", PersianName: "بانک مرکزی جمهوری اسلامی ایران"},
	"011": {Code: "011", Name: "Bank of Industry and Mine", PersianName: "بانک صنعت و معدن"},
	"012": {Code: "012", Name: "Bank Mellat", PersianName: "بانک ملت"},
	"013": {Code: "013", Name: "Bank Refah", PersianName: "بانک رفاه"},
	"014": {Code: "014", Name: "Bank Maskan", PersianName: "بانک مسکن"},
	"015": {Code: "015", Name: "Bank Sepah", PersianName: "بانک سپه"},
	"016": {Code: "016", Name: "Bank Keshavarzi", PersianName: "بانک کشاورزی"},
	"017": {Code: "017", Name: "Bank Melli Iran", PersianName: "بانک ملی ایران"},
	"018": {Code: "018", Name: "Bank Tejarat", PersianName: "بانک تجارت"},
	"019": {Code: "019", Name: "Bank Saderat Iran", PersianName: "بانک صادرات ایران"},
	"020": {Code: "020", Name: "Export Development Bank", PersianName: "بانک توسعه صادرات"},
	"021": {Code: "021", Name: "Post Bank of Iran", PersianName: "پست بانک ایران"},
	"022": {Code: "022", Name: "Tose'e Ta'avon Bank", PersianName: "بانک توسعه تعاون"},
	"053": {Code: "053", Name: "Karafarin Bank", PersianName: "بانک کارآفرین"},
	"054": {Code: "054", Name: "Parsian Bank", PersianName: "بانک پارسیان"},
	"055": {Code: "055", Name: "Eghtesad Novin Bank", PersianName: "بانک اقتصاد نوین"},
	"056": {Code: "056", Name: "Saman Bank", PersianName: "بانک سامان"},
	"057": {Code: "057", Name: "Pasargad Bank", PersianName: "بانک پاسارگاد"},
	"058": {Code: "058", Name: "Sarmayeh Bank", PersianName: "بانک سرمایه"},
	"059": {Code: "059", Name: "Sina Bank", PersianName: "بانک سینا"},
	"061": {Code: "061", Name: "Shahr Bank", PersianName: "بانک شهر"},
	"062": {Code: "062", Name: "Ayandeh Bank", PersianName: "بانک آینده"}, // historically Tat/Ayandeh lineage; verify before relying
	"063": {Code: "063", Name: "Ansar Bank", PersianName: "بانک انصار"},
	"064": {Code: "064", Name: "Tourism Bank", PersianName: "بانک گردشگری"},
	"066": {Code: "066", Name: "Day Bank", PersianName: "بانک دی"},
	"069": {Code: "069", Name: "Iran Zamin Bank", PersianName: "بانک ایران زمین"},
}

// binToCode is a PARTIAL, well-known BIN→bank-code map.
// Unknown BINs are not invalid; they simply have no metadata.
// Sources: commonly published Iranian BIN lists in open documentation.
// Do not treat as complete. Updated: 2026-11-11
var binToCode = map[string]string{
	"603799": "017", // Melli
	"589210": "015", // Sepah
	"627884": "018", // Tejarat (commonly listed)
	"610433": "012", // Mellat
	"628023": "014", // Maskan
	"627961": "011", // Industry and Mine
	"603770": "016", // Keshavarzi
	"627648": "020", // Export Development
	"627760": "019", // Saderat
	"502229": "054", // Parsian
	"621986": "056", // Saman
	"627412": "053", // Karafarin
	"622106": "057", // Pasargad
	"627488": "057", // Pasargad
	"639347": "061", // Shahr
}
