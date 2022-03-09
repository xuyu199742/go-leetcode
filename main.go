package main

import "fmt"

type ListNode struct {
	Val int
	Nex *ListNode
}

func main() {
	fmt.Println("hello word")
	nums := []int{1, 2, 3, 4, 4, 5}
	//fmt.Println()
	removeElement(nums, 4)
	//fmt.Println(nums)
}

// 两数之和 map
func towSumMap(target int, nums []int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if val, ok := m[v]; ok {
			if k < val {
				return []int{k, val}
			} else {
				return []int{val, k}
			}
		}
		m[target-v] = k
	}

	return nil
}

// 两数之和双for
func towSumFor(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//单链表反转
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var curr = head

	for curr != nil {
		var next = curr.Nex
		curr.Nex = prev
		prev = curr
		curr = next
	}

	return prev
}

//数组二分查找
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}
	return -1
}

//数组移除元素
func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	//[]int{1, 2, 3, 4, 4, 5}
	for i := 0; i < length; i++ {
		if nums[i] == val {
			for j := i + 1; j < length; j++ {
				fmt.Println(j)
				//nums[i-1] = nums[j]
			}
			//length--
			//i--
		}
	}

	return length
}

//字符串反转
func reverseString(s []byte) {
	r, l := 0, len(s)-1
	for r < l {
		s[r], s[l] = s[l], s[r]
		l--
		r++
	}
}
