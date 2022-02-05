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
	if length == 0 || length == 1{
		return "*"
	}
	for i := 0; i < length; i++ {
		if i >= units || length <= units {
			list[i] = "*"
		}
	}
	return strings.Join(list, "")
}

func WordSplit(arr [2]string) string {

	if len(arr[0]) == 0 { // empty string input will return "not possible"
		return "not possible"
	}

	// initialized required variables
	input := arr[0]
	words := strings.Split(arr[1], ",")
	var result []string

	// looks for the dictionary words in the input string
	for i := 0; i < len(words); i++ {
		if strings.Contains(input, words[i]) {
			result = append(result, words[i])
		}
	}

	// if none were found OR some words were found THEN the return value will be "not possible"
	if result == nil || len(strings.Join(result, "")) < len(input){
		return "not possible"
	}

	// Below for loops arranges the index of the return value

	// appends the first word
	inputArr := strings.Split(input, "")
	var finalResult []string
	for i := 0; i < len(result); i++ {
		if strings.Contains(result[i], inputArr[0]) {
			finalResult = append(finalResult, result[i])
		}
	}

	// appends the remaining word
	for i := 0; i < len(result); i++ {
		if !strings.Contains(finalResult[0], result[i]) {
			finalResult = append(finalResult, result[i])
		}
	}

	// returns the final value
	return strings.Join(finalResult, ",")
}

func VariadicSet(i ...interface{}) []interface{} {
	items := make(map[interface{}]bool) // create an (items) map with Key-Type of (Interface) and Value-Type of (Boolean)
	var result []interface{}
	for _, item := range i { // loops through i Interface's items
		n := items[item] // initialize i's nth (item's) boolean value to n
		if !n {          // IF n is !false which means i's nth (item) isn't available in items variable, THEN
			items[item] = true            // add (item) interface as Key-Type and (true) as Value-Type to items map
			result = append(result, item) // and append (item) to (result) interface
		}
	}
	return result
}
