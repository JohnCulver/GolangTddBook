package wallet

import "testing"

func TestWallet(t *testing.T) {

	//TODO write assertBalance here https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#refactor-1

	t.Run("Deposit BTC", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw BTC", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
