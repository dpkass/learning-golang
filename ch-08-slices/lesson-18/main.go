package main

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	hasUpper, hasDigit := false, false
	for _, c := range password {
		if 'A' <= c && c <= 'Z' {
			hasUpper = true
		}
		if '0' <= c && c <= '9' {
			hasDigit = true
		}
	}
	return hasUpper && hasDigit
}
