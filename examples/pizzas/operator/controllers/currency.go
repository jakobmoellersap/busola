package controllers

import (
	"fmt"
	"strconv"
	"strings"
)

// USD represents US dollar amount in terms of cents
type USD int64

// ToUSD converts a float64 to USD
// e.g. 1.23 to $1.23, 1.345 to $1.35
func ToUSD(f float64) USD {
	return USD((f * 100) + 0.5)
}

// Float64 converts a USD to float64
func (m USD) Float64() float64 {
	x := float64(m)
	x = x / 100
	return x
}

// Multiply safely multiplies a USD value by a float64, rounding
// to the nearest cent.
func (m USD) Multiply(f float64) USD {
	x := (float64(m) * f) + 0.5
	return USD(x)
}

// String returns a formatted USD value
func (m USD) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("$%.2f", x)
}

func NewUSD(usd string) (USD, error) {
	if fval, err := strconv.ParseFloat(strings.ReplaceAll(strings.TrimPrefix(usd, "$"), ",", "."), 64); err != nil {
		return 0, err
	} else {
		usdVal := ToUSD(fval)
		return usdVal, nil
	}
}
