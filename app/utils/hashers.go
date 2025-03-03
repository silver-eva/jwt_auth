package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

type bryptHasher struct {
	complixity int
}

func NewBcryptHasher(complixity interface{}) *bryptHasher {
	return &bryptHasher{
		complixity: complixity.(int),
	}
}

func (b *bryptHasher) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), b.complixity)
	return base64.StdEncoding.EncodeToString(hashed), err
}

func (b *bryptHasher) CheckPasswordHash(password, hash string) bool {
	decodedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(decodedHash, []byte(password))
	return err == nil
}

type argon2Hasher struct {
	complixity string // salt
	iterations uint32
	memory     uint32
	threads    uint8
	keyLen     uint32
}

func NewArgon2Hasher(complixity interface{}) *argon2Hasher {
	return &argon2Hasher{
		complixity: complixity.(string),
		iterations: 3,
		memory:     64 * 1024,
		threads:    4,
		keyLen:     32,
	}
}

func (a *argon2Hasher) HashPassword(password string) (string, error) {
	saltBytes := []byte(a.complixity)
	hashed := argon2.IDKey(
		[]byte(password), 
		saltBytes, 
		a.iterations,
		a.memory,
		a.threads,
		a.keyLen,
	)
	return base64.StdEncoding.EncodeToString(hashed), nil
}

func (a *argon2Hasher) CheckPasswordHash(password, hash string) bool {
	decodedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}
	saltBytes := []byte(a.complixity)
	otherHashed := argon2.IDKey(
		[]byte(password), 
		saltBytes, 
		a.iterations,
		a.memory,
		a.threads,
		a.keyLen,
	)
	return string(decodedHash) == string(otherHashed)
	
}