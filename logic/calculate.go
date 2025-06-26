package logic

import (
	"avg_cost_finder/model"
	"avg_cost_finder/utils"
)

// Simulasikan penerapan satu order pada qty dan avg
// order.JenisInstruksi: 1 = BUY, 2 = SELL
func Calculate(qty int64, avg float64, order model.Transaksi) (int64, float64) {
	if order.Type == "BUY" {
		newQty := qty + order.Qty
		newAvg := ((float64(order.Qty) * order.Price) + (avg * float64(qty))) / float64(newQty)
		return newQty, utils.Round(newAvg, 4)
	} else if order.Type == "SELL" {
		return qty - order.Qty, avg
	}
	return qty, avg
}
