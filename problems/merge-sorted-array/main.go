package main

import "fmt"

func merge(nums1 []int, _ int, nums2 []int, n int) {
	for i := len(nums1) - 1; i >= n; i-- {
		nums1[i] = nums1[i-n]
	}

	curr1 := n
	curr2 := 0

	i := 0
	for curr1 < len(nums1) && curr2 < len(nums2) {
		if nums1[curr1] <= nums2[curr2] {
			nums1[i] = nums1[curr1]
			curr1++
		} else {
			nums1[i] = nums2[curr2]
			curr2++
		}
		i++
	}

	for curr1 < len(nums1) {
		nums1[i] = nums1[curr1]
		curr1++
		i++
	}

	for curr2 < len(nums2) {
		nums1[i] = nums2[curr2]
		curr2++
		i++
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
	fmt.Printf("nums1=%#v\n", nums1)
}
