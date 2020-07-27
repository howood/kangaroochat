package actor

import (
	"github.com/howood/kangaroochat/infrastructure/encrypt"
	"github.com/howood/kangaroochat/library/utils"
)

var (
	usetype      = utils.GetOsEnv("PASSOWRD_USETYPE", "scrypt")
	scryptN      = utils.GetOsEnvInt("PASSOWRD_SCRYPTN", 32768)
	scryptR      = utils.GetOsEnvInt("PASSOWRD_SCRYPTR", 8)
	scryptP      = utils.GetOsEnvInt("PASSOWRD_SCRYPTP", 1)
	scryptkeyLen = utils.GetOsEnvInt("PASSOWRD_SCRYPTKEYLEN", 32)
)

//PasswordOperator struct
type PasswordOperator struct {
}

//GetHashedPassword get hashed password
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

//ComparePassword compare hashed password and string password
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
