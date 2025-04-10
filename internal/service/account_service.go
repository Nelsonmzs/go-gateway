package service

type AccountService struct {

	repository domain.AccountRepository

	func NewAccountService(repository domain.AccountRepository) *AccountService {
		return &AccountService{
			repository: repository,
		}
	}
	
	func (s *AccountService) CreateAccount(input *dto.CreateAccountInput) (*dto.AccountOutput, error) {
		account := dto.ToAccount(input)

		existingAccount, err := s.repository.FindByAPIKey(account.APIKey)

		if (err != nil && err != domain.ErrAccountNotFound) {
			return nil, err
		}

		if (existingAccount != nil) {
			return nil, domain.ErrorAccountDuplicatedKey
		}

		err = s.repository.Save(account)

		if (err != nil) {
			return nil, err
		}

		return dto.FromAccount(account), nil
	}
	

}