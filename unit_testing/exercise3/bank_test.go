package exercise3

import (
	"testing"
)

type MockDatabase struct {
    balance float64
}

func (m *MockDatabase) SaveTransaction(accountID string, amount float64) error {
    m.balance += amount
    return nil
}

func (m *MockDatabase) GetBalance(accountID string) (float64, error) {
    return m.balance, nil
}

type BankTest struct {
	name string
	accountID string
	amount float64
	want float64
	wantErr bool
}

type BankTestErr struct {
	name string
	accountID string
	amount float64
	wantErr bool
}


func TestBank_Deposit(t *testing.T) {
	mockDB := &MockDatabase{}
	tests := []BankTestErr{
		{"Normal case", "1sa2", 100, false},
		{"Normal case big number", "1sa2", 10000000000000, false},
		{"Normal case float number", "1sa2", 0.112321, false},
		{"Negative deposit", "1sa2", -100, true},
		{"Not account ID", "", 100, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bank{
				db: mockDB,
			}
			if err := b.Deposit(tt.accountID, tt.amount); (err != nil) != tt.wantErr {
				t.Errorf("Bank.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBank_Withdraw(t *testing.T) {
	mockDB := &MockDatabase{}
	tests := []BankTestErr{
		{"Normal case", "1sa2", 100, false},
		{"Normal case big number", "1sa2", 10000000000000, false},
		{"Normal case float number", "1sa2", 0.112321, false},
		{"Negative deposit", "1sa2", -100, true},
		{"Not account ID", "", 100, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bank{
				db: mockDB,
			}
			if err := b.Withdraw(tt.accountID, tt.amount); (err != nil) != tt.wantErr {
				t.Errorf("Bank.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBank_GetBalance(t *testing.T) {
	mockDB := &MockDatabase{}
	tests := []BankTest{
		{"Zero balance", "12as", 0, 0, false},
		{"Deposit 30 balance", "12as", 30, 30, false},
		{"Withdraw balance", "12as", -30, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bank{
				db: mockDB,
			}

			if tt.amount >= 0 {
				b.Deposit(tt.accountID, tt.amount)
			} else {
				b.Withdraw(tt.accountID, tt.amount * -1)
			}
			got, err := b.GetBalance(tt.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bank.GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bank.GetBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
