package model

type Wallet struct {
	Id           uint   `json:"id"`
	CustomerName string `json:"customer_name"`
	Balance      uint   `json:"balance"`
}

type AmountPaid struct {
	Amount uint   `json:"money"`
	Type   string `json:"type"`
}

type Transaction struct {
	Id       uint   `json:"id"`
	WalletId uint   `json:"walletId"`
	Amount   uint   `json:"amount"`
	Type     string `json:"type"`
	Status   string `json:"status"`
}
