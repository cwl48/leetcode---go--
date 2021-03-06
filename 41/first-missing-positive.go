package firstmisspositve

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode-cn.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (34.56%)
 * Total Accepted:    6.7K
 * Total Submissions: 19.4K
 * Testcase Example:  '[1,2,0]'
 *
 * 给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
 *
 * 示例 1:
 *
 * 输入: [1,2,0]
 * 输出: 3
 *
 *
 * 示例 2:
 *
 * 输入: [3,4,-1,1]
 * 输出: 2
 *
 *
 * 示例 3:
 *
 * 输入: [7,8,9,11,12]
 * 输出: 1
 *
 *
 * 说明:
 *
 * 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
 *
 */
func firstMissingPositive(nums []int) int {

	m := make(map[int]bool, len(nums))
	max := 0
	for _, v := range nums {
		m[v] = true
		if v > max {
			max = v
		}
	}

	for i := 1; i < max; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return max + 1
}
