package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	a := uint64(x)
	b  := uint64(y)
	if a+b > math.MaxUint32 {
		return x+y, true
	}
	return x+y, false
}

func CeilNumber(f float64) float64 {
	return 0
}

func AlphabetSoup(s string) string {
	n := strings.Split(s, "")
	sort.Strings(n)
	return strings.Join(n, "")
}

func StringMask(s string, n uint) string {
	return ""
}

func WordSplit(arr [2]string) string {
	return ""
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
