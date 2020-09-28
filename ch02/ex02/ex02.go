// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/martinlindhe/unit"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = append(args, "0.0")
	}
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		// Temprature
		f := unit.FromFahrenheit(t)
		c := unit.FromCelsius(t)
		fmt.Printf("%g°F = %g°C, %g°C = %g°F\n",
			t, f.Celsius(), t, c.Fahrenheit())

		// Length
		feet := unit.Length(t) * unit.Foot
		metter := unit.Length(t) * unit.Meter
		fmt.Printf("%g Feet = %g Metters, %g Metters = %g Feet\n",
			t, feet.Meters(), t, metter.Feet())
	}
}

//!-
