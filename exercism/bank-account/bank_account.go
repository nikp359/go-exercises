package account

import "sync"

// Account ...
type Account struct {
	deposit int64
	isOpen  bool
	m       sync.RWMutex
}

// Open ...
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		deposit: initialDeposit,
		isOpen:  true,
	}
}

// Balance ...
func (a *Account) Balance() (balance int64, ok bool) {
	a.m.RLock()
	defer a.m.RUnlock()

	return a.deposit, a.isOpen
}

// Deposit ...
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()

	if !a.isOpen {
		return 0, false
	}

	if a.deposit+amount < 0 {
		return a.deposit, false
	}

	a.deposit += amount

	return a.deposit, a.isOpen
}

// Close ...
func (a *Account) Close() (payout int64, ok bool) {
	a.m.Lock()
	defer func() {
		if a.isOpen {
			a.deposit = 0
			a.isOpen = false
		}
		a.m.Unlock()
	}()

	return a.deposit, a.isOpen
}
