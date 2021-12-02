// https://adventofcode.com/2021/day/1
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var (
		x, delta, increases int
		window, depth       []int
	)
	// It's helpful to have the test data be the same type as the file so strings, even though they are ints.
	// var input = []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}

	// Read in the file as a splice of strings
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		//Do something (I didn't do anything, because I know what's in the file and didn't expect weird errors)
	}
	input := strings.Split(string(content), "\n")

	// Convert the strings to ints.
	for _, i := range input {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		depth = append(depth, j)
	}

	// Find the three value sum for each window centered on x.
	// Slice starts at 0 so 1 is the earliest window
	// The last entry in the slice doesn't have a +1 so can't use it.
	// Append each window to the empty window slice
	for x = 1; x < (len(depth) - 1); x++ {
		delta = depth[x-1] + depth[x] + depth[x+1]
		window = append(window, delta)
	}
	// Once we have the windows, it's the same math as before.
	// I was getting 1712 with strings and 1713 with ints...types.
	for x = 1; x < (len(window)); x++ {
		if window[x-1] < window[x] {
			increases += 1
		}
	}
	fmt.Println(increases)
}
