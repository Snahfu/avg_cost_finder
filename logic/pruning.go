package logic

import (
	"avg_cost_finder/model"
)

// ✅ Pruning 1: Cek apakah transaksi tersisa hanya SELL
func HanyaSisaSell(trans []model.Transaksi, idx int) bool {
	for i := idx; i < len(trans); i++ {
		if trans[i].Type == "BUY" {
			return false
		}
	}
	return true
}

// ✅ Pruning 2: Target Avg naik, tapi tidak ada BUY yang bisa menaikkan avg
func TidakAdaBuyYangBisaMenaikkanAvg(trans []model.Transaksi, idx int, avgNow, targetAvg float64) bool {
	if avgNow >= targetAvg {
		return false // tidak perlu naik lagi
	}
	for i := idx; i < len(trans); i++ {
		if trans[i].Type == "BUY" && trans[i].Price > avgNow {
			return false // ada BUY yang bisa dorong avg naik
		}
	}
	return true // semua BUY <= avgNow → mustahil naik
}

// ✅ Pruning 3: Target Avg turun, tapi tidak ada BUY yang bisa menurunkan avg
func TidakAdaBuyYangBisaMenurunkanAvg(trans []model.Transaksi, idx int, avgNow, targetAvg float64) bool {
	if avgNow <= targetAvg {
		return false // tidak perlu turun lagi
	}
	for i := idx; i < len(trans); i++ {
		if trans[i].Type == "BUY" && trans[i].Price < avgNow {
			return false // ada BUY yang bisa tarik avg turun
		}
	}
	return true // semua BUY >= avgNow → mustahil turun
}
