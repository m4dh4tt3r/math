package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Generate a slice of uint64 1..n given n and return
// the slice
func genSeqSlice(n uint64) []uint64 {
	var seqSlice []uint64

	for i := uint64(1); i <= n; i++ {
		seqSlice = append(seqSlice, i)
	}

	return seqSlice
}

// Given a slice of uint64, randomly permutate the
// elements and return the permutated slice
func permutateSlice(s []uint64) []uint64 {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })

	return s
}

// Given a slice, print its elements
func printSlice(s []uint64) {
	for i, v := range s {
		if i < len(s)-1 {
			fmt.Printf("%v ", v)
		} else {
			fmt.Printf("%v\n", v)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s n\n", os.Args[0])
	} else {
		n, _ := strconv.ParseUint(os.Args[1], 10, 64)
		printSlice(permutateSlice(genSeqSlice(n)))
	}
}
