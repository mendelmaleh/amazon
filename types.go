package amazon

import (
	"bytes"
	"fmt"
	"time"
)

// partially generated with git.sr.ht/~mendelmaleh/csvgen

type DateUS struct {
	time.Time
}

func (d *DateUS) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	t, err := time.Parse("01/02/06", string(data))
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

type DateISO struct {
	time.Time
}

func (d *DateISO) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05", string(data))
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

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

type Tracking struct {
	Carrier string
	Number  string
}

func (t *Tracking) UnmarshalText(data []byte) error {
	parts := bytes.FieldsFunc(data, func(r rune) bool {
		return r == '(' || r == ')'
	})

	if len(parts) != 2 {
		return fmt.Errorf("expected 2 parts from tracking, got %d", len(parts))
	}

	t.Carrier = string(parts[0])
	t.Number = string(parts[1])

	return nil
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

	Currency             string `csv:"Currency"`
	ListPricePerUnit     string `csv:"List Price Per Unit"`
	PurchasePricePerUnit string `csv:"Purchase Price Per Unit"`
	TaxExemptionApplied  string `csv:"Tax Exemption Applied"`
	TaxExemptionType     string `csv:"Tax Exemption Type"`
	ExemptionOptOut      string `csv:"Exemption Opt-Out"`
	ItemSubtotal         string `csv:"Item Subtotal"`
	ItemSubtotalTax      string `csv:"Item Subtotal Tax"`
	ItemTotal            string `csv:"Item Total"`
	PoLineNumber         string `csv:"PO Line Number"`
}

type Order struct {
	Common
	OrderInfo

	ShippingCharge      string `csv:"Shipping Charge"`
	Subtotal            string `csv:"Subtotal"`
	TaxBeforePromotions string `csv:"Tax Before Promotions"`
	TaxCharged          string `csv:"Tax Charged"`
	TotalCharged        string `csv:"Total Charged"`
	TotalPromotions     string `csv:"Total Promotions"`
}

type Return struct {
	Common
	ItemInfo

	Date   DateUS `csv:"Return Date"`
	Reason string `csv:"Return Reason"`
}
