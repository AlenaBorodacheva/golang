package main

import "unicode"

func valid(pass string) string {
	wrong := "Wrong password"

	if len(pass) < 5 {
		return wrong
	}

	for _, r := range pass {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return wrong
		}
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return wrong
		}
	}

	return "OK"
}
