package internal

import "fmt"

// 通知
type Notification struct {
}

func (n *Notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
