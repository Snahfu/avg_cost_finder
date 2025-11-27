package logic

import (
	"avg_cost_finder/model"
	"avg_cost_finder/utils"
)

// Simulasikan penerapan satu order pada qty dan avg
func Calculate(qty, avg float64, order model.Transaksi) (float64, float64, bool) {
	if order.Type == "BUY" { // Buy
		newQty := qty + float64(order.Qty)
		if newQty <= 0 {
			return qty, avg, false
		}
		newAvg := (float64(order.Qty)*order.Price + avg*qty) / newQty
		return newQty, utils.Round(newAvg, 6), true
	} else if order.Type == "SELL" { // Sell
		newQty := qty - float64(order.Qty)
		if newQty < 0 {
			return qty, avg, false
		}
		return newQty, avg, true
	}
	return qty, avg, false
}
