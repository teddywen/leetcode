package com.teddy.leetcode;

import java.util.Objects;
import java.util.TreeMap;

/**
 * @link https://leetcode-cn.com/problems/contains-duplicate-iii/
 */
class Solution220 {

    public static void main(String[] args) {
        int[] nums = new int[]{-2147483648, 2147483647};
        int k = 1;
        int t = 1;
        System.out.println((new Solution220()).containsNearbyAlmostDuplicate(nums, k, t));
    }

    public boolean containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        TreeMap<Long, Integer> tree = new TreeMap<>();

        int left = 0;
        int right = 0;

        while (right < nums.length) {
            Long c = ((Integer) nums[right]).longValue();
            if (tree.containsKey(c)) {
                tree.put(c, tree.get(c) + 1);
            } else {
                tree.put(c, 1);
            }
            right++;

            while (right - left - 1 > k) {
                Long c2 = ((Integer) nums[left]).longValue();
                Integer ct = tree.get(c2);
                if (ct > 1) {
                    tree.put(c2, ct - 1);
                } else {
                    tree.remove(c2);
                }
                left++;
            }

            if (tree.get(c) > 1) {
                return true;
            }
            Long ch = tree.higherKey(c);
            if (Objects.nonNull(ch) && Math.abs(ch - c) <= t) {
                return true;
            }
            Long cl = tree.lowerKey(c);
            if (Objects.nonNull(cl) && Math.abs(cl - c) <= t) {
                return true;
            }
        }

        return false;
    }
}