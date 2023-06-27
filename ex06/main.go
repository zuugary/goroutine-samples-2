package main

import "sync"

type Balance struct {
	amount   float64
	currency string
	mu       *sync.RWMutex
}

func NewBalance(amount float64, currency string) *Balance {
	return &Balance{
		amount:   amount,
		currency: currency,
		mu:       &sync.RWMutex{},
	}
}

func (b *Balance) Add(i float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.amount += i
}

func (b *Balance) Amount() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.amount
}
