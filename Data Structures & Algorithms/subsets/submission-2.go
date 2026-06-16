func subsets(nums []int) [][]int {
    sort.Ints(nums)  // ← Must sort first!
    result := [][]int{}
    current := []int{}
    
    var backtrack func(start int)
    backtrack = func(start int) {
        subset := make([]int, len(current))
        copy(subset, current)
        result = append(result, subset)
        
        for i := start; i < len(nums); i++ {
            // Skip duplicates at same recursion level
            if i > start && nums[i] == nums[i-1] {
                continue
            }
            
            current = append(current, nums[i])
            backtrack(i + 1)
            current = current[:len(current)-1]
        }
    }
    
    backtrack(0)
    return result
}