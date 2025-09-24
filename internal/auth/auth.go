package auth

import (
    "errors"
    "time"
    "golang.org/x/crypto/bcrypt"
    "github.com/google/uuid"
)

// Dummy user struct — ideally load this from DB
type User struct {
    Username     string
    PasswordHash string
    Token        string
    TokenExpiry  time.Time
}

var users = map[string]*User{
    "admin": {
        Username:     "admin",
        PasswordHash: "$2a$12$uFjHaJvAF0AhIjN.H7pBIOTkt7QDPebENpRY3Z4XzZXBe61K/YJOO", // "password"
    },
}

// Authenticates a user and returns a new token
func Authenticate(username, password string) (string, error) {
    user, ok := users[username]
    if !ok {
        return "", errors.New("user not found")
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    // Generate a new token
    token := uuid.New().String()
    user.Token = token
    user.TokenExpiry = time.Now().Add(24 * time.Hour)

    return token, nil
}

// Validates an auth token
func ValidateToken(token string) (*User, error) {
    for _, user := range users {
        if user.Token == token && time.Now().Before(user.TokenExpiry) {
            return user, nil
        }
    }
    return nil, errors.New("invalid or expired token")
}

// Optional: Helper to create a hashed password
func HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash), err
}

