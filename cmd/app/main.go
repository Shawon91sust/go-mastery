package main

import (
	"flag"
	"fmt"
)

type Money struct {
	AmountCents int64  // 0 by default (zero value)
	Currency    string // "" by default (zero value)
}

func (m Money) String() string {
	// Ensure safe defaults
	cur := m.Currency
	if cur == "" {
		cur = "USD"
	}

	sign := ""
	cents := m.AmountCents
	if cents < 0 {
		sign = "-"
		cents = -cents
	}

	dollars := cents / 100
	remainder := cents % 100

	return fmt.Sprintf("%s%s %d.%02d", sign, cur, dollars, remainder)
}

func main() {
	amount := flag.Int64("cents", 0, "amount in cents (e.g., 12345 = 123.45)")
	currency := flag.String("cur", "USD", "currency code (e.g., USD, BDT)")
	flag.Parse()

	m := Money{
		AmountCents: *amount,
		Currency:    *currency,
	}

	fmt.Println(m.String())

	// Zero value demo (important in Go)
	var z Money
	fmt.Printf("Raw struct with %%#v: %#v\n", z)
	fmt.Printf("Default formatting (Stringer): %v\n", z)
	fmt.Printf("Explicit String(): %s\n", z.String())
}
