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
		// {ID: 1, Type: "SELL", Qty: 30000, Price: 9550},
		// {ID: 4, Type: "BUY", Qty: 50000, Price: 9275},
		// {ID: 2, Type: "SELL", Qty: 20000, Price: 9425},
		// {ID: 3, Type: "SELL", Qty: 20000, Price: 9475},
		{ID: 1, Type: "SELL", Qty: 50000, Price: 2570},
		{ID: 2, Type: "SELL", Qty: 20000, Price: 2570},
		{ID: 3, Type: "SELL", Qty: 30000, Price: 2570},
		{ID: 4, Type: "SELL", Qty: 30000, Price: 2570},
		{ID: 6, Type: "SELL", Qty: 30000, Price: 2570},
		{ID: 7, Type: "SELL", Qty: 20000, Price: 2570},
		{ID: 5, Type: "SELL", Qty: 50000, Price: 2560},
		{ID: 8, Type: "SELL", Qty: 100000, Price: 2560},
		{ID: 9, Type: "BUY", Qty: 50000, Price: 2500},
		{ID: 10, Type: "BUY", Qty: 50000, Price: 2500},
	}
	var startQty float64
	var startAvg float64
	var targetAvg float64
	var targetQty float64

	startQty = 1248800
	targetQty = 1018800
	startAvg = 2533.535745
	targetAvg = 2530.932676
	// startQty = 485000
	// targetQty = 465000
	// startAvg = 10070.0659
	// targetAvg = 9991.3465

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
			fmt.Printf("  Step %d: %s %d @ %.2f → Qty=%.1f, Avg=%.4f\n",
				i+1, order.Type, order.Qty, order.Price, qty, avg)
		}

		fmt.Printf("\n✅ Final result: Avg=%.4f (target was %.4f)\n", avg, targetAvg)

	} else {
		fmt.Println("❌ NO SOLUTION FOUND")
		fmt.Printf("Target average %.4f cannot be achieved with the given orders\n", targetAvg)
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
