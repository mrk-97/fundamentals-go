package main

type CheckingAccount struct {
	id           int
	balance      (float64)
	transactions ([]string)
}

type SavingsAccount struct {
	id           int
	balance      (float64)
	interestRate (float64)
	transactions ([]string)
}

type Account interface {
	deposit(amount float64)
	withdraw(amount float64) error
	getBalance() float64
	printStatement()
}

type Bank struct {
	accounts      (map[int]Account)
	nextAccountId (int)
}

func openCheckingAccount() int {

}

func openSavingsAccount(interestRate float64) int {

}

func getAccount(id int) (Account, error) {

}

func bankDeposit(id int, amount float64) error {

}

func bankWithdraw(id int, amount float64) error {

}

func main() {

}
