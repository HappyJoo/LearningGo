package main

import (
	"testing"
	//"fmt"
)
func TestWallet(t *testing.T) {
	// wallet := Wallet{}

	// wallet.Deposit(Bitcoin(10))

	// got := wallet.Balance()

	// // fmt.Printf("address %v\n", &wallet.balance)
	// want := Bitcoin(110)

	// if got != want {
	// 	t.Errorf("got %s want %s", got, want)
	// }

//////////////////////////////////////////
	// t.Run("Deposit", func(t *testing.T) {
	// 	wallet := Wallet{}

	// 	wallet.Deposit(Bitcoin(10))

	// 	got := wallet.Balance()

	// 	want := Bitcoin(10)

	// 	if got != want {
	// 		t.Errorf("got %s want %s", got, want)
	// 	}
	// })

	// t.Run("Withdraw", func(t *testing.T) {
	// 	wallet := Wallet{balance: Bitcoin(20)}

	// 	wallet.Withdraw(Bitcoin(10))

	// 	got := wallet.Balance()

	// 	want := Bitcoin(10)

	// 	if got != want {
	// 		t.Errorf("got %s want %s", got, want)
	// 	}
	// })

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but want one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}