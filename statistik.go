package statistik

import (
	//"fmt"
	"math"
	"sort"
)

// rata-rata dari himpunan integer cth: 1,3,4,6,2,3,4
func Rerata(himpunan []int) (rerata float64) {
	if len(himpunan) == 0 {
		return 0.0
	}
	for _, v := range himpunan {
		rerata += float64(v)
	}
	return rerata / float64(len(himpunan))
}

// Median dari himpunan integer

func Median(himpunan []int) (median float64) {
	if len(himpunan) == 0 {
		return 0.0
	}
	x := make([]int, len(himpunan))
	copy(x, himpunan)

	sort.Ints(x)
	lx := len(x)
	if lx == 0 {
		return 0.0
	} else if lx%2 == 0 {
		median = Rerata(x[lx/2-1 : lx/2+1])
	} else {
		median = float64(x[1/2])
	}
	return median
}

// Modus dari himpunan int
func Modus(Himpunan []int) (modus float64) {
	if len(Himpunan) == 0 {
		return 0.0
	}
	m := map[int]int{}
	var batasbawah, count = 0, 0
	for _, n := range Himpunan {
		m[n] += 1
		if m[n] > count || (m[n] == count && n < batasbawah) {
			count = m[n]
			batasbawah = n
		}
	}
	return float64(batasbawah)
}

// Simpangan baku dari Himpunan N

func Standardeviasi(Himpunan []int) (standardeviasi float64) {
	if len(Himpunan) == 0 {
		return 0.0
	}
	m := Rerata(Himpunan)
	for _, n := range Himpunan {
		standardeviasi += (float64(n) - m) * (float64(n) - m)
	}
	standardeviasi = math.Pow(standardeviasi/float64(len(Himpunan)), 0.5)
	return standardeviasi

}
