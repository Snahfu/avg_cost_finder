package logic

import (
	"avg_cost_finder/model"
	"avg_cost_finder/utils"
	"math"
)

// func Backtrack(path []model.Transaksi, idx int, qtyNow int64, nominalNow float64, targetQty int64, targetAvg float64, all []model.Transaksi) {
// 	if idx == len(all) {
// 		if qtyNow == targetQty {
// 			avgNow := nominalNow / float64(qtyNow)
// 			if utils.Round(avgNow, 4) == utils.Round(targetAvg, 4) {
// 				fmt.Println("✅ Solusi ditemukan:")
// 				for _, t := range path {
// 					fmt.Printf("  - %d %s %d @ %.2f\n", t.ID, t.Type, t.Qty, t.Price)
// 				}
// 				fmt.Printf("Qty akhir: %d | Avg akhir: %.4f\n", qtyNow, avgNow)
// 			}
// 		}
// 		return
// 	}

// 	fmt.Print("Masih mencari")

// 	t := all[idx]

// 	// Gunakan function Calculate untuk menghitung qty dan avg baru
// 	qtyNew, avgNew := Calculate(qtyNow, nominalNow/float64(qtyNow), t)

// 	// Hitung nominal baru berdasarkan avg baru
// 	var nominalNew float64
// 	if qtyNew == 0 {
// 		nominalNew = 0
// 	} else {
// 		nominalNew = float64(qtyNew) * avgNew
// 	}

// 	// Pruning sederhana
// 	// Pruning 1: sisa hanya SELL, tapi avg belum cocok
// 	if HanyaSisaSell(all, idx) && utils.Round(avgNew, 4) != utils.Round(targetAvg, 4) {
// 		return
// 	}

// 	// Pruning 2: avg harus naik, tapi tidak ada BUY > avgNow
// 	if targetAvg > avgNew && TidakAdaBuyYangBisaMenaikkanAvg(all, idx, avgNew, targetAvg) {
// 		return
// 	}

// 	// Pruning 3: avg harus turun, tapi tidak ada BUY < avgNow
// 	if targetAvg < avgNew && TidakAdaBuyYangBisaMenurunkanAvg(all, idx, avgNew, targetAvg) {
// 		return
// 	}

// 	Backtrack(append(path, t), idx+1, qtyNew, nominalNew, targetQty, targetAvg, all)
// }

func BacktrackOrderSequence(orders []model.Transaksi, startQty, startAvg, targetQty, targetAvg float64) ([]model.Transaksi, bool) {
	n := len(orders)
	used := make([]bool, n)
	var solution []model.Transaksi

	var backtrack func(currentQty, currentAvg float64, depth int) bool
	backtrack = func(currentQty, currentAvg float64, depth int) bool {
		// Base case: all orders have been used
		if depth == n {
			// Check if we've reached both target quantity and average
			qtyMatch := math.Abs(currentQty-targetQty) < 0.0001
			avgMatch := math.Abs(currentAvg-targetAvg) < 0.0001
			return qtyMatch && avgMatch
		}

		// Try each unused order
		for i := 0; i < n; i++ {
			if used[i] {
				continue // Skip if already used
			}

			// Try this order
			newQty, newAvg, valid := Calculate(currentQty, currentAvg, orders[i])
			if !valid {
				continue // Skip invalid moves
			}

			// 🎯 GET REMAINING INFO AND PRUNING CHECK - RIGHT HERE
			// Simulate what would remain after using this order
			tempUsed := make([]bool, len(used))
			copy(tempUsed, used)
			tempUsed[i] = true // Simulate using this order

			remainingInfo := utils.GetRemainingOrdersDetailed(tempUsed, orders)

			// === Pruning Section ===
			if remainingInfo.BuyCount == 0 && remainingInfo.SellCount > 0 && newQty > 0 {
				// SELL tidak akan mengubah avg, jadi hanya lanjut jika avg sudah sama target
				if math.Abs(newAvg-targetAvg) > 0.0001 {
					continue // prune
				}
				// Kalau avg sudah sama target, biarkan lanjut
			}

			// Make the choice (assuming not prunned)
			used[i] = true
			solution = append(solution, orders[i])
			// fmt.Printf("%s📝 Added to solution: [", "")
			// for j, ord := range solution {
			// 	if j > 0 {
			// 		fmt.Print(", ")
			// 	}
			// 	fmt.Printf("%d", ord.ID)
			// }
			// fmt.Printf("]\n")

			// Recursively try to complete the solution
			if backtrack(newQty, newAvg, depth+1) {
				return true // Found a solution
			}

			// Backtrack: undo the choice
			used[i] = false
			solution = solution[:len(solution)-1]
		}

		return false // No solution found at this level
	}

	// Start the backtracking process
	found := backtrack(startQty, startAvg, 0)

	if found {
		// Return a copy of the solution
		result := make([]model.Transaksi, len(solution))
		copy(result, solution)
		return result, true
	}

	return nil, false
}
