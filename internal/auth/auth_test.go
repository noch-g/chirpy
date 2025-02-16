package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	err = CheckPassword(password, hashedPassword)
	if err != nil {
		t.Fatalf("Failed to check password: %v", err)
	}

	err = CheckPassword("wrongpassword", hashedPassword)
	if err == nil {
		t.Fatalf("Expected error for wrong password, got nil")
	}
}

func TestValidateJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "test-secret"
	expiresIn := time.Hour * 24

	tokenString, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("Failed to make JWT: %v", err)
	}
	if tokenString == "" {
		t.Fatalf("Expected non-empty token string, got empty")
	}

	validUserID, err := ValidateJWT(tokenString, tokenSecret)
	if err != nil {
		t.Fatalf("Failed to validate JWT: %v", err)
	}

	if validUserID != userID {
		t.Fatalf("Expected user ID %v, got %v", userID, validUserID)
	}

	invalidToken := "invalid-token"
	_, err = ValidateJWT(tokenString, invalidToken)
	if err == nil {
		t.Fatalf("Expected error for invalid token, got nil")
	}

	modifiedToken := tokenString + "modified"
	_, err = ValidateJWT(modifiedToken, tokenSecret)
	if err == nil {
		t.Fatalf("Expected error for modified token, got nil")
	}
}

func TestMakeJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "test-secret"
	expiresIn := time.Millisecond * 300

	tokenString, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("Failed to make JWT: %v", err)
	}

	_, err = ValidateJWT(tokenString, tokenSecret)
	if err != nil {
		t.Fatalf("Expected no error for token, got %v", err)
	}

	time.Sleep(expiresIn)

	_, err = ValidateJWT(tokenString, tokenSecret)
	if err == nil {
		t.Fatalf("Expected error for expired token, got nil")
	}
}
