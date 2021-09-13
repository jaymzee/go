package main

import "strconv"

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func utoa(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func itoa(i int64) string {
	return strconv.FormatInt(i, 10)
}
