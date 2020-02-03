package entities

import (
	. "gotest.tools/assert"
	"testing"
	"time"
)

func TestNewCurrency(t *testing.T) {
	data := []struct {
		name string
		rate float64
	}{
		{"IDR", 15091.51},
		{"USD", 1.1029},
		{"JPY", 122.31},
	}
	now := time.Now().UTC()
	for _, d := range data {
		c := NewCurrency(d.name, d.rate, now)
		Assert(t, c.Name() == d.name, "name attribute was not the expected: %s != %s", c.Name(), d.name)
		Assert(t, c.Rate() == d.rate, "rate attribute was not the expected: %f != %f", c.Rate, d.rate)
		Assert(t, c.LastUpdate() == now, "name attribute was not the expected: %s != %s", c.LastUpdate(), now)
	}
}

func TestCurrencyConvertToEuros(t *testing.T) {
	var amount uint64 = 1000000
	data := []struct {
		name     string
		rate     float64
		expected uint64
	}{
		{"IDR", 15091.51, 15091510000},
		{"USD", 1.1029, 1102900},
		{"JPY", 122.31, 122310000},
	}
	now := time.Now().UTC()
	for _, d := range data {
		c := NewCurrency(d.name, d.rate, now)
		converted := c.ConvertToEuros(amount)
		Assert(t, converted == d.expected, "expected amount from ConvertToEuros was not the expected: %d != %d", converted, d.expected)
	}
}

func TestCurrencyConvertEuros(t *testing.T) {
	var amount uint64 = 1000000
	data := []struct {
		name     string
		rate     float64
		expected uint64
	}{
		{"IDR", 15091.51, 66},
		{"USD", 1.1029, 906700},
		{"JPY", 122.31, 8175},
	}
	now := time.Now().UTC()
	for _, d := range data {
		c := NewCurrency(d.name, d.rate, now)
		converted := c.ConvertEuros(amount)
		Assert(t, converted == d.expected, "expected amount from ConvertEuros was not the expected: %d != %d", converted, d.expected)
	}
}

func TestCurrencyToString(t *testing.T) {
	var amount uint64 = 1000000
	data := []struct {
		name     string
		rate     float64
		expected string
	}{
		{"IDR", 15091.51, "IDR 1.000000"},
		{"USD", 1.1029, "USD 1.000000"},
		{"JPY", 122.31, "JPY 1.000000"},
	}
	now := time.Now().UTC()
	for _, d := range data {
		c := NewCurrency(d.name, d.rate, now)
		repr := c.ToString(int64(amount))
		Assert(t, repr == d.expected, "expected representation from ToString was not the expected: %s != %s", repr, d.expected)
	}
}
