package logic

import (
	"avg_cost_finder/model"
	"avg_cost_finder/utils"
	"fmt"
)

func Backtrack(path []model.Transaksi, idx int, qtyNow int64, nominalNow float64, targetQty int64, targetAvg float64, all []model.Transaksi) {
	if idx == len(all) {
		if qtyNow == targetQty {
			avgNow := nominalNow / float64(qtyNow)
			if utils.Round(avgNow, 4) == utils.Round(targetAvg, 4) {
				fmt.Println("✅ Solusi ditemukan:")
				for _, t := range path {
					fmt.Printf("  - %d %s %d @ %.2f\n", t.ID, t.Type, t.Qty, t.Price)
				}
				fmt.Printf("Qty akhir: %d | Avg akhir: %.4f\n", qtyNow, avgNow)
			}
		}
		return
	}

	fmt.Print("Masih mencari")

	t := all[idx]

	// Gunakan function Calculate untuk menghitung qty dan avg baru
	qtyNew, avgNew := Calculate(qtyNow, nominalNow/float64(qtyNow), t)

	// Hitung nominal baru berdasarkan avg baru
	var nominalNew float64
	if qtyNew == 0 {
		nominalNew = 0
	} else {
		nominalNew = float64(qtyNew) * avgNew
	}

	// Pruning sederhana
	// Pruning 1: sisa hanya SELL, tapi avg belum cocok
	if HanyaSisaSell(all, idx) && utils.Round(avgNew, 4) != utils.Round(targetAvg, 4) {
		return
	}

	// Pruning 2: avg harus naik, tapi tidak ada BUY > avgNow
	if targetAvg > avgNew && TidakAdaBuyYangBisaMenaikkanAvg(all, idx, avgNew, targetAvg) {
		return
	}

	// Pruning 3: avg harus turun, tapi tidak ada BUY < avgNow
	if targetAvg < avgNew && TidakAdaBuyYangBisaMenurunkanAvg(all, idx, avgNew, targetAvg) {
		return
	}

	Backtrack(append(path, t), idx+1, qtyNew, nominalNew, targetQty, targetAvg, all)
}
