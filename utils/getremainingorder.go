package utils

import "avg_cost_finder/model"

type RemainingInfo struct {
	Orders     []model.Transaksi
	BuyCount   int
	SellCount  int
	BuyOrders  []model.Transaksi
	SellOrders []model.Transaksi
}

func GetRemainingOrdersDetailed(used []bool, orders []model.Transaksi) RemainingInfo {
	var info RemainingInfo

	for i, order := range orders {
		if !used[i] {
			info.Orders = append(info.Orders, order)
			if order.Type == "BUY" {
				info.BuyCount++
				info.BuyOrders = append(info.BuyOrders, order)
			} else {
				info.SellCount++
				info.SellOrders = append(info.SellOrders, order)
			}
		}
	}
	return info
}
