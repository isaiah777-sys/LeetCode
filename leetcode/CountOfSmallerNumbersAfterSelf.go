package leetcode

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		count := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		result[i] = count
	}
	return result
}

// Another approach using a Binary Indexed Tree (Fenwick Tree) for better performance
type BIT struct {
	tree []int
}

func (b *BIT) update(index int, value int) {
	for index < len(b.tree) {
		b.tree[index] += value
		index += index & -index
	}
}

func (b *BIT) query(index int) int {
	sum := 0
	for index > 0 {
		sum += b.tree[index]
		index -= index & -index
	}
	return sum
}

func countSmallerBIT(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := make([]int, len(nums))
	// Coordinate compression
	ranks := make(map[int]int)
	for _, num := range nums {
		ranks[num] = 0
	}
	rank := 1
	for num := range ranks {
		ranks[num] = rank
		rank++
	}
	bit := &BIT{tree: make([]int, len(ranks)+1)}
	for i := len(nums) - 1; i >= 0; i-- {
		rank := ranks[nums[i]]
		result[i] = bit.query(rank - 1)
		bit.update(rank, 1)
	}
	return result
}
