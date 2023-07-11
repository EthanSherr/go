package main

func main() {

}

func fourSum(_nums []int, target int) [][]int {
	numsMap := make(map[int]bool, len(_nums))
	for _, n := range _nums {
		numsMap[n] = true
	}
	nums := make([]int, 0)
	for n, _ := range numsMap {
		nums = append(nums, n)
	}

	results := make([][]int, 0)
	for i, n1 := range nums {
		for j, n2 := range nums[i+1:] {
			for _, twoSumResult := range twoSum(nums[j+1:], target-n2-n1) {
				results = append(results, append(twoSumResult, n1, n2))
			}
		}
	}
	return results
}

func twoSum(nums []int, target int) [][]int {
	indices := make([][]int, 0)
	visited := make(map[int]int)

	for i, n := range nums {
		if j, ok := visited[target-n]; ok {
			indices = append(indices, []int{n, nums[j]})
		}
		visited[n] = i
	}

	return indices
}
