package main

import (
	"avg_cost_finder/logic"
	"avg_cost_finder/model"
	"fmt"
	// "fmt"
)

func main() {
	// INPUT SECTION
	// Hardcoded transaksi
	transactions := []model.Transaksi{
		{ID: 1, Type: "BUY", Qty: 30000, Price: 7650},
		{ID: 2, Type: "SELL", Qty: 30000, Price: 7950},
	}
	/*
		[
			{608 2 80000 1330}
			{620 1 50000 1340}
			{621 2 80000 1340}
			{622 2 80000 1340}
			{624 2 100000 1338.915}
			{627 2 100000 1335}
			{628 2 50000 1335}
			{629 2 80000 1330}
			{630 2 100000 1330}
		]
	*/
	//
	var startQty float64
	var startAvg float64
	var targetAvg float64
	var targetQty float64

	startQty = 275000
	targetQty = 275000
	startAvg = 8913.093104
	targetAvg = 8788.854438

	// Find solution with detailed logging
	solution, found := logic.BacktrackOrderSequence(transactions, startQty, startAvg, targetQty, targetAvg)
	if found {
		fmt.Println("🎉 SOLUTION FOUND!")
		fmt.Println("\n📋 Optimal order sequence:")
		for i, order := range solution {
			fmt.Printf("  %d. ID:%d %s %d @ %.2f\n",
				i+1, order.ID, order.Type, order.Qty, order.Price)
		}

		fmt.Println("\n🔍 Step-by-step verification:")
		qty := startQty
		avg := startAvg
		fmt.Printf("  Initial: Qty=%.1f, Avg=%.4f\n", qty, avg)

		for i, order := range solution {
			qty, avg, _ = logic.Calculate(qty, avg, order)
			fmt.Printf("  Step %d: %s %d @ %.2f → Qty=%.1f, Avg=%.6f\n",
				i+1, order.Type, order.Qty, order.Price, qty, avg)
		}

		fmt.Printf("\n✅ Final result: Avg=%.4f (target was %.6f)\n", avg, targetAvg)

	} else {
		fmt.Println("❌ NO SOLUTION FOUND")
		fmt.Printf("Target average %.6f cannot be achieved with the given orders\n", targetAvg)
	}
	// START PROCESS

	// numWorkers := 8 // Total Worker

	// if len(trans) <= 10 {
	// 	fmt.Println("🔍 Menggunakan Fixed Permutation Worker...")

	// 	allPerm := logic.GeneratePermutations(trans)
	// 	fmt.Println("Total permutasi:", len(allPerm))

	// 	// Fixed worker logic
	// 	batchSize := len(allPerm) / numWorkers
	// 	resultChan := make(chan []model.Transaksi)
	// 	var wg logic.SafeWaitGroup

	// 	for i := 0; i < numWorkers; i++ {
	// 		start := i * batchSize
	// 		end := start + batchSize
	// 		if i == numWorkers-1 {
	// 			end = len(allPerm)
	// 		}

	// 		wg.Add(1)
	// 		go logic.Worker(logic.WorkerParams{
	// 			PermBatch:  allPerm[start:end],
	// 			QtyAwal:    qtyAwal,
	// 			AvgAwal:    avgAwal,
	// 			TargetQty:  targetQty,
	// 			TargetAvg:  targetAvg,
	// 			ResultChan: resultChan,
	// 			Wg:         &wg.WaitGroup,
	// 		})
	// 	}

	// 	go func() {
	// 		wg.Wait()
	// 		close(resultChan)
	// 	}()

	// 	found := false
	// 	for res := range resultChan {
	// 		fmt.Println("✅ Solusi ditemukan:")
	// 		for _, t := range res {
	// 			fmt.Printf("  - %d %s %d @ %.2f\n", t.ID, t.Type, t.Qty, t.Price)
	// 		}
	// 		found = true
	// 	}

	// 	if !found {
	// 		fmt.Println("❌ Tidak ditemukan solusi.")
	// 	}

	// } else {
	// 	fmt.Println("🎲 Menggunakan Random Permutation Worker...")

	// 	totalSimulasi := 3628800 // 10! simulasi target
	// 	batchPerWorker := totalSimulasi / numWorkers

	// 	resultChan := make(chan []model.Transaksi)
	// 	doneChan := make(chan struct{})

	// 	for i := 0; i < numWorkers; i++ {
	// 		go logic.RandomWorker(logic.RandomWorkerParam{
	// 			Transaksi:  trans,
	// 			Simulasi:   batchPerWorker,
	// 			QtyAwal:    qtyAwal,
	// 			AvgAwal:    avgAwal,
	// 			TargetQty:  targetQty,
	// 			TargetAvg:  targetAvg,
	// 			ResultChan: resultChan,
	// 			DoneChan:   doneChan,
	// 		})
	// 	}

	// 	selesai := 0
	// 	found := false

	// 	for !found && selesai < numWorkers {
	// 		select {
	// 		case res := <-resultChan:
	// 			fmt.Println("✅ Solusi ditemukan:")
	// 			for _, t := range res {
	// 				fmt.Printf("  - %d %s %d @ %.2f\n", t.ID, t.Type, t.Qty, t.Price)
	// 			}
	// 			found = true
	// 		case <-doneChan:
	// 			selesai++
	// 		}
	// 	}

	// 	if !found {
	// 		fmt.Println("❌ Tidak ditemukan solusi setelah", totalSimulasi, "simulasi.")
	// 	}
	// }
}
