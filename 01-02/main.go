package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	nums, err := ints(f)
	if err != nil {
		log.Fatal(err)
	}
	answer := recurringSum(nums)

	log.Println(answer)
}

// Repeatedly sums the given integers until the same sum occurs twice.
func recurringSum(nums []int) int {
	sum := 0
	found := map[int]bool{0: true}
	for {
		for _, n := range nums {
			sum += n
			if found[sum] {
				return sum
			}
			found[sum] = true
		}
	}
}

// Ints gets a slice of ints from a line-separated list.
func ints(r io.Reader) ([]int, error) {
	ints := []int{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}
	return ints, nil
}
