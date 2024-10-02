package web

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	password := "123456"
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encrypted)
	//assert.NoError(t, err)
}
