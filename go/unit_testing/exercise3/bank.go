package exercise3

import "errors"

type Database interface {
	SaveTransaction(accountID string, amount float64) error
	GetBalance(accountID string) (float64, error)
}

type Bank struct {
	db Database
}

func NewBank(db Database) *Bank {
	return &Bank{db: db}
}

func (b *Bank) Deposit(accountID string, amount float64) error {
	if accountID == "" {
		return errors.New("not account id")
	}
	if amount == 0 {
		return nil
	}

	if amount < 0 {
		return errors.New("amount can not be negative")
	}

	err := b.db.SaveTransaction(accountID, amount)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bank) Withdraw(accountID string, amount float64) error {
	if accountID == "" {
		return errors.New("not account id")
	}
	if amount == 0 {
		return nil
	}

	if amount < 0 {
		return errors.New("amount can not be positive")
	}

	balance, err := b.GetBalance(accountID)
	if err != nil {
		return err
	}

	if balance - amount < 0 {
		return errors.New("cant withdraw, not enought balance")
	}

	err = b.db.SaveTransaction(accountID, -amount)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bank) GetBalance(accountID string) (float64, error) {
	balance, err := b.db.GetBalance(accountID)
	if err != nil {
		return 0, err
	}

	return balance, nil
}
