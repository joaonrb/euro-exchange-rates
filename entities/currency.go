package entities

import (
	"fmt"
	"time"
)

const micros float64 = 1000000.

// Currency defines the currency functionality
type Currency interface {
	// Name
	Name() string
	// Rate holds the rate for this currency in contrast with Euro.
	Rate() float64
	// LastUpdate holds the date this currency was last updated.
	LastUpdate() time.Time
	// ConvertToEuros convert "amount" of this currency to euros. Amount is in micros.
	ConvertToEuros(amount uint64) uint64
	// ConvertEuros convert "amount" euros to this currency. Amount is in micros.
	ConvertEuros(amount uint64) uint64
	// ToString returns the string representation of the "amount" in this currency.
	ToString(amount int64) string
}

type currency struct {
	name       string
	rate       float64
	lastUpdate time.Time
}

// NewCurrency is the currency constructor
func NewCurrency(name string, rate float64, lastUpdate time.Time) Currency {
	return &currency{
		name:       name,
		rate:       rate,
		lastUpdate: lastUpdate,
	}
}

func (c *currency) Name() string {
	return c.name
}

func (c *currency) Rate() float64 {
	return c.rate
}

func (c *currency) LastUpdate() time.Time {
	return c.lastUpdate
}

func (c *currency) ConvertToEuros(amount uint64) uint64 {
	return uint64(float64(amount) * c.Rate())
}

func (c *currency) ConvertEuros(amount uint64) uint64 {
	return uint64(float64(amount) / c.Rate())
}

func (c *currency) ToString(amount int64) string {
	return fmt.Sprintf("%s %f", c.Name(), float64(amount)/micros)
}
