package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	nums2Counts := make(map[int]int)
	for fileScanner.Scan() {
		row := strings.Fields(fileScanner.Text())
		first, _ := strconv.Atoi(row[0])
		second, _ := strconv.Atoi(row[1])
		nums1 = append(nums1, first)
		nums2 = append(nums2, second)
		nums2Counts[second]++
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	fmt.Println(part1(nums1, nums2))
	fmt.Println(part2(nums1, nums2Counts))
}

func part1(nums1 []int, nums2 []int) int {
	var output int
	for i := 0; i < len(nums1); i++ {
		diff := nums1[i] - nums2[i]
		if diff < 0 {
			diff *= -1
		}
		output += diff
	}
	return output
}

func part2(nums1 []int, nums2Counts map[int]int) int {
	var output int
	for i := 0; i < len(nums1); i++ {
		output += nums1[i] * nums2Counts[nums1[i]]
	}
	return output
}
