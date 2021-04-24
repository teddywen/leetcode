package com.teddy.leetcode;

import java.util.Deque;
import java.util.LinkedList;

class Solution179 {

    public static void main(String[] args) {
        int[] nums = new int[]{0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
        System.out.println(new Solution179().longestSubarray(nums, 2));
    }

    public int longestSubarray(int[] nums, int limit) {
        //这里维护的是最大值们对应的下标
        Deque<Integer> maxQ = new LinkedList<>();
        Deque<Integer> minQ = new LinkedList<>();
        int ans = 0;
        //窗口左沿
        int start = 0;
        //窗口右沿
        for (int end = 0; end < nums.length; end++) {
            //右沿元素进入窗口、维护最大值单调队列
            while (!maxQ.isEmpty() && nums[maxQ.peekLast()] < nums[end]) {
                maxQ.pollLast();
            }
            maxQ.add(end);
            //右沿元素进入窗口、维护最小值单调队列
            while (!minQ.isEmpty() && nums[minQ.peekLast()] > nums[end]) {
                minQ.pollLast();
            }
            minQ.add(end);

            //如果当前窗口的最大值最小值的差大于 limit，则不断缩小窗口（左沿++），直至最大值变小或者最小值变大从而满足 limit 限制
            while (!maxQ.isEmpty() && !minQ.isEmpty() && nums[maxQ.peek()] - nums[minQ.peek()] > limit) {
                if (maxQ.peek() <= start) maxQ.poll();
                if (minQ.peek() <= start) minQ.poll();
                start++;
            }
            ans = Math.max(ans, end - start + 1);
        }
        return ans;
    }
}
