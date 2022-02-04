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
	return math.Ceil(f * 4) / 4
}

func AlphabetSoup(s string) string {
	n := strings.Split(s, "")
	sort.Strings(n)
	return strings.Join(n, "")
}

func StringMask(s string, n uint) string {
	length := len(s)
	units := int(n)
	list := strings.Split(s, "")
	for i := 0; i < length; i++ {
		if i >= units {
			list[i] = "*"
		} else if length == 1 {
			list[i] = "*"
		} else if length <= units {
			list[i] = "*"
		}
	}
	if length == 0 {
		return "*"
	}
	return strings.Join(list, "")
}

func WordSplit(arr [2]string) string {
	return ""
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
