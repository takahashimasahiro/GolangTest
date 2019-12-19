package main

type Wallet struct {
	balance int
}

// func (w Wallet) Deposit(amount int) {
// 	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
// 	w.balance += amount
// }

// func (w Wallet) Balance() int {
// 	return w.balance
// }

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// ポインタ(*)を指定することでテストを通す
func (w *Wallet) Balance() int {
	return w.balance
}
