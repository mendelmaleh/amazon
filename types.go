package amazon

// generated with git.sr.ht/~mendelmaleh/csvgen

type ABC struct {
	BuyerName              string `csv:"Buyer Name"`
	GroupName              string `csv:"Group Name"`
	OrderDate              string `csv:"Order Date"`
	OrderId                string `csv:"Order ID"`
	OrderingCustomerEmail  string `csv:"Ordering Customer Email"`
	PurchaseOrderNumber    string `csv:"Purchase Order Number"`
	ShipmentDate           string `csv:"Shipment Date"`
	ShippingAddressCity    string `csv:"Shipping Address City"`
	ShippingAddressName    string `csv:"Shipping Address Name"`
	ShippingAddressState   string `csv:"Shipping Address State"`
	ShippingAddressStreet1 string `csv:"Shipping Address Street 1"`
	ShippingAddressStreet2 string `csv:"Shipping Address Street 2"`
	ShippingAddressZip     string `csv:"Shipping Address Zip"`
	Website                string `csv:"Website"`
}

type AB struct {
	CarrierNameTrackingNumber string `csv:"Carrier Name & Tracking Number"`
	OrderStatus               string `csv:"Order Status"`
	PaymentInstrumentType     string `csv:"Payment Instrument Type"`
}

type AC struct {
	AsinIsbn          string `csv:"ASIN/ISBN"`
	Category          string `csv:"Category"`
	Quantity          string `csv:"Quantity"`
	Seller            string `csv:"Seller"`
	SellerCredentials string `csv:"Seller Credentials"`
	Title             string `csv:"Title"`
}

type A struct {
	ABC
	AB
	AC

	Condition            string `csv:"Condition"`
	Currency             string `csv:"Currency"`
	ExemptionOptOut      string `csv:"Exemption Opt-Out"`
	ItemSubtotal         string `csv:"Item Subtotal"`
	ItemSubtotalTax      string `csv:"Item Subtotal Tax"`
	ItemTotal            string `csv:"Item Total"`
	ListPricePerUnit     string `csv:"List Price Per Unit"`
	PoLineNumber         string `csv:"PO Line Number"`
	PurchasePricePerUnit string `csv:"Purchase Price Per Unit"`
	ReleaseDate          string `csv:"Release Date"`
	TaxExemptionApplied  string `csv:"Tax Exemption Applied"`
	TaxExemptionType     string `csv:"Tax Exemption Type"`
	UnspscCode           string `csv:"UNSPSC Code"`
}

type B struct {
	ABC
	AB

	ShippingCharge      string `csv:"Shipping Charge"`
	Subtotal            string `csv:"Subtotal"`
	TaxBeforePromotions string `csv:"Tax Before Promotions"`
	TaxCharged          string `csv:"Tax Charged"`
	TotalCharged        string `csv:"Total Charged"`
	TotalPromotions     string `csv:"Total Promotions"`
}

type C struct {
	ABC
	AC

	ReturnDate   string `csv:"Return Date"`
	ReturnReason string `csv:"Return Reason"`
}
