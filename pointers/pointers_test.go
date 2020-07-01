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


	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}