package statistik
import (
        "fmt"
        "math"
)

func Mean(data Interface) (mean float64){
        Len := data.Len()
        for i:=0;i<Len;i++{
                mean += (data.Value(i) - mean) / float64(i+1)
        }
}
