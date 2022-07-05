// The objective of the Memento is to have some sort of token that acts like a point in time were
// we can roll back the system to

package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memento // here we have an array of mementos that we can use to roll back the system
	current int
}

func (b *BankAccount) String() string {
	return fmt.Sprint("Balance = $", b.balance,
		", current = ", b.current)
}

func NewBankAccount(balance int) *BankAccount {
	b := &BankAccount{balance: balance}
	b.changes = append(b.changes, &Memento{balance}) // the creation of the bank account is a change so we log it to the set of mementos
	return b
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	m := Memento{b.balance}
	b.changes = append(b.changes, &m) // the deposit is a change so we log it to the set of mementos
	b.current++
	fmt.Println("Deposited", amount,
		", balance is now", b.balance)
	return &m
}

func (b *BankAccount) Restore(m *Memento) {
	if m != nil {
		b.balance -= m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current] // the undo() is a change so we log it to the set of mementos
		b.balance = m.Balance
		return m
	}
	return nil // nothing to undo
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current] // the redo() is a change so we log it to the set of mementos
		b.balance = m.Balance
		return m
	}
	return nil
}

func main() {
	ba := NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1:", ba)
	ba.Undo()
	fmt.Println("Undo 2:", ba)
	ba.Redo()
	fmt.Println("Redo:", ba)
}
