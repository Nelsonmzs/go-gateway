package repository

type AccountRepository struct {

	db *sql.DB

}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Save(account *domain.Account) error {

	stmt, err := r.db.Prepare("INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?)"
)
	
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		account.ID, account.Name, account.Email, account.APIKey, account.Balance, account.CreatedAt, account.UpdatedAt
	)
	
	if err != nil {
		return err
	}

	return nil
	
	
}

func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {

	var account = domain.Account
	var created_at, updated_at time.Time

	err := r.db.QueryRow(
		'SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = $1',
		apiKey,
	).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&created_at,
		&updated_at,
	)

	if (err == sql.ErrNoRows) {
		return nil, domain.ErrAccountNotFound
	}

	if (err != nil) {
		return nil, err
	}

	account.CreatedAt = created_at
	account.UpdatedAt = updated_at

	return &account, nill

}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {

	var account = domain.Account
	var created_at, updated_at time.Time

	err := r.db.QueryRow(
		'SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE id = $1',
		id,
	).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&created_at,
		&updated_at,
	)

	if (err == sql.ErrNoRows) {
		return nil, domain.ErrAccountNotFound
	}

	if (err != nil) {
		return nil, err
	}

	account.CreatedAt = created_at
	account.UpdatedAt = updated_at

	return &account, nill

}


func (r *AccountRepository) Update(account *domain.Account) error {

	tx, err := r.db.Begin()

	if (err != nil) {
		return err
	}

	defer tx.Rollback()

	var current_balance float64

	err = tx.QueryRow('SELECT balance FROM accounts WHERE id = $1' FOR UPDATE, account.ID)
	.Scan(&current_balance)

	if (err != nil) {
		return err
	}

	tx.Exec(UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3, account.Balance, time.Now(), account.ID)

	if (err != nil) {
		return err
	}

	tx.Commit()
	
	
}
