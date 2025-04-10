package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID string
	Name string
	Email string
	APIKey string
	Balance float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateAPIKey() string {
	key := make([]byte, 16)
	rand.Read(key)
	return hex.EncodeToString(key)
}

func NewAccount(name, email string) (*Account, error) {

	account := &Account{

		ID: uuid.New().String(),
		Name: name,
		Email: email,
		APIKey: generateAPIKey(),
		Balance: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
	
}

func addBalance( a *Account, amount float64) {
	a.Balance += amount
	a.UpdatedAt = time.Now()
}
	

