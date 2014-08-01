package mean
import (
	"fmt"
	"math"
	"sort"
)

// Mean/rata-rata dari himpunan integer
func Mean(himpunan []int) (mean float64) {
	if len(himpunan) == 0 {
		return 0.
	}
	for _, n := range himpunan {
		mean += float64(n)	
	}
	return mean / float64(len(himpunan))

}
