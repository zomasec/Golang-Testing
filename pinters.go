package pointers

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.balance -= amount
	return nil
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
