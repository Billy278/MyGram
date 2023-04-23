package cripto

import "golang.org/x/crypto/bcrypt"

func GenerateHash(payload string) (hashed string, err error) {
	// Hashing the password with the default cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	return string(hash), err
}

func CompareHash(hash, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return
	}
	return
}
