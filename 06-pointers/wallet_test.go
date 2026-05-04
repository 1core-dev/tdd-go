package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted error but didn't got one")
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

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(100)
		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})
}
