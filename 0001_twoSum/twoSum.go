package _001_twoSum


//解法一
func twoSum1(nums []int, target int) []int {
	length := len(nums)
	for k,v := range nums {
		for i := k + 1 ; i < length; i++ {
			if  v + nums[i] == target {
				return []int{k,i}
			}
		}
	}
	return nil
}
//解法二
func twoSum2(nums []int, target int) []int {

	index := make(map[int]int, len(nums))

	for k,v := range nums {
		if i, ok := index[target-v]; ok {
			return []int{k,i}
		}
		index[v] = k
	}
	return  nil
}

