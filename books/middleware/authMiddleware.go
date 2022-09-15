package middleware

func BasicAuth(email, password string) (bool, error) {
	if email == "admin@gmail.com" && password == "admin" {
		return true, nil
	}
	return false, nil
}
