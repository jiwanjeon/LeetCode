func twoSum(nums []int, target int) []int {
    // irriating over both slices
    // checking for the target
    for i, left := range(nums) {
        for j, right := range(nums) {
            if left + right == target && i != j {
                return []int{i, j}
            }
        }
    }
    
    // just return nothing since there will always be an answer
    return nil
}