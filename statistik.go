package statistik

import (
	"fmt"
	//"math"
)

// rata-rata dari himpunan integer cth: 1,3,4,6,2,3,4
func Rerata(himpunan []int) (rerata float64) {
	if len(himpunan) == 0 {
		return 0.0
	}
	for _, v := range himpunan {
		rerata += float64(v)
	}
	return mean / float64(len(himpunan))
}
