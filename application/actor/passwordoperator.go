package actor

import (
	"github.com/howood/kangaroochat/infrastructure/encrypt"
)

const (
	usetype      = "scrypt"
	scryptN      = 32768
	scryptR      = 8
	scryptP      = 1
	scryptkeyLen = 32
)

type PasswordOperator struct {
}

func (po PasswordOperator) GetHashedPassword(password string) (string, string, error) {
	passwordhash := encrypt.PasswordHash{
		Type:         usetype,
		ScryptN:      scryptN,
		ScryptR:      scryptR,
		ScryptP:      scryptP,
		ScryptKeylen: scryptkeyLen,
	}
	return passwordhash.GetHashed(password)
}

func (po PasswordOperator) ComparePassword(hashedpassword, password, salt string) error {
	passwordhash := encrypt.PasswordHash{
		Type:         usetype,
		ScryptN:      scryptN,
		ScryptR:      scryptR,
		ScryptP:      scryptP,
		ScryptKeylen: scryptkeyLen,
	}
	return passwordhash.Compare(hashedpassword, password, salt)
}
