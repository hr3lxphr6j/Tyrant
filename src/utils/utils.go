package utils

import (
	"fmt"
	"math"
)

var (
	sizeUnits = []string{"B", "KB", "MB", "GB", "TB", "PB"}
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func HumanSize(n uint64) string {
	if n == 0 {
		return "0B"
	}
	i := IntMin(int(math.Log2(float64(n))/math.Log2(1024)), len(sizeUnits)-1)
	return fmt.Sprintf("%.2f%s", float64(n)/(math.Pow(1024, float64(i))), sizeUnits[i])
}
