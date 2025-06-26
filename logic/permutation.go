package logic

import (
	"avg_cost_finder/model"
)

// GeneratePermutations menghasilkan semua permutasi dari slice transaksi
func GeneratePermutations(trans []model.Transaksi) [][]model.Transaksi {
	var result [][]model.Transaksi
	permute(trans, 0, &result)
	return result
}

// Helper rekursif untuk permutasi
func permute(arr []model.Transaksi, start int, result *[][]model.Transaksi) {
	if start == len(arr)-1 {
		// Buat salinan slice agar tidak berubah saat backtrack
		perm := make([]model.Transaksi, len(arr))
		copy(perm, arr)
		*result = append(*result, perm)
		return
	}

	for i := start; i < len(arr); i++ {
		arr[start], arr[i] = arr[i], arr[start] // swap
		permute(arr, start+1, result)
		arr[start], arr[i] = arr[i], arr[start] // swap balik (backtrack)
	}
}
