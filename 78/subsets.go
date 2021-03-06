package subsets


/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (69.54%)
 * Total Accepted:    10.3K
 * Total Submissions: 14.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */
func subsets(nums []int) [][]int {

	res := make([][]int, 0)
	// 给一个空数组
	res = append(res, []int{})
	for i := range nums {
		tmp := make([][]int, 0)
		for _, v := range res {
			ele := make([]int, 0)
			for _, e := range v {
				ele = append(ele, e)
			}
			ele = append(ele, nums[i])
			tmp = append(tmp, ele)
		}
		for _, v := range tmp {
			res = append(res, v)
		}
	}
	return res
}
