package amazon

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	parts := strings.FieldsFunc(string(data), func(r rune) bool {
		return r == '(' || r == ')'
	})

	if len(parts) != 2 {
		return fmt.Errorf("expected 2 parts from tracking, got %d", len(parts))
	}

	t.Carrier = parts[0]
	t.Number = parts[1]

	return nil
}

type Currency struct {
	Symbol string
	Cents  int
}

func (c *Currency) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	trimmed := strings.TrimLeftFunc(string(data), func(r rune) bool {
		return !unicode.IsDigit(r)
	})

	if s := len(data) - len(trimmed); s > 0 {
		c.Symbol = string(data[:s])
	}

	whole, cents, ok := strings.Cut(trimmed, ".")
	if !ok {
		return fmt.Errorf("error splitting wholes and cents: %q", data)
	}

	iw, err := strconv.Atoi(whole)
	if err != nil {
		return fmt.Errorf("error parsing wholes: %q", data)
	}

	ic, err := strconv.Atoi(cents)
	if err != nil {
		return fmt.Errorf("error parsing cents: %q", data)
	}

	c.Cents = iw*100 + ic

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
