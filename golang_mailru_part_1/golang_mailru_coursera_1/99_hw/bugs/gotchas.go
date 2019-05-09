package main

import (
	"strconv"
	"sort"
)

func ReturnInt() int {
	return 1;
}

func ReturnFloat() float32 {
	return 1.1;
}

func ReturnIntArray() ([3]int) {
	return [3]int{1, 3, 4};
}

func ReturnIntSlice() ([]int) {
	return []int{1, 2, 3};
}

func IntSliceToString(in []int) (res string) {
	for _, val := range in {
		res += strconv.Itoa(val)
	}
	return res;
}

func MergeSlices(fst []float32, scnd []int32) (res[]int) {
	for _, val := range fst {
		res = append(res, int(val))
	}
	for _, val := range scnd {
		res = append(res, int(val))
	}
	return
}

func GetMapValuesSortedByKey(in map[int]string) (res[]string) {
	keys := make([]int, 0, len(in))
	for k := range in {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		res = append(res, in[k])
	}
	return
}