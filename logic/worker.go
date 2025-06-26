// ===============================
// === SKENARIO DATA <= 10 ITEM ===
// Kode di bawah ini digunakan untuk kasus di mana jumlah transaksi (data) kurang dari atau sama dengan 10.
// Pada skenario ini, semua permutasi urutan transaksi akan di-generate dan diperiksa satu per satu secara exhaustif.
// Hal ini dimungkinkan karena jumlah permutasi masih dalam batas wajar untuk diproses secara paralel oleh worker.
// ===============================
package logic

import (
	"avg_cost_finder/model"
	"avg_cost_finder/utils"
	"sync"
)

// WorkerParams digunakan untuk worker fixed permutation
type WorkerParams struct {
	PermBatch  [][]model.Transaksi
	QtyAwal    int64
	AvgAwal    float64
	TargetQty  int64
	TargetAvg  float64
	ResultChan chan []model.Transaksi
	Wg         *sync.WaitGroup
}

func Worker(p WorkerParams) {
	defer p.Wg.Done()

	nominalAwal := p.AvgAwal * float64(p.QtyAwal)

	for _, perm := range p.PermBatch {
		qtyNow := p.QtyAwal
		nominalNow := nominalAwal
		valid := true
		path := []model.Transaksi{}

		for _, t := range perm {
			qtyNew, avgNew := Calculate(qtyNow, nominalNow/float64(qtyNow), t)
			if qtyNew < 0 {
				valid = false
				break
			}
			nominalNow = float64(qtyNew) * avgNew
			qtyNow = qtyNew
			path = append(path, t)
		}

		if !valid || qtyNow != p.TargetQty {
			continue
		}

		avgNow := nominalNow / float64(qtyNow)
		if utils.Round(avgNow, 4) == utils.Round(p.TargetAvg, 4) {
			p.ResultChan <- path
			return
		}
	}
}
