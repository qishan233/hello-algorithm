package sort

func partition(nums []int, start, end int) int {
	// 退出条件
	if start == end {
		return end
	}

	if start > end {
		return -1
	}

	// 哨兵就位
	x := nums[end]

	// 指针就位
	left := start - 1
	right := start

	// 双指针开始遍历
	for ; right < end; right++ {
		if nums[right] <= x {
			left++
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	// 保持循环不变性
	left++

	// 哨兵就位
	nums[left], nums[end] = nums[end], nums[left]

	return left
}

func sort1(nums []int) {
	sortInternal(nums, 0, len(nums)-1)
}
func sortInternal(nums []int, start, end int) {
	if start >= end {
		return
	}

	p := partition(nums, start, end)

	sortInternal(nums, start, p-1)
	sortInternal(nums, p+1, end)
}

func partition2(nums []int, start, end int) int {
	// 退出条件
	if start == end {
		return end
	}

	if start > end {
		return -1
	}

	// 哨兵就位
	x := nums[end]

	// 指针就位
	left := start - 1
	right := start

	for left < right {
		// 左指针一直前进
		for ; left < right; left++ {
			if nums[left] >= x {
				break
			}
		}
		// 右指针一直前进
		for ; right > left; right-- {
			if nums[right] < x {
				break
			}
		}

		// 交换
		nums[left], nums[end] = nums[end], nums[left]

		// 保持循环不变性
		left++
	}

	nums[right], nums[end] = nums[end], nums[right]

	return right
}

func sort2(nums []int) {
	sortInternal2(nums, 0, len(nums)-1)
}
func sortInternal2(nums []int, start, end int) {
	if start >= end {
		return
	}

	p := partition(nums, start, end)

	sortInternal(nums, start, p-1)
	sortInternal(nums, p+1, end)
}
