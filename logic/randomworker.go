// // ===============================
// // === SKENARIO DATA <= 10 ITEM ===
// // Kode di bawah ini digunakan untuk kasus di mana jumlah transaksi (data) kurang dari atau sama dengan 10.
// // Pada skenario ini, semua permutasi urutan transaksi akan di-generate dan diperiksa satu per satu secara exhaustif.
// // Hal ini dimungkinkan karena jumlah permutasi masih dalam batas wajar untuk diproses secara paralel oleh worker.
// // ===============================
package logic

// import (
// 	"avg_cost_finder/model"
// 	"avg_cost_finder/utils"
// 	"math/rand"
// 	"time"
// )

// // RandomWorkerParam digunakan untuk random permutation
// type RandomWorkerParam struct {
// 	Transaksi  []model.Transaksi
// 	Simulasi   int
// 	QtyAwal    int64
// 	AvgAwal    float64
// 	TargetQty  int64
// 	TargetAvg  float64
// 	ResultChan chan []model.Transaksi
// 	DoneChan   chan struct{}
// }

// func RandomWorker(p RandomWorkerParam) {
// 	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	nominalAwal := float64(p.QtyAwal) * p.AvgAwal

// 	for i := 0; i < p.Simulasi; i++ {
// 		copyTrans := make([]model.Transaksi, len(p.Transaksi))
// 		copy(copyTrans, p.Transaksi)
// 		randGen.Shuffle(len(copyTrans), func(i, j int) {
// 			copyTrans[i], copyTrans[j] = copyTrans[j], copyTrans[i]
// 		})

// 		qtyNow := p.QtyAwal
// 		nominalNow := nominalAwal
// 		valid := true
// 		path := []model.Transaksi{}

// 		for _, t := range copyTrans {
// 			qtyNew, avgNew := Calculate(qtyNow, nominalNow/float64(qtyNow), t)
// 			if qtyNew < 0 {
// 				valid = false
// 				break
// 			}
// 			nominalNow = float64(qtyNew) * avgNew
// 			qtyNow = qtyNew
// 			path = append(path, t)
// 		}

// 		if !valid || qtyNow != p.TargetQty {
// 			continue
// 		}

// 		avgNow := nominalNow / float64(qtyNow)
// 		if utils.Round(avgNow, 4) == utils.Round(p.TargetAvg, 4) {
// 			p.ResultChan <- path
// 			return
// 		}
// 	}

// 	// Jika tidak ada hasil yang cocok, beri sinyal selesai
// 	p.DoneChan <- struct{}{}
// }
