package main

import "fmt"

func main() {
    nums := []int{2, 19, 7, 15}
    target := 9

    result := twoSum(nums, target)
    fmt.Println(result) // Output: [0 2]
}

func twoSum(nums []int, target int) []int {
    // Create a map to store the complement of each element and its index
    complementMap := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if index, found := complementMap[complement]; found {
            return []int{index, i}
        }
        complementMap[num] = i
    }
    return []int{}
}