package model

type Transaksi struct {
	ID    int64
	Type  string // "BUY" or "SELL"
	Qty   int64
	Price float64
}
