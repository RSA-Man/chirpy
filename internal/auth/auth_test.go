package auth

import (
	"testing"
)

// Test CheckPasswordHash tests the CheckPasswordHash function.
func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := CheckPasswordHash(password, hash)
	if err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Errorf("CheckPasswordHash returned false")
	}
}

// Test HashPassword tests the HashPassword function.
func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	} else if hash == "" {
		t.Errorf("HashPassword returned an empty string")
	}
}
