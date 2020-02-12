package entities

import (
	"fmt"
	"time"
)

// CurrencyRate defines the currencyRate functionality
type CurrencyRate interface {
	// Name
	Name() string
	// Rate holds the rate for this currencyRate in contrast with Euro.
	Rate() float64
	// LastUpdate holds the date this currencyRate was last updated.
	LastUpdate() time.Time
	// ConvertToEuros convert "amount" of this currencyRate to euros. Amount is in micros.
	ConvertToEuros(amount int64) int64
	// ConvertEuros convert "amount" euros to this currencyRate. Amount is in micros.
	ConvertEuros(amount int64) int64
	// ToString returns the string representation of the "amount" in this currencyRate.
	ToString(amount int64) string
}

type currencyRate struct {
	name       string
	rate       float64
	lastUpdate time.Time
}

// NewCurrencyRate is the currencyRate constructor
func NewCurrencyRate(name string, rate float64, lastUpdate time.Time) CurrencyRate {
	return &currencyRate{
		name:       name,
		rate:       rate,
		lastUpdate: lastUpdate,
	}
}

func (c *currencyRate) Name() string {
	return c.name
}

func (c *currencyRate) Rate() float64 {
	return c.rate
}

func (c *currencyRate) LastUpdate() time.Time {
	return c.lastUpdate
}

func (c *currencyRate) ConvertToEuros(amount int64) int64 {
	return int64(float64(amount) * c.Rate())
}

func (c *currencyRate) ConvertEuros(amount int64) int64 {
	return int64(float64(amount) / c.Rate())
}

func (c *currencyRate) ToString(amount int64) string {
	return fmt.Sprintf("%s %f", c.Name(), float64(amount)/micros)
}
