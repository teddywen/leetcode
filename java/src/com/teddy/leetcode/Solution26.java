package com.teddy.leetcode;

import java.util.Arrays;
import java.util.Objects;
import java.util.stream.Collectors;

/**
 * @link https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
 */
class Solution26 {

    public static void main(String[] args) {
        int[] nums = new int[]{0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
        System.out.println(new Solution26().removeDuplicates(nums));
        System.out.println(Arrays.stream(nums).mapToObj(String::valueOf).collect(Collectors.joining(",")));
    }

    public int removeDuplicates(int[] nums) {
        Integer last = null;
        int removeNums = 0;
        for (int i = 0; i < nums.length; i++) {
            if (Objects.nonNull(last) && last.equals(nums[i])) {
                removeNums++;
                continue;
            }
            last = nums[i];
            if (removeNums > 0) {
                nums[i - removeNums] = nums[i];
            }
        }
        return nums.length - removeNums;
    }
}