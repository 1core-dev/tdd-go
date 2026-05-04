package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		wallet.Withdraw(10)
		assertBalance(t, wallet, 10)
	})
}
