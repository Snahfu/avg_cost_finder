package logic

import (
	"avg_cost_finder/model"
	"fmt"
	"math"
)

// Helper function to verify a solution
func verifySolution(orders []model.Transaksi, startQty, startAvg, targetQty, targetAvg float64) bool {
	qty := startQty
	avg := startAvg

	fmt.Printf("Initial: Qty=%.1f, Avg=%.4f\n", qty, avg)

	for i, order := range orders {
		newQty, newAvg, valid := Calculate(qty, avg, order)
		if !valid {
			fmt.Printf("Invalid order at step %d\n", i+1)
			return false
		}

		qty = newQty
		avg = newAvg

		fmt.Printf("Step %d: %s %d @ %.2f -> Qty=%.1f, Avg=%.4f\n",
			i+1, order.Type, order.Qty, order.Price, qty, avg)
	}

	qtyMatch := math.Abs(qty-targetQty) < 0.0001
	avgMatch := math.Abs(avg-targetAvg) < 0.0001

	fmt.Printf("Final: Qty=%.1f (target=%.1f), Avg=%.4f (target=%.4f)\n",
		qty, targetQty, avg, targetAvg)
	fmt.Printf("Matches target: Qty=%v, Avg=%v\n", qtyMatch, avgMatch)

	return qtyMatch && avgMatch
}
