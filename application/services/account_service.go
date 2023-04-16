package services

import (
	"financial-api/domain/entities"
	"financial-api/domain/repositories"
)

type AccountService struct {
	accountRepository repositories.AccountRepository
}

func NewAccountService(accountRepository repositories.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (s *AccountService) FindAll() ([]entities.Account, error) {
	return s.accountRepository.FindAll()
}

func (s *AccountService) FindByID(id int) (*entities.Account, error) {
	return s.accountRepository.FindByID(id)
}

func (s *AccountService) Create(account *entities.Account) error {
	return s.accountRepository.Create(account)
}

func (s *AccountService) Update(account *entities.Account) error {
	return s.accountRepository.Update(account)
}

func (s *AccountService) Delete(id int) error {
	account, err := s.accountRepository.FindByID(id)
	if err != nil {
		return err
	}
	return s.accountRepository.Delete(account)
}
