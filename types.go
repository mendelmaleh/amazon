package amazon

// partially generated with git.sr.ht/~mendelmaleh/csvgen

type Common struct {
	Shipping Shipping `csv:",inline"`

	ID    string `csv:"Order ID"`
	Date  DateUS `csv:"Order Date"`
	Email string `csv:"Ordering Customer Email"`
	Buyer string `csv:"Buyer Name"`
	Group string `csv:"Group Name"`

	OrderNumber string `csv:"Purchase Order Number"`
	Website     string `csv:"Website"`
}

type Shipping struct {
	Date    DateUS `csv:"Shipment Date"`
	Name    string `csv:"Shipping Address Name"`
	Street1 string `csv:"Shipping Address Street 1"`
	Street2 string `csv:"Shipping Address Street 2"`
	ZIP     string `csv:"Shipping Address Zip"`
	City    string `csv:"Shipping Address City"`
	State   string `csv:"Shipping Address State"`
}

type OrderInfo struct {
	Status                string   `csv:"Order Status"`
	PaymentInstrumentType string   `csv:"Payment Instrument Type"`
	Tracking              Tracking `csv:"Carrier Name & Tracking Number"`
}

type ItemInfo struct {
	AsinIsbn          string `csv:"ASIN/ISBN"`
	Quantity          int    `csv:"Quantity"`
	Title             string `csv:"Title"`
	Category          string `csv:"Category"`
	Seller            string `csv:"Seller"`
	SellerCredentials string `csv:"Seller Credentials"`
}

type Item struct {
	Common
	OrderInfo
	ItemInfo

	Condition   string  `csv:"Condition"`
	UnspscCode  string  `csv:"UNSPSC Code"`
	ReleaseDate DateISO `csv:"Release Date"`

	Currency             string   `csv:"Currency"`
	ListPricePerUnit     Currency `csv:"List Price Per Unit"`
	PurchasePricePerUnit Currency `csv:"Purchase Price Per Unit"`
	TaxExemptionApplied  string   `csv:"Tax Exemption Applied"`
	TaxExemptionType     string   `csv:"Tax Exemption Type"`
	ExemptionOptOut      string   `csv:"Exemption Opt-Out"`
	ItemSubtotal         Currency `csv:"Item Subtotal"`
	ItemSubtotalTax      Currency `csv:"Item Subtotal Tax"`
	ItemTotal            Currency `csv:"Item Total"`
	PoLineNumber         string   `csv:"PO Line Number"`
}

type Order struct {
	Common
	OrderInfo

	ShippingCharge      Currency `csv:"Shipping Charge"`
	Subtotal            Currency `csv:"Subtotal"`
	TaxBeforePromotions Currency `csv:"Tax Before Promotions"`
	TaxCharged          Currency `csv:"Tax Charged"`
	TotalCharged        Currency `csv:"Total Charged"`
	TotalPromotions     Currency `csv:"Total Promotions"`
}

type Return struct {
	Common
	ItemInfo

	Date   DateUS `csv:"Return Date"`
	Reason string `csv:"Return Reason"`
}
