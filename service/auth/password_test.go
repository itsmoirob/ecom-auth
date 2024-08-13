package auth

import "testing"

func TestHasPassw(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash to not be empty")
	}

	if hash == "password" {
		t.Errorf("expected hash to not be equal to password")
	}
}


func TestComparePassword(t *testing.T) {
    hash, err := HashPassword("password")
    if err != nil {
        t.Errorf("error hashing password: %v", err)
    }

    if !ComparePasswords(hash, "password") {
        t.Errorf("expected password to match hash")
    }

    if ComparePasswords(hash, "password1") {
        t.Errorf("expected password to not match hash")
    }
}