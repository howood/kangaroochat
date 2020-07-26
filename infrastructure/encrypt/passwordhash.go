package encrypt

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

const SaltBufLen = 14
const HashtypeScrypt = "scrypt"
const HashtypeBcrypt = "bcrypt"
const WrongPasswordMessage = "Wrong Password"

type PasswordHash struct {
	Type         string
	ScryptN      int
	ScryptR      int
	ScryptP      int
	ScryptKeylen int
}

func (ph PasswordHash) getSalt() string {
	b := make([]byte, SaltBufLen)

	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		fmt.Println("error:", err)
	}

	salt := base64.StdEncoding.EncodeToString(b)
	return salt
}

func (ph PasswordHash) hashWithScrypt(password, saltstr string) (string, error) {
	salt := []byte(saltstr)
	converted, err := scrypt.Key([]byte(password), salt, ph.ScryptN, ph.ScryptR, ph.ScryptP, ph.ScryptKeylen)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(converted[:]), nil
}

func (ph PasswordHash) comparePasswordWithScrypt(hashedpassword, password, saltstr string) error {
	hashed, err := ph.hashWithScrypt(password, saltstr)
	if err != nil {
		return err
	}
	if hashedpassword != hashed {
		return errors.New(WrongPasswordMessage)
	}
	return nil
}

func (ph PasswordHash) hashWithBcrypt(password string) (string, error) {
	converted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(converted[:]), nil
}

func (ph PasswordHash) comparePasswordWithBcrypt(hashedpassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
}

func (ph PasswordHash) GetHashed(password string) (string, string, error) {
	switch ph.Type {
	case HashtypeScrypt:
		saltstr := ph.getSalt()
		hashedpasswd, err := ph.hashWithScrypt(password, saltstr)
		return hashedpasswd, saltstr, err
	case HashtypeBcrypt:
		hashedpasswd, err := ph.hashWithBcrypt(password)
		return hashedpasswd, "", err
	default:
		hashedpasswd, err := ph.hashWithBcrypt(password)
		return hashedpasswd, "", err
	}
}

func (ph PasswordHash) Compare(hashedpassword, password, saltstr string) error {
	switch ph.Type {
	case HashtypeScrypt:
		return ph.comparePasswordWithScrypt(hashedpassword, password, saltstr)
	case HashtypeBcrypt:
		return ph.comparePasswordWithBcrypt(hashedpassword, password)
	default:
		return ph.comparePasswordWithBcrypt(hashedpassword, password)
	}
}
