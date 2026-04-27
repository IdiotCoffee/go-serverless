package validators

import "regexp"

func isEmailValid(email string) bool {
	var rxEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}
	return true
}
