package main

import (
	"avg_cost_finder/logic"
	"avg_cost_finder/model"
	"fmt"
)

func main() {
	// Hardcoded transaksi
	trans := []model.Transaksi{
		{ID: 1, Type: "SELL", Qty: 50000, Price: 2570},
		{ID: 2, Type: "SELL", Qty: 20000, Price: 2570},
		{ID: 3, Type: "SELL", Qty: 30000, Price: 2570},
		{ID: 4, Type: "SELL", Qty: 30000, Price: 2570},
		{ID: 5, Type: "SELL", Qty: 50000, Price: 2560},
		{ID: 6, Type: "SELL", Qty: 30000, Price: 2570},
		{ID: 7, Type: "SELL", Qty: 20000, Price: 2570},
		{ID: 8, Type: "SELL", Qty: 100000, Price: 2560},
		{ID: 9, Type: "BUY", Qty: 50000, Price: 2500},
		{ID: 10, Type: "BUY", Qty: 50000, Price: 2500},
	}
	var qtyAwal int64
	var targetQty int64

	qtyAwal = 1248800
	avgAwal := 2533.5357
	targetQty = 1018800
	targetAvg := 2530.9326

	numWorkers := 8 // Total Worker

	if len(trans) <= 10 {
		fmt.Println("🔍 Menggunakan Fixed Permutation Worker...")

		allPerm := logic.GeneratePermutations(trans)
		fmt.Println("Total permutasi:", len(allPerm))

		// Fixed worker logic
		batchSize := len(allPerm) / numWorkers
		resultChan := make(chan []model.Transaksi)
		var wg logic.SafeWaitGroup

		for i := 0; i < numWorkers; i++ {
			start := i * batchSize
			end := start + batchSize
			if i == numWorkers-1 {
				end = len(allPerm)
			}

			wg.Add(1)
			go logic.Worker(logic.WorkerParams{
				PermBatch:  allPerm[start:end],
				QtyAwal:    qtyAwal,
				AvgAwal:    avgAwal,
				TargetQty:  targetQty,
				TargetAvg:  targetAvg,
				ResultChan: resultChan,
				Wg:         &wg.WaitGroup,
			})
		}

		go func() {
			wg.Wait()
			close(resultChan)
		}()

		found := false
		for res := range resultChan {
			fmt.Println("✅ Solusi ditemukan:")
			for _, t := range res {
				fmt.Printf("  - %d %s %d @ %.2f\n", t.ID, t.Type, t.Qty, t.Price)
			}
			found = true
		}

		if !found {
			fmt.Println("❌ Tidak ditemukan solusi.")
		}

	} else {
		fmt.Println("🎲 Menggunakan Random Permutation Worker...")

		totalSimulasi := 3628800 // 10! simulasi target
		batchPerWorker := totalSimulasi / numWorkers

		resultChan := make(chan []model.Transaksi)
		doneChan := make(chan struct{})

		for i := 0; i < numWorkers; i++ {
			go logic.RandomWorker(logic.RandomWorkerParam{
				Transaksi:  trans,
				Simulasi:   batchPerWorker,
				QtyAwal:    qtyAwal,
				AvgAwal:    avgAwal,
				TargetQty:  targetQty,
				TargetAvg:  targetAvg,
				ResultChan: resultChan,
				DoneChan:   doneChan,
			})
		}

		selesai := 0
		found := false

		for !found && selesai < numWorkers {
			select {
			case res := <-resultChan:
				fmt.Println("✅ Solusi ditemukan:")
				for _, t := range res {
					fmt.Printf("  - %d %s %d @ %.2f\n", t.ID, t.Type, t.Qty, t.Price)
				}
				found = true
			case <-doneChan:
				selesai++
			}
		}

		if !found {
			fmt.Println("❌ Tidak ditemukan solusi setelah", totalSimulasi, "simulasi.")
		}
	}
}
