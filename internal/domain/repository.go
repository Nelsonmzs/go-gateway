package domain

type AccountRepository interface {

	Save(account *Account) error

	FindByAPIKey(apiKey string) (*Account, error)

	FindByID(id string) (*Account, error)

	Update(account *Account)
	
}


