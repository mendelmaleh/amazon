package amazon

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

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
