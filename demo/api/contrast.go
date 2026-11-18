package main

import (
	"fmt"
	"math"
	"regexp"
)

var hexRe = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)

func lum(hex string) float64 {
	var r, g, b int
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	f := func(c int) float64 {
		v := float64(c) / 255
		if v <= 0.03928 {
			return v / 12.92
		}
		return math.Pow((v+0.055)/1.055, 2.4)
	}
	return 0.2126*f(r) + 0.7152*f(g) + 0.0722*f(b)
}

// contrast returns the WCAG contrast ratio between two #rrggbb colours (1..21).
func contrast(a, b string) float64 {
	la, lb := lum(a), lum(b)
	if la < lb {
		la, lb = lb, la
	}
	return (la + 0.05) / (lb + 0.05)
}

// minContrast is the lowest ratio we accept for a code to stay scannable.
const minContrast = 3.0
