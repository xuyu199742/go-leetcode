package main

import "fmt"

type ListNode struct {
	Val int
	Nex *ListNode
}
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func main() {
	fmt.Println("hello word")
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

//删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	show := 1
	for fast := 1; fast < length; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[show] = nums[fast]
			show++
		}
	}

	return show
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

//二叉树最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(l, r int) int {
	if l > r {
		return l
	}

	return r
}
